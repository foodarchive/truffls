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
	"log"

	"github.com/foodarchive/truffls/internal/config"
	pkgConfig "github.com/foodarchive/truffls/pkg/config"
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
		if err := pkgConfig.Load(config.AppName, cfgFile); err != nil {
			log.Fatal(err)
		}
	})

	pf := rootCmd.PersistentFlags()
	pf.StringVar(&cfgFile, "config", "", "config filepath")
	pf.Bool("debug", false, "debugging mode")

	if err := pkgConfig.BindFlags(pf.Lookup("config"), pf.Lookup("debug")); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
