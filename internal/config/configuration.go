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
	"fmt"
	"os"
	"path"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/stoewer/go-strcase"

	log "github.com/sirupsen/logrus"
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

var DefaultSettings = map[string]any{
	LOG_FILE:          path.Join(ConfigDir, "bntp.log"),
	CONSOLE_LOG_LEVEL: log.ErrorLevel,
	FILE_LOG_LEVEL:    log.InfoLevel,
}

var pendingLogMessage []message

func addPendingLogMessage(level log.Level, format string, values ...any) {
	pendingLogMessage = append(pendingLogMessage, newMessage(level, fmt.Sprintf(format, values...)))
}

func formatValidationError(e validator.FieldError) string {
	var message string

	// TODO: Add test to validate if all used validator tags are handled here
	switch e.Tag() {
	case VALIDATOR_LOGRUS_LOG_LEVEL:
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

func InitConfig() {

	pendingLogMessage = make([]message, 0, 5)

	if PassedConfigPath != "" {
		viper.SetConfigFile(PassedConfigPath)
	}

	goaoi.ForeachSliceUnsafe(ConfigSearchPaths, viper.AddConfigPath)

	for k, v := range DefaultSettings {
		viper.SetDefault(k, v)
	}

	viper.SetEnvPrefix("BNTP")
	viper.SetConfigName(ConfigFileBaseName)
	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()
	if err != nil {
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
		consoleLogLevel, _ = log.ParseLevel(viper.GetString(CONSOLE_LOG_LEVEL))
		fileLogLevel, _ = log.ParseLevel(viper.GetString(FILE_LOG_LEVEL))
	}
	// ******************** Log pending messages ********************* //

	helper.Logger = *helper.NewDefaultLogger(viper.GetString(LOG_FILE), consoleLogLevel, fileLogLevel)

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
