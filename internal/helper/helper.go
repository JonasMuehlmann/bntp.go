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

package helper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"runtime"
	"time"

	"github.com/JonasMuehlmann/goaoi"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

const (
	// NOTE: This is the only format time marshaling supports, it is not configurable...
	DateFormat           string = time.RFC3339
	LogMessageEmptyInput        = "Returning early after receiving empty input"
)

var (
	EmptyInputError             = errors.New("empty input")
	NonExistentPrimaryDataError = errors.New("the primary data to work with does not exist")
	NopUpdaterError             = errors.New("The updater will leave the data unchanged")
)

type NilInputError struct {
	BadFieldOrParameter string
}

func (err NilInputError) Error() string {
	if err.BadFieldOrParameter == "" {
		return "Input contains a nil pointer"
	}
	return "Input contains a nil pointer in parameter or struct field " + err.BadFieldOrParameter
}

type IneffectiveOperationError struct {
	Inner error
}

func (err IneffectiveOperationError) Error() string {
	return fmt.Sprintf("The operation had no effect: %v", err.Inner)
}

func (err IneffectiveOperationError) Unwrap() error {
	return err.Inner
}

func (err IneffectiveOperationError) Is(other error) bool {
	var thisZero IneffectiveOperationError
	return other.(IneffectiveOperationError) != thisZero
}

func (err IneffectiveOperationError) As(target any) bool {
	var thisZero IneffectiveOperationError
	isTarget := target.(IneffectiveOperationError) != thisZero

	if isTarget {
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
	}

	return isTarget
}

type DuplicateInsertionError struct {
	Inner error
}

func (err DuplicateInsertionError) Error() string {
	return fmt.Sprintf("The operation would insert a duplicate: %v", err.Inner)
}

func (err DuplicateInsertionError) Unwrap() error {
	return err.Inner
}

func (err DuplicateInsertionError) Is(other error) bool {
	var thisZero DuplicateInsertionError
	return other.(DuplicateInsertionError) != thisZero
}

func (err DuplicateInsertionError) As(target any) bool {
	var thisZero DuplicateInsertionError
	isTarget := target.(DuplicateInsertionError) != thisZero

	if isTarget {
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
	}

	return isTarget
}

func NewDefaultLogger(logFile string, consoleLogLevel log.Level, fileLogLevel log.Level) *log.Logger {
	callerPrettyfier := func(f *runtime.Frame) (string, string) {
		filename := path.Base(f.File)

		return fmt.Sprintf("in function %s()", path.Base(f.Function)), fmt.Sprintf("%s:%d", filename, f.Line)
	}

	formatter := &log.TextFormatter{
		FullTimestamp:    true,
		CallerPrettyfier: callerPrettyfier,
	}
	consoleFormatter := &log.TextFormatter{
		FullTimestamp:    true,
		ForceColors:      true,
		CallerPrettyfier: callerPrettyfier,
	}

	logFileHandle, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o755)
	if err != nil {
		panic("Failed to set up file logger with path " + logFile)
	}

	mainLogger := log.New()
	mainLogger.SetFormatter(formatter)
	mainLogger.SetOutput(ioutil.Discard)
	mainLogger.SetReportCaller(true)

	// Log levels are ordered from most to least serve
	consoleLogLevels, _ := goaoi.TakeWhileSlice(log.AllLevels, goaoi.IsLessThanEqualPartial(consoleLogLevel))
	fileLoglevels, _ := goaoi.TakeWhileSlice(log.AllLevels, goaoi.IsLessThanEqualPartial(fileLogLevel))

	mainLogger.AddHook(&ConsoleLoggerHook{
		Formatter: consoleFormatter,
		LogLevels: consoleLogLevels,
	})
	mainLogger.AddHook(&writer.Hook{
		Writer:    logFileHandle,
		LogLevels: fileLoglevels,
	})

	return mainLogger
}

type ConsoleLoggerHook struct {
	Formatter log.Formatter
	LogLevels []log.Level
}

func (hook *ConsoleLoggerHook) Fire(entry *log.Entry) error {
	formatted, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = os.Stderr.Write(formatted)

	return err
}

func (hook *ConsoleLoggerHook) Levels() []log.Level {
	return hook.LogLevels
}
