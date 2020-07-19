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
	"time"

	"github.com/foodarchive/truffls/pkg/config"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type configTest struct {
	Foo     string `config:"bar"`
	Elapsed time.Duration
	Arr     []string
}

func TestLoad(t *testing.T) {
	testCases := []struct {
		namespace string
		cfgFile   string
		assertion assert.ErrorAssertionFunc
		name      string
	}{
		{
			namespace: "",
			assertion: assert.Error,
			name:      "ErrorEmptyNamespace",
		},
		{
			namespace: "foo",
			cfgFile:   "",
			assertion: assert.Error,
			name:      "ErrorEmptyConfigFile",
		},
		{
			namespace: "foo",
			cfgFile:   "noop.yml",
			assertion: assert.Error,
			name:      "ErrorNonExistingConfigFile",
		},
		{
			namespace: "foo",
			cfgFile:   "./testdata/config_test.yml",
			assertion: assert.NoError,
			name:      "SuccessValidConfigFile",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.assertion(t, config.Load(tc.namespace, tc.cfgFile))
		})
	}
}

func TestUnmarshal(t *testing.T) {
	require.NoError(t, config.Load("config_test", "./testdata/config_test.yml"))

	var c configTest
	assert.NoError(t, config.Unmarshal(&c))
	assert.Equal(t, c.Foo, "baz")
	assert.Equal(t, c.Elapsed, 5*time.Second)
	assert.Equal(t, c.Arr, []string{"hello", "world"})
}

func TestBindFlags(t *testing.T) {
	require.NoError(t, config.Load("config_test", "./testdata/config_test.yml"))

	f := pflag.NewFlagSet("config_flag_test", pflag.ContinueOnError)
	f.String("bar", "", "")
	f.String("elapsed", "", "")

	require.NoError(t, f.Set("bar", "yolo"))
	require.NoError(t, f.Set("elapsed", "10s"))

	assert.NoError(t, config.BindFlags(f.Lookup("bar"), f.Lookup("elapsed")))

	var c configTest
	require.NoError(t, config.Unmarshal(&c))
	assert.Equal(t, c.Foo, "yolo")
	assert.Equal(t, c.Elapsed, 10*time.Second)
}
