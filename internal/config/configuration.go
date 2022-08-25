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
	"io"
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
	"github.com/JonasMuehlmann/bntp.go/internal/marshallers"
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

type message struct {
	Level log.Level
	Msg   string
}

type ConfigManager struct {
	Viper             *viper.Viper
	DBOverride        *sql.DB
	FSOverride        afero.Fs
	PendingLogMessage []message
	Logger            *log.Logger

	PassedConfigPath   string
	ConfigDir          string
	ConfigFileBaseName string
	ConfigSearchPaths  []string
}

func NewConfigManager(stderr io.Writer, dbOverride *sql.DB, fsOverride afero.Fs) *ConfigManager {
	m := &ConfigManager{Viper: viper.New(), ConfigFileBaseName: "bntp"}

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	m.ConfigDir = path.Join(userConfigDir, "bntp")

	m.ConfigSearchPaths = append(m.ConfigSearchPaths, m.ConfigDir)
	m.ConfigSearchPaths = append(m.ConfigSearchPaths, cwd)

	if dbOverride != nil {
		m.DBOverride = dbOverride
	}
	if fsOverride != nil {
		m.FSOverride = fsOverride
	}

	m.PendingLogMessage = make([]message, 0, 5)

	if m.PassedConfigPath != "" {
		m.Viper.SetConfigFile(m.PassedConfigPath)
	}

	goaoi.ForeachSliceUnsafe(m.ConfigSearchPaths, m.Viper.AddConfigPath)

	err = m.setDefaultsFromStructOrMap(m.GetDefaultSettings())
	if err != nil {
		m.addPendingLogMessage(log.FatalLevel, "Error setting default values: %v", err)
	}

	m.Viper.SetEnvPrefix("BNTP")
	m.Viper.SetConfigName(m.ConfigFileBaseName)
	m.Viper.AutomaticEnv() // read in environment variables that match

	err = m.Viper.ReadInConfig()
	if err != nil && errors.Is(err, &viper.ConfigFileNotFoundError{}) {
		m.addPendingLogMessage(log.FatalLevel, "Error reading config: %v", err)
	}

	//*********************    Validate settings    ********************//
	var config Config

	err = m.Viper.Unmarshal(&config)
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
			m.addPendingLogMessage(log.FatalLevel, "Error processing setting %v: %v", strcase.SnakeCase(e.Field()), formatValidationError(e))
		}
	} else {
		consoleLogLevel, _ = log.ParseLevel(m.Viper.GetString(ConsoleLogLevel))
		fileLogLevel, _ = log.ParseLevel(m.Viper.GetString(FileLogLevel))
	}
	// ******************** Log pending messages ********************* //

	m.Logger = helper.NewDefaultLogger(m.Viper.GetString(LogFile), consoleLogLevel, fileLogLevel, stderr)

	for _, message := range m.PendingLogMessage {
		m.Logger.Log(message.Level, message.Msg)
		if message.Level == log.FatalLevel {
			log.Exit(1)
		}
	}

	if m.PassedConfigPath != "" {
		m.Logger.Infof("Using config file %v", m.Viper.ConfigFileUsed())
	} else {
		m.Logger.Info("Using internal configuration")
	}

	return m
}

func (m *ConfigManager) NewDBFromConfig() *sql.DB {
	if m.DBOverride != nil {
		return m.DBOverride
	}

	dataSource := m.Viper.GetString(DB_DataSource)
	args := m.Viper.GetStringSlice(DB_Args)

	openArgs := dataSource
	if len(args) > 0 {
		openArgs += "?"

		openArgs += strings.Join(args, "&")
	}

	db, err := sql.Open(m.Viper.GetString(DB_Driver), openArgs)
	if err != nil {
		panic(err)
	}

	return db
}

func (m *ConfigManager) GetDefaultDBConfig() DBConfig {
	return DBConfig{
		Driver:     "sqlite3",
		DataSource: path.Join(m.ConfigDir, "bntp_db.sql"),
	}
}

func (m *ConfigManager) GetDefaultSettings() map[string]any {
	defaultTagRepository := TagsRepositoryConfig{
		DB: m.GetDefaultDBConfig(),
	}

	return map[string]any{
		LogFile:         path.Join(m.ConfigDir, "bntp.log"),
		ConsoleLogLevel: log.ErrorLevel.String(),
		FileLogLevel:    log.InfoLevel.String(),
		DB:              m.GetDefaultDBConfig(),
		Backend: BackendConfig{
			Bookmarkmanager: BookmarkManagerConfig{
				BookmarkRepository: BookmarkRepositoryConfig{
					DB:            m.GetDefaultDBConfig(),
					TagRepository: defaultTagRepository,
				},
			},
			TagsManager: TagsManagerConfig{
				TagsRepository: defaultTagRepository,
			},
			DocumentManager: DocumentManagerConfig{
				DocumentRepository: DocumentRepositoryConfig{
					DB:            m.GetDefaultDBConfig(),
					TagRepository: defaultTagRepository,
				},
			},
			DocumentContentManager: DocumentContentManagerConfig{
				DocumentContentRepository: DocumentContentRepositoryConfig{
					DB: m.GetDefaultDBConfig(),
				},
			},
		},
	}
}

// ********************    Repository builders    ********************//
// TODO: Allow using non-sql repositories.
func (m *ConfigManager) NewBookmarkRepositoryFromConfig(logger *log.Logger, repoDB *sql.DB, tagRepository repository.TagRepository) repository.BookmarkRepository {
	bookmarkRepository := new(sqlite3Repository.Sqlite3BookmarkRepository)

	bookmarkRepositoryAbstract, err := bookmarkRepository.New(sqlite3Repository.Sqlite3BookmarkRepositoryConstructorArgs{Logger: logger, DB: repoDB, TagRepository: tagRepository})
	if err != nil {
		panic(err)
	}

	bookmarkRepository, _ = bookmarkRepositoryAbstract.(*sqlite3Repository.Sqlite3BookmarkRepository)

	return bookmarkRepository
}

func (m *ConfigManager) NewTagsRepositoryFromConfig(logger *log.Logger, repoDB *sql.DB) repository.TagRepository {
	tagsRepository := new(sqlite3Repository.Sqlite3TagRepository)

	tagsRepositoryAbstract, err := tagsRepository.New(sqlite3Repository.Sqlite3TagRepositoryConstructorArgs{Logger: logger, DB: repoDB})
	if err != nil {
		panic(err)
	}

	tagsRepository, _ = tagsRepositoryAbstract.(*sqlite3Repository.Sqlite3TagRepository)

	return tagsRepository
}
func (m *ConfigManager) NewDocumentRepositoryFromConfig(logger *log.Logger, repoDB *sql.DB, tagRepository repository.TagRepository) repository.DocumentRepository {
	documentRepository := new(sqlite3Repository.Sqlite3DocumentRepository)

	documentRepositoryAbstract, err := documentRepository.New(sqlite3Repository.Sqlite3DocumentRepositoryConstructorArgs{Logger: logger, DB: repoDB, TagRepository: tagRepository})
	if err != nil {
		panic(err)
	}

	documentRepository, _ = documentRepositoryAbstract.(*sqlite3Repository.Sqlite3DocumentRepository)

	return documentRepository
}

// TODO: Allow non-fs content repositories
func (m *ConfigManager) NewDocumentContentRepositoryFromConfig(logger *log.Logger, fs afero.Fs) repository.DocumentContentRepository {
	documentContentRepository := new(fsRepository.FSDocumentContentRepository)

	documentContentRepositoryAbstract, err := documentContentRepository.New(fsRepository.FSDocumentContentRepositoryConstructorArgs{Logger: logger, Fs: fs})
	if err != nil {
		panic(err)
	}

	documentContentRepository, _ = documentContentRepositoryAbstract.(*fsRepository.FSDocumentContentRepository)

	return documentContentRepository
}

// TODO: Implement proper  hook parsing
// ********************    Manager builders    ********************//
func (m *ConfigManager) NewBookmarkManagerFromConfig(logger *log.Logger, repo repository.BookmarkRepository) libbookmarks.BookmarkManager {
	hooks := bntp.NewHooks[domain.Bookmark]()

	bookmarkManager, err := libbookmarks.NewBookmarkManager(logger, hooks, repo)
	if err != nil {
		panic(err)
	}

	return bookmarkManager
}

func (m *ConfigManager) NewTagsManagerFromConfig(logger *log.Logger, repo repository.TagRepository) libtags.TagManager {
	hooks := new(bntp.Hooks[domain.Tag])
	tagsManager, err := libtags.NewTagmanager(logger, hooks, repo)
	if err != nil {
		panic(err)
	}

	return tagsManager
}
func (m *ConfigManager) NewDocumentManagerFromConfig(logger *log.Logger, repo repository.DocumentRepository) libdocuments.DocumentManager {
	hooks := new(bntp.Hooks[domain.Document])
	documentManager, err := libdocuments.NewDocumentManager(logger, hooks, repo)
	if err != nil {
		panic(err)
	}

	return documentManager
}

func (m *ConfigManager) NewDocumentContentManagerFromConfig(logger *log.Logger, repo repository.DocumentContentRepository) libdocuments.DocumentContentManager {
	hooks := new(bntp.Hooks[string])
	documentContentManager, err := libdocuments.NewDocumentContentRepository(logger, hooks, repo)
	if err != nil {
		panic(err)
	}

	return documentContentManager
}

// **********************    Set up backend    **********************//
func (m *ConfigManager) NewBackendFromConfig() *backend.Backend {
	newBackend := new(backend.Backend)

	db := m.NewDBFromConfig()

	var fs afero.Fs
	if m.FSOverride != nil {
		fs = m.FSOverride
	} else {
		// TODO: Extend config to allow creating custom fs
		fs = afero.NewOsFs()
	}

	tagRepository := m.NewTagsRepositoryFromConfig(m.Logger, db)
	newBackend.BookmarkManager = m.NewBookmarkManagerFromConfig(m.Logger, m.NewBookmarkRepositoryFromConfig(m.Logger, db, tagRepository))
	newBackend.TagManager = m.NewTagsManagerFromConfig(m.Logger, tagRepository)
	newBackend.DocumentManager = m.NewDocumentManagerFromConfig(m.Logger, m.NewDocumentRepositoryFromConfig(m.Logger, db, tagRepository))
	newBackend.DocumentContentManager = m.NewDocumentContentManagerFromConfig(m.Logger, m.NewDocumentContentRepositoryFromConfig(m.Logger, fs))

	newBackend.Marshallers = make(map[string]marshallers.Marshaller)
	newBackend.Unmarshallers = make(map[string]marshallers.Unmarshaller)

	newBackend.Marshallers["json"] = new(marshallers.JsonMarshaller)
	newBackend.Unmarshallers["json"] = new(marshallers.JsonUnmarshaller)

	newBackend.Marshallers["yaml"] = new(marshallers.YamlMarshaller)
	newBackend.Unmarshallers["yaml"] = new(marshallers.YamlUnmarshaller)

	newBackend.Marshallers["csv"] = new(marshallers.CsvMarshaller)
	newBackend.Unmarshallers["csv"] = new(marshallers.CsvUnmarshaller)

	return newBackend
}

//******************************************************************//
//                            Private API                           //
//******************************************************************//

func (m *ConfigManager) addPendingLogMessage(level log.Level, format string, values ...any) {
	m.PendingLogMessage = append(m.PendingLogMessage, newMessage(level, fmt.Sprintf(format, values...)))
}

func (m *ConfigManager) bfsSetDefault(path string, val any) error {
	switch v := val.(type) {
	case map[string]any:
		for k, v := range v {
			newPath := path
			if newPath == "" {
				newPath = k
			} else {
				newPath += "." + k
			}

			m.bfsSetDefault(newPath, v)
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

				err := m.bfsSetDefault(newPath, v)
				if err != nil {
					return err
				}

			}
		} else {
			m.Viper.SetDefault(path, v)
		}
	}

	return nil
}

func (m *ConfigManager) setDefaultsFromStructOrMap(defaults any) error {
	return m.bfsSetDefault("", defaults)
}

func newMessage(level log.Level, msg string) message {
	message := message{}

	message.Level = level
	message.Msg = msg

	return message
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
