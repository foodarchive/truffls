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

// Config logging configuration.
type Config struct {
	// Level log level config.
	//
	// Supported values: `trace`, `debug`, `info`, `warn`, `fatal`, `panic`.
	// Default value "info".
	Level            string

	// Output log output.
	//
	// Supported values: `console`, `stderr`, `stdout`.
	// For `console` print log with pretty output,
	// `stderr` and `stdout` print log with JSON format.
	Output           string

	// EnableStackTrace set this to true to enable stack stacktrace.
	EnableStackTrace bool

	// ShowCaller if enabled, print log caller filename and line number.
	ShowCaller       bool
}
