// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/bntp"
	"github.com/JonasMuehlmann/bntp.go/bntp/backend"
	"github.com/JonasMuehlmann/bntp.go/bntp/libbookmarks"
	"github.com/JonasMuehlmann/bntp.go/bntp/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/bntp/libtags"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stoewer/go-strcase"

	"github.com/JonasMuehlmann/bntp.go/model/domain"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"

	fsRepository "github.com/JonasMuehlmann/bntp.go/model/repository/fs"
	// mssqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/mssql"
	// mysqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/mysql"
	// psqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/psql".
	sqlite3Repository "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"
)

var (
	PassedConfigPath   string
	ConfigDir          string
	ConfigFileBaseName = "bntp"
	ConfigSearchPaths  []string
)

type message struct {
	Level log.Level
	Msg   string
}

func newMessage(level log.Level, msg string) message {
	message := message{}

	message.Level = level
	message.Msg = msg

	return message
}

func GetDefaultDBConfig() DBConfig {
	return DBConfig{
		Driver:     "sqlite3",
		DataSource: path.Join(ConfigDir, "bntp_db.sql"),
	}
}

func GetDefaultSettings() map[string]any {
	return map[string]any{
		LogFile:         path.Join(ConfigDir, "bntp.log"),
		ConsoleLogLevel: log.ErrorLevel.String(),
		FileLogLevel:    log.InfoLevel.String(),
		DB:              GetDefaultDBConfig(),
		Backend: BackendConfig{
			Bookmarkmanager: BookmarkManagerConfig{
				BookmarkRepository: BookmarkRepositoryConfig{
					DB: GetDefaultDBConfig(),
				},
			},
			TagsManager: TagsManagerConfig{
				TagsRepository: TagsRepositoryConfig{
					DB: GetDefaultDBConfig(),
				},
			},
			DocumentManager: DocumentManagerConfig{
				DocumentRepository: DocumentRepositoryConfig{
					DB: GetDefaultDBConfig(),
				},
			},
			DocumentContentManager: DocumentContentManagerConfig{
				DocumentContentRepository: DocumentContentRepositoryConfig{
					DB: GetDefaultDBConfig(),
				},
			},
		},
	}
}

var pendingLogMessage []message

func addPendingLogMessage(level log.Level, format string, values ...any) {
	pendingLogMessage = append(pendingLogMessage, newMessage(level, fmt.Sprintf(format, values...)))
}

func formatValidationError(e validator.FieldError) string {
	var message string

	// TODO: Add test to validate if all used validator tags are handled here
	switch e.Tag() {
	case ValidatorLogrusLogLevel:
		message = fmt.Sprintf("value %q is invalid, allowed values are %v", e.Value(), log.AllLevels)
	case "required":
		message = "setting is required"
	case "file":
		message = fmt.Sprintf("value %q is not a path", e.Value())
	default:
		message = e.Error()
	}

	return message
}

func bfsSetDefault(path string, val any) error {
	switch v := val.(type) {
	case map[string]any:
		for k, v := range v {
			newPath := path
			if newPath == "" {
				newPath = k
			} else {
				newPath += "." + k
			}

			bfsSetDefault(newPath, v)
		}
	default:
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Struct {
			var asMap map[string]any
			err := mapstructure.Decode(v, &asMap)
			if err != nil {
				return err
			}

			for k, v := range asMap {
				newPath := path
				if newPath == "" {
					newPath = k
				} else {
					newPath += "." + k
				}

				err := bfsSetDefault(newPath, v)
				if err != nil {
					return err
				}

			}
		} else {
			viper.SetDefault(path, v)
		}
	}

	return nil
}

func setDefaultsFromStructOrMap(defaults any) error {
	return bfsSetDefault("", defaults)
}

func InitConfig() {

	pendingLogMessage = make([]message, 0, 5)

	if PassedConfigPath != "" {
		viper.SetConfigFile(PassedConfigPath)
	}

	goaoi.ForeachSliceUnsafe(ConfigSearchPaths, viper.AddConfigPath)

	err := setDefaultsFromStructOrMap(GetDefaultSettings())
	if err != nil {
		addPendingLogMessage(log.FatalLevel, "Error setting default values: %v", err)
	}

	viper.SetEnvPrefix("BNTP")
	viper.SetConfigName(ConfigFileBaseName)
	viper.AutomaticEnv() // read in environment variables that match

	err = viper.ReadInConfig()
	if err != nil && errors.Is(err, &viper.ConfigFileNotFoundError{}) {
		addPendingLogMessage(log.FatalLevel, "Error reading config: %v", err)
	}

	//*********************    Validate settings    ********************//
	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	var consoleLogLevel log.Level
	var fileLogLevel log.Level

	err = ConfigValidator.Struct(config)
	if err != nil {
		consoleLogLevel = log.ErrorLevel
		fileLogLevel = log.InfoLevel

		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			addPendingLogMessage(log.FatalLevel, "Error processing setting %v: %v", strcase.SnakeCase(e.Field()), formatValidationError(e))
		}
	} else {
		consoleLogLevel, _ = log.ParseLevel(viper.GetString(ConsoleLogLevel))
		fileLogLevel, _ = log.ParseLevel(viper.GetString(FileLogLevel))
	}
	// ******************** Log pending messages ********************* //

	helper.Logger = *helper.NewDefaultLogger(viper.GetString(LogFile), consoleLogLevel, fileLogLevel)

	for _, message := range pendingLogMessage {
		helper.Logger.Log(message.Level, message.Msg)
		if message.Level == log.FatalLevel {
			helper.Logger.Exit(1)
		}
	}

	if PassedConfigPath != "" {
		helper.Logger.Infof("Using config file %v", viper.ConfigFileUsed())
	} else {
		helper.Logger.Info("Using internal configuration")
	}
}

func NewDBFromConfig() *sql.DB {
	dataSource := viper.GetString(DB_DataSource)
	args := viper.GetStringSlice(DB_Args)

	openArgs := dataSource
	if len(args) > 0 {
		openArgs += "?"

		openArgs += strings.Join(args, "&")
	}

	db, err := sql.Open(viper.GetString(DB_Driver), openArgs)
	if err != nil {
		panic(err)
	}

	return db
}

//********************    Repository builders    ********************//
// TODO: Allow using non-sql repositories.
func NewBookmarkRepositoryFromConfig(repoDB *sql.DB) repository.BookmarkRepository {
	bookmarkRepository := new(sqlite3Repository.Sqlite3BookmarkRepository)

	bookmarkRepositoryAbstract, err := bookmarkRepository.New(sqlite3Repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: repoDB})
	if err != nil {
		panic(err)
	}

	bookmarkRepository, _ = bookmarkRepositoryAbstract.(*sqlite3Repository.Sqlite3BookmarkRepository)

	return bookmarkRepository
}

func NewTagsRepositoryFromConfig(repoDB *sql.DB) repository.TagRepository {
	tagsRepository := new(sqlite3Repository.Sqlite3TagRepository)

	tagsRepositoryAbstract, err := tagsRepository.New(sqlite3Repository.Sqlite3TagRepositoryConstructorArgs{DB: repoDB})
	if err != nil {
		panic(err)
	}

	tagsRepository, _ = tagsRepositoryAbstract.(*sqlite3Repository.Sqlite3TagRepository)

	return tagsRepository
}
func NewDocumentRepositoryFromConfig(repoDB *sql.DB) repository.DocumentRepository {
	documentRepository := new(sqlite3Repository.Sqlite3DocumentRepository)

	documentRepositoryAbstract, err := documentRepository.New(sqlite3Repository.Sqlite3DocumentRepositoryConstructorArgs{DB: repoDB})
	if err != nil {
		panic(err)
	}

	documentRepository, _ = documentRepositoryAbstract.(*sqlite3Repository.Sqlite3DocumentRepository)

	return documentRepository
}

// TODO: Allow non-fs content repositories
func NewDocumentContentRepositoryFromConfig(fs afero.Fs) repository.DocumentContentRepository {
	documentContentRepository := new(fsRepository.FSDocumentContentRepository)

	documentContentRepositoryAbstract, err := documentContentRepository.New(fs)
	if err != nil {
		panic(err)
	}

	documentContentRepository, _ = documentContentRepositoryAbstract.(*fsRepository.FSDocumentContentRepository)

	return documentContentRepository
}

// TODO: Implement proper  hook parsing
//********************    Manager builders    ********************//
func NewBookmarkManagerFromConfig(repo repository.BookmarkRepository) libbookmarks.BookmarkManager {
	hooks := new(bntp.Hooks[domain.Bookmark])

	bookmarkManager, err := libbookmarks.NewBookmarkManager(hooks, repo)
	if err != nil {
		panic(err)
	}

	return bookmarkManager
}

func NewTagsManagerFromConfig(repo repository.TagRepository) libtags.TagManager {
	hooks := new(bntp.Hooks[domain.Tag])
	tagsManager, err := libtags.NewTagmanager(hooks, repo)
	if err != nil {
		panic(err)
	}

	return tagsManager
}
func NewDocumentManagerFromConfig(repo repository.DocumentRepository) libdocuments.DocumentManager {
	hooks := new(bntp.Hooks[domain.Document])
	documentManager, err := libdocuments.NewDocumentManager(hooks, repo)
	if err != nil {
		panic(err)
	}

	return documentManager
}

func NewDocumentContentManagerFromConfig(repo repository.DocumentContentRepository) libdocuments.DocumentContentManager {
	hooks := new(bntp.Hooks[string])
	documentContentManager, err := libdocuments.NewDocumentContentRepository(hooks, repo)
	if err != nil {
		panic(err)
	}

	return documentContentManager
}

// **********************    Set up backend    **********************//
func NewBackendFromConfig() *backend.Backend {
	newBackend := new(backend.Backend)

	db := NewDBFromConfig()

	// TODO: Extend config to allow creating custom fs
	fs := afero.NewOsFs()

	newBackend.BookmarkManager = NewBookmarkManagerFromConfig(NewBookmarkRepositoryFromConfig(db))
	newBackend.TagManager = NewTagsManagerFromConfig(NewTagsRepositoryFromConfig(db))
	newBackend.DocumentManager = NewDocumentManagerFromConfig(NewDocumentRepositoryFromConfig(db))
	newBackend.DocumentContentManager = NewDocumentContentManagerFromConfig(NewDocumentContentRepositoryFromConfig(fs))

	return newBackend
}

func init() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ConfigDir = path.Join(userConfigDir, "bntp")

	ConfigSearchPaths = append(ConfigSearchPaths, ConfigDir)
	ConfigSearchPaths = append(ConfigSearchPaths, cwd)
}
