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
	"path/filepath"
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
	"github.com/JonasMuehlmann/goaoi/functional"
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
	CacheDir           string
	TempDir            string
	ConfigFileBaseName string
	ConfigSearchPaths  []string
}

func NewConfigManager(stderr io.Writer, dbOverride *sql.DB, fsOverride afero.Fs) (m *ConfigManager, err error) {
	m = &ConfigManager{Viper: viper.New(), ConfigFileBaseName: "bntp"}

	var fs afero.Fs
	if m.FSOverride != nil {
		fs = m.FSOverride
	} else {
		// TODO: Extend config to allow creating custom fs
		fs = afero.NewOsFs()
	}

	//****************    Make sure config dir exists    ***************//
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return
	}

	tempDir := os.TempDir()
	m.TempDir = path.Join(tempDir, "bntp")
	m.CacheDir = path.Join(userCacheDir, "bntp")
	m.ConfigDir = path.Join(userConfigDir, "bntp")

	doesTempDirExist, err := afero.Exists(fs, m.TempDir)
	if err != nil {
		return
	}

	if !doesTempDirExist {
		err = fs.MkdirAll(m.TempDir, 0o744)
		if err != nil {
			return
		}
	}

	doesCacheDirExist, err := afero.Exists(fs, m.CacheDir)
	if err != nil {
		return
	}

	if !doesCacheDirExist {
		err = fs.MkdirAll(m.CacheDir, 0o744)
		if err != nil {
			return
		}
	}

	doesConfigDirExist, err := afero.Exists(fs, m.ConfigDir)
	if err != nil {
		return
	}

	if !doesConfigDirExist {
		err = fs.MkdirAll(m.ConfigDir, 0o744)
		if err != nil {
			return
		}
	}

	//********************    Set default values    ********************//
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
		m.addPendingLogMessage(log.ErrorLevel, "Error setting default values: %v", err)
	}

	m.Viper.SetEnvPrefix("BNTP")
	m.Viper.SetConfigName(m.ConfigFileBaseName)
	m.Viper.AutomaticEnv() // read in environment variables that match

	err = m.Viper.ReadInConfig()
	if err != nil && errors.Is(err, &viper.ConfigFileNotFoundError{}) {
		m.addPendingLogMessage(log.ErrorLevel, "Error reading config: %v", err)
	}

	logFile := m.Viper.GetString(LogFile)

	doesLogFileExist, err := afero.Exists(fs, logFile)
	if err != nil {
		return
	}

	if !doesLogFileExist {
		err = afero.WriteFile(fs, logFile, []byte("foo"), 0o644)
		if err != nil {
			return
		}
	}

	//*********************    Validate settings    ********************//
	var config Config

	err = m.Viper.Unmarshal(&config)
	if err != nil {
		return
	}

	var consoleLogLevel log.Level
	var fileLogLevel log.Level

	err = ConfigValidator.Struct(config)
	if err != nil {
		consoleLogLevel = log.ErrorLevel
		fileLogLevel = log.InfoLevel

		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			m.addPendingLogMessage(log.ErrorLevel, "Error processing setting %v: %v", strcase.SnakeCase(e.Field()), formatValidationError(e))
		}
	} else {
		consoleLogLevel, _ = log.ParseLevel(m.Viper.GetString(ConsoleLogLevel))
		fileLogLevel, _ = log.ParseLevel(m.Viper.GetString(FileLogLevel))
	}
	// ******************** Log pending messages ********************* //
	m.Logger = helper.NewDefaultLogger(logFile, consoleLogLevel, fileLogLevel, stderr)

	for _, message := range m.PendingLogMessage {
		m.Logger.Log(message.Level, message.Msg)
	}

	if m.PassedConfigPath != "" {
		m.Logger.Infof("Using config file %v", m.Viper.ConfigFileUsed())
	} else {
		m.Logger.Info("Using internal configuration")
	}

	return
}

func (m *ConfigManager) NewDBFromConfig() (db *sql.DB, err error) {
	if m.DBOverride != nil {
		return m.DBOverride, nil
	}

	dataSource := m.Viper.GetString(DB_DataSource)
	args := m.Viper.GetStringSlice(DB_Args)

	openArgs := dataSource
	if len(args) > 0 {
		openArgs += "?"

		openArgs += strings.Join(args, "&")
	}

	db, err = sql.Open(m.Viper.GetString(DB_Driver), openArgs)
	if err != nil {
		return
	}

	return
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
		LogFile:         path.Join(m.CacheDir, "bntp.log"),
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
func (m *ConfigManager) NewBookmarkRepositoryFromConfig(logger *log.Logger, repoDB *sql.DB, tagRepository repository.TagRepository) (repo repository.BookmarkRepository, err error) {
	repo = new(sqlite3Repository.Sqlite3BookmarkRepository)

	bookmarkRepositoryAbstract, err := repo.New(sqlite3Repository.Sqlite3BookmarkRepositoryConstructorArgs{Logger: logger, DB: repoDB, TagRepository: tagRepository})
	if err != nil {
		return
	}

	repo, _ = bookmarkRepositoryAbstract.(*sqlite3Repository.Sqlite3BookmarkRepository)

	return
}

func (m *ConfigManager) NewTagsRepositoryFromConfig(logger *log.Logger, repoDB *sql.DB) (repo repository.TagRepository, err error) {
	repo = new(sqlite3Repository.Sqlite3TagRepository)

	tagsRepositoryAbstract, err := repo.New(sqlite3Repository.Sqlite3TagRepositoryConstructorArgs{Logger: logger, DB: repoDB})
	if err != nil {
		return
	}

	repo, _ = tagsRepositoryAbstract.(*sqlite3Repository.Sqlite3TagRepository)

	return
}
func (m *ConfigManager) NewDocumentRepositoryFromConfig(logger *log.Logger, repoDB *sql.DB, tagRepository repository.TagRepository) (repo repository.DocumentRepository, err error) {
	repo = new(sqlite3Repository.Sqlite3DocumentRepository)

	documentRepositoryAbstract, err := repo.New(sqlite3Repository.Sqlite3DocumentRepositoryConstructorArgs{Logger: logger, DB: repoDB, TagRepository: tagRepository})
	if err != nil {
		return
	}

	repo, _ = documentRepositoryAbstract.(*sqlite3Repository.Sqlite3DocumentRepository)

	return
}

// TODO: Allow non-fs content repositories
func (m *ConfigManager) NewDocumentContentRepositoryFromConfig(logger *log.Logger, fs afero.Fs) (repo repository.DocumentContentRepository, err error) {
	repo = new(fsRepository.FSDocumentContentRepository)

	documentContentRepositoryAbstract, err := repo.New(fsRepository.FSDocumentContentRepositoryConstructorArgs{Logger: logger, Fs: fs})
	if err != nil {
		return
	}

	repo, _ = documentContentRepositoryAbstract.(*fsRepository.FSDocumentContentRepository)

	return
}

// TODO: Implement proper  hook parsing
// ********************    Manager builders    ********************//
func (m *ConfigManager) NewBookmarkManagerFromConfig(logger *log.Logger, repo repository.BookmarkRepository) (manager libbookmarks.BookmarkManager, err error) {
	hooks := bntp.NewHooks[domain.Bookmark]()

	manager, err = libbookmarks.NewBookmarkManager(logger, hooks, repo)
	if err != nil {
		return
	}

	return
}

func (m *ConfigManager) NewTagsManagerFromConfig(logger *log.Logger, repo repository.TagRepository) (manager libtags.TagManager, err error) {
	hooks := new(bntp.Hooks[domain.Tag])
	manager, err = libtags.NewTagmanager(logger, hooks, repo)
	if err != nil {
		return
	}

	return
}
func (m *ConfigManager) NewDocumentManagerFromConfig(logger *log.Logger, repo repository.DocumentRepository) (manager libdocuments.DocumentManager, err error) {
	hooks := new(bntp.Hooks[domain.Document])
	manager, err = libdocuments.NewDocumentManager(logger, hooks, repo)
	if err != nil {
		return
	}

	return
}

func (m *ConfigManager) NewDocumentContentManagerFromConfig(logger *log.Logger, repo repository.DocumentContentRepository) (manager libdocuments.DocumentContentManager, err error) {
	hooks := new(bntp.Hooks[string])
	manager, err = libdocuments.NewDocumentContentManager(logger, hooks, repo)
	if err != nil {
		return
	}

	return
}

// **********************    Set up backend    **********************//
func (m *ConfigManager) NewBackendFromConfig() (newBackend *backend.Backend, err error) {
	newBackend = new(backend.Backend)

	db, err := m.NewDBFromConfig()
	if err != nil {
		return
	}

	var fs afero.Fs
	if m.FSOverride != nil {
		fs = m.FSOverride
	} else {
		// TODO: Extend config to allow creating custom fs
		fs = afero.NewOsFs()
	}

	var tagRepository repository.TagRepository
	var documentRepository repository.DocumentRepository
	var bookmarkRepository repository.BookmarkRepository
	var documentContentRepository repository.DocumentContentRepository

	//********************    Create repositories    *******************//
	tagRepository, err = m.NewTagsRepositoryFromConfig(m.Logger, db)
	if err != nil {
		return
	}
	bookmarkRepository, err = m.NewBookmarkRepositoryFromConfig(m.Logger, db, tagRepository)
	if err != nil {
		return
	}
	documentRepository, err = m.NewDocumentRepositoryFromConfig(m.Logger, db, tagRepository)
	if err != nil {
		return
	}

	documentContentRepository, err = m.NewDocumentContentRepositoryFromConfig(m.Logger, fs)
	if err != nil {
		return
	}

	//**********************    Create Managers    *********************//
	newBackend.BookmarkManager, err = m.NewBookmarkManagerFromConfig(m.Logger, bookmarkRepository)
	if err != nil {
		return
	}

	newBackend.TagManager, err = m.NewTagsManagerFromConfig(m.Logger, tagRepository)
	if err != nil {
		return
	}

	newBackend.DocumentManager, err = m.NewDocumentManagerFromConfig(m.Logger, documentRepository)
	if err != nil {
		return
	}

	newBackend.DocumentContentManager, err = m.NewDocumentContentManagerFromConfig(m.Logger, documentContentRepository)
	if err != nil {
		return
	}

	newBackend.Marshallers = make(map[string]marshallers.Marshaller)
	newBackend.Unmarshallers = make(map[string]marshallers.Unmarshaller)

	newBackend.Marshallers["json"] = new(marshallers.JsonMarshaller)
	newBackend.Unmarshallers["json"] = new(marshallers.JsonUnmarshaller)

	newBackend.Marshallers["yaml"] = new(marshallers.YamlMarshaller)
	newBackend.Unmarshallers["yaml"] = new(marshallers.YamlUnmarshaller)

	newBackend.Marshallers["csv"] = new(marshallers.CsvMarshaller)
	newBackend.Unmarshallers["csv"] = new(marshallers.CsvUnmarshaller)

	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	parentDirs := strings.Split(cwd, string(os.PathSeparator))

	iProjectRoot, err := goaoi.FindIfSlice(parentDirs, functional.AreEqualPartial("bntp.go"))
	if err != nil {
		return
	}

	schemaDirPath := string(os.PathSeparator) + filepath.Join(filepath.Join(parentDirs[:iProjectRoot+1]...), "schema")

	schemaFiles, err := os.ReadDir(schemaDirPath)
	if err != nil {
		return
	}

	newBackend.DBProviderSchemas = make(map[string]string)

	for _, schemaFile := range schemaFiles {
		schemaFileName := schemaFile.Name()
		if schemaFile.Type().IsDir() || strings.Contains(schemaFileName, "test") {
			continue
		}

		providerName := strings.Split(schemaFileName, "_")[1]
		providerName = strings.Split(providerName, ".")[0]

		newBackend.DBProviderSchemas[providerName] = filepath.Join(schemaDirPath, schemaFileName)
	}

	// newBackend.DBProviderSchemas["sqlite"] = os.ReadFile()

	return
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
