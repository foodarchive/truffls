// Copyright The Truffls Contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"io"
	stdlog "log"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const (
	outConsole = "console"
	outStderr  = "stderr"
)

type (
	// Event zerolog.Event alias.
	Event = zerolog.Event
	// Hook zerolog.Hook alias.
	Hook = zerolog.Hook
	// Logger zerolog.Logger alias.
	Logger = zerolog.Logger
	// Level zerolog.Level alias.
	Level = zerolog.Level
)

var (
	Log Logger
)

// Init initialize logger based on Config.
func Init(config Config) {
	level := zerolog.InfoLevel
	plevel, err := zerolog.ParseLevel(config.Level)
	if err == nil && plevel != zerolog.NoLevel {
		level = plevel
	}
	zerolog.SetGlobalLevel(level)

	// Maybe add more write, eg: syslog, journald, etc,
	// or supports multiple writer.
	var w io.Writer
	switch strings.ToLower(config.Out) {
	case outConsole:
		w = zerolog.ConsoleWriter{Out: os.Stdout}
	case outStderr:
		w = os.Stderr
	default:
		w = os.Stdout
	}

	if config.StackTrace {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	Log = zerolog.New(w).With().Timestamp().Logger()
	if config.Caller {
		Log = Log.With().Caller().Logger()
	}

	// replace standard logger with zerolog
	hook := Log.Hook(NoLevelWarnHook{})
	stdlog.SetFlags(0)
	stdlog.SetOutput(hook)
}

// Debug starts logging with debug level.
func Debug() *Event {
	return Log.Debug()
}

// Info starts logging with info level.
func Info() *Event {
	return Log.Info()
}

// Error starts logging with error level.
func Error() *Event {
	return Log.Error()
}

// Info starts logging with fatal level.
// Note: under the hood it will call os.Exit(1).
func Fatal() *Event {
	return Log.Fatal()
}

// WithHook returns a logger with the h Hook.
func WithHook(h Hook) Logger {
	return Log.Hook(h)
}
