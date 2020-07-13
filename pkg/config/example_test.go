// Copyright (c) 2020 The Truffls Authors.
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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/foodarchive/truffls/pkg/config"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

type (
	ServerConfig struct {
		Addr string
		Port uint
	}

	Config struct {
		AppName string `config:"app_name"`
		Env     string
		Debug   bool
		Server  ServerConfig `config:"http_server"`
	}
)

var (
	configBytes = []byte(`
app_name: foo
debug: true
http_server:
  addr: localhost
  port: 3000`)
)

func Example() {
	configPath := fmt.Sprintf("%s/config_sample.yml", filepath.Dir("."))
	if err := ioutil.WriteFile(configPath, configBytes, 0644); err != nil {
		log.Fatal(errors.Wrap(err, "fail to write config file"))
	}

	defer func() {
		_ = os.Remove(configPath)
	}()

	_ = pflag.String("env", "development", "setup app env")
	pflag.Parse()

	_ = os.Setenv("SIMPLE_APP_NAME", "simple")
	defer func() {
		_ = os.Unsetenv("SIMPLE_APP_NAME")
	}()

	config.Load("simple", configPath)
	config.BindFlags(pflag.Lookup("env"))

	var c Config
	if err := config.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)
	// Output:
	// {AppName:simple Env:development Debug:true Server:{Addr:localhost Port:3000}}
}
