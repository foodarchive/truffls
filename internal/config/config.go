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

package config

import (
	pkgconfig "github.com/foodarchive/truffls/pkg/config"
	pkglog "github.com/foodarchive/truffls/pkg/log"
)

type server struct {
	Host    string
	Port    string
	GinMode string
}

type config struct {
	Debug  bool
	Server server
	Log    pkglog.Config
}

var (
	// AppName is dynamically set by the toolchain or overridden by the Makefile.
	AppName = "truffls"
	// Version is dynamically set by the toolchain or overridden by the Makefile.
	Version = "DEV"
	// BuildDate is dynamically set at build time in the Makefile.
	BuildDate = "2020-07-01" // YYYY-MM-DD
	// Debug turn on/of debugging mode.
	Debug bool
	// Server configuration.
	Server server
	// Log configuration for logging package.
	Log pkglog.Config
)

// Load config file, return error when failed to load config file
// or unmarshaling config struct.
func Load(configFile string) error {
	if err := pkgconfig.Load(AppName, configFile); err != nil {
		return err
	}

	var c config
	if err := pkgconfig.Unmarshal(&c); err != nil {
		return err
	}

	Debug = c.Debug
	Server = c.Server
	Log = c.Log

	if Debug {
		Log.Level = "debug"
	}

	return nil
}
