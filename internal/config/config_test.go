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

	. "github.com/foodarchive/truffls/internal/config"
	pkgConfig "github.com/foodarchive/truffls/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	pkgConfig.Load("truffls", "./testdata/config_test.yml")
	c, err := New()

	assert.NoError(t, err)
	assert.True(t, c.Debug)
}
