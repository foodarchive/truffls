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

package log_test

import (
	stdLog "log"
	"time"

	"github.com/foodarchive/truffls/pkg/log"
	"github.com/rs/zerolog"
)

func setup() {
	// Use static timestamp for testing.
	zerolog.TimestampFunc = func() time.Time {
		return time.Date(2020, 7, 18, 21, 1, 05, 0, time.UTC)
	}

	zerolog.CallerMarshalFunc = func(_ string, _ int) string {
		return "/app/controller/root.go:113"
	}
}

func Example() {
	setup()
	log.Init(log.Config{})

	log.Info().Msg("hello world")
	// Output:
	// {"level":"info","time":"2020-07-18T21:01:05Z","message":"hello world"}
}

func Example_StdLog() {
	setup()
	log.Init(log.Config{})

	stdLog.Printf("hello from standard log")
	// Output:
	// {"time":"2020-07-18T21:01:05Z","level":"warn","message":"hello from standard log"}
}

func Example_WithCaller() {
	setup()
	log.Init(log.Config{ShowCaller: true, Level: "debug"})

	log.Debug().Msg("something went wrong")
	// Output:
	// {"level":"debug","time":"2020-07-18T21:01:05Z","caller":"/app/controller/root.go:113","message":"something went wrong"}
}
