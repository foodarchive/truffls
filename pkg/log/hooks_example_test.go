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

	"github.com/foodarchive/truffls/pkg/log"
)

func ExampleNoLevelDebugHook_Run() {
	setup()
	log.Init(log.Config{})

	hook := log.Log.Hook(log.NoLevelDebugHook{})
	stdLog.SetFlags(0)
	stdLog.SetOutput(hook)

	stdLog.Print("debugging logs...")
	// Output:
	// {"time":"2020-07-18T21:01:05Z","level":"debug","message":"debugging logs..."}
}

func ExampleNoLevelErrorHook_Run() {
	setup()
	log.Init(log.Config{})

	hook := log.Log.Hook(log.NoLevelErrorHook{})
	stdLog.SetFlags(0)
	stdLog.SetOutput(hook)

	stdLog.Print("error logs...")
	// Output:
	// {"time":"2020-07-18T21:01:05Z","level":"error","message":"error logs..."}
}
