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

package cmd

import (
	stdLog "log"

	"github.com/foodarchive/truffls/internal/config"
	pkgCfg "github.com/foodarchive/truffls/pkg/config"
	"github.com/foodarchive/truffls/pkg/log"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use: config.AppName,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
)

func init() {
	cobra.OnInitialize(func() {
		config.Init(cfgFile)
		log.Init(config.Log)
	})

	pf := rootCmd.PersistentFlags()
	pf.StringVar(&cfgFile, "config", "./config.yml", "config filepath")
	pf.BoolVar(&config.Debug, "debug", false, "debugging mode")

	err := pkgCfg.BindFlags(pf.Lookup("config"), pf.Lookup("debug"))
	if err != nil {
		stdLog.Fatal(err)
	}
}

// Execute run commandline, exit with status 1 if there's an error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		stdLog.Fatal(err)
	}
}
