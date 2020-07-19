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
	"log"

	pkgConfig "github.com/foodarchive/truffls/pkg/config"
	pkgLog "github.com/foodarchive/truffls/pkg/log"
)

type server struct {
	Host    string
	Port    string
	GinMode string
}

type config struct {
	Debug  bool
	Server server
	Log    pkgLog.Config
}

var (
	// AppName is dynamically set by the toolchain or overridden by the Makefile.
	AppName = "truffls"

	// Version is dynamically set by the toolchain or overridden by the Makefile.
	Version = "DEV"

	// BuildDate is dynamically set at build time in the Makefile.
	BuildDate = "2020-07-01" // YYYY-MM-DD

	Debug  bool
	Server server
	Log    pkgLog.Config
)

func Init(configFile string) {
	if err := pkgConfig.Load(AppName, configFile); err != nil {
		log.Fatal(err)
	}

	var c config
	if err := pkgConfig.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}

	Debug = c.Debug
	Server = c.Server
	Log = c.Log

	if Debug {
		Log.Level = "debug"
	}
}
