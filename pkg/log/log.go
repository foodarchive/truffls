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
	stdLog "log"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const (
	outConsole = "console"
	outStdErr  = "stderr"
)

var (
	l zerolog.Logger
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
	switch strings.ToLower(config.Output) {
	case outConsole:
		w = zerolog.ConsoleWriter{Out: os.Stdout}
	case outStdErr:
		w = os.Stderr
	default:
		w = os.Stdout
	}

	if config.EnableStackTrace {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	l = zerolog.New(w).With().Timestamp().Logger()
	if config.ShowCaller {
		l = l.With().Caller().Logger()
	}

	// replace standard logger with zerolog
	hook := l.Hook(noLevelHook{})
	stdLog.SetFlags(0)
	stdLog.SetOutput(hook)
}

// Debug starts logging with debug level.
func Debug() *zerolog.Event {
	return l.Debug()
}

// Info starts logging with info level.
func Info() *zerolog.Event {
	return l.Info()
}

// Error starts logging with error level.
func Error() *zerolog.Event {
	return l.Error()
}

// Info starts logging with fatal level.
func Fatal() *zerolog.Event {
	return l.Fatal()
}
