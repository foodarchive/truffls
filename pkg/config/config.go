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

// Package config is a https://github.com/spf13/viper wrapper
// with some convention.
package config

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type (
	// Flag is pflag.Flag alias.
	Flag = pflag.Flag
)

// Load loading configuration from yaml file with namespace.
func Load(namespace, filename string) {
	if namespace == "" {
		log.Fatal("config namespace must be provided")
	}

	if filename != "" {
		viper.SetConfigFile(filename)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/" + namespace)
		viper.AddConfigPath("$HOME/." + namespace)
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix(namespace)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

// Unmarshal unmarshals the config into a Struct.
func Unmarshal(v interface{}) error {
	return viper.Unmarshal(v, func(c *mapstructure.DecoderConfig) {
		c.TagName = "config"
		c.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		)
	})
}

// BindFlags binds multiple keys to pflag.Flag.
func BindFlags(flags ...*Flag) {
	for _, flag := range flags {
		_ = viper.BindPFlag(flag.Name, flag)
	}
}
