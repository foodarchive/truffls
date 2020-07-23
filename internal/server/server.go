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

package server

import (
	"net/http"

	"github.com/foodarchive/truffls/internal/config"
	"github.com/foodarchive/truffls/internal/server/handler"
	pkgserver "github.com/foodarchive/truffls/pkg/server"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Start starts HTTP server.
func Start() error {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.Path("/").Methods(http.MethodGet).HandlerFunc(handler.Root)

	n := negroni.New()
	recoveryMw := negroni.NewRecovery()
	if config.Debug {
		recoveryMw.PrintStack = true
		recoveryMw.StackSize = 1 << 20
	}

	n.Use(recoveryMw)
	n.UseHandler(r)

	srv := pkgserver.New(n,
		pkgserver.WithAddr(config.Server.Host, config.Server.Port),
		pkgserver.WithCertFile(config.Server.TLS.CertFile, config.Server.TLS.KeyFile),
		pkgserver.WithAutoTLS(config.Server.AutoTLS.Host, config.Server.AutoTLS.CacheDir),
	)

	var err error
	{
		switch {
		case config.Server.TLS.Enabled:
			err = srv.StartTLS()
		case config.Server.AutoTLS.Enabled:
			err = srv.StartAutoTLS()
		default:
			err = srv.Start()
		}

		if err != nil {
			return err
		}
	}

	return srv.Shutdown()
}
