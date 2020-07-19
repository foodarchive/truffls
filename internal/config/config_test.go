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

package config_test

import (
	"testing"

	"github.com/foodarchive/truffls/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	config.Load("./testdata/config_test.yml")

	assert.True(t, config.Debug)
	assert.Equal(t, config.Server.Host, "www.example.com")
	assert.Equal(t, config.Server.Port, "3000")
}
