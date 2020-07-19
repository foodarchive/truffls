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
	"os"
	"testing"

	"github.com/foodarchive/truffls/pkg/log"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	testCases := []struct {
		config   log.Config
		expected log.Logger
		name     string
	}{
		{
			config:   log.Config{},
			expected: zerolog.New(os.Stdout).With().Timestamp().Logger(),
			name:     "DefaultConfig",
		},
		{
			config:   log.Config{Out: "stderr"},
			expected: zerolog.New(os.Stderr).With().Timestamp().Logger(),
			name:     "StderrOutput",
		},
		{
			config:   log.Config{Out: "console"},
			expected: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger(),
			name:     "ConsoleOutput",
		},
		{
			config:   log.Config{Caller: true},
			expected: zerolog.New(os.Stdout).With().Timestamp().Caller().Logger(),
			name:     "WithCaller",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			log.Init(tc.config)
			assert.Equal(t, log.Log, tc.expected)
		})
	}
}
