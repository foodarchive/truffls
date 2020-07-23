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
	"github.com/rs/zerolog"
)

// NoLevelWarnHook replace no level with warn log level.
type NoLevelWarnHook struct{}

// Run add warn level to the log message when log level not provided.
func (h NoLevelWarnHook) Run(e *Event, level Level, _ string) {
	if level == zerolog.NoLevel {
		e.Str(zerolog.LevelFieldName, zerolog.WarnLevel.String())
	}
}

// NoLevelDebugHook replace no level with warn log level.
type NoLevelDebugHook struct{}

// Run add debug level to the log message when log level not provided.
func (h NoLevelDebugHook) Run(e *Event, level Level, _ string) {
	if level == zerolog.NoLevel {
		e.Str(zerolog.LevelFieldName, zerolog.DebugLevel.String())
	}
}

// NoLevelErrorHook replace no level with warn log level.
type NoLevelErrorHook struct{}

// Run add error level to the log message when log level not provided.
func (h NoLevelErrorHook) Run(e *Event, level Level, _ string) {
	if level == zerolog.NoLevel {
		e.Str(zerolog.LevelFieldName, zerolog.ErrorLevel.String())
	}
}
