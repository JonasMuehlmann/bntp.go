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
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/JonasMuehlmann/goaoi"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

const DateFormat string = "helper.DateFormat"

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
