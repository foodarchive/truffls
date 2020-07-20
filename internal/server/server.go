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
	"github.com/foodarchive/truffls/internal/config"
	"github.com/foodarchive/truffls/internal/server/handler"
	"github.com/foodarchive/truffls/pkg/log"
	pkgserver "github.com/foodarchive/truffls/pkg/server"
	"github.com/gin-gonic/gin"
)

// Start starts HTTP server.
func Start() error {
	srv := pkgserver.New(
		pkgserver.WithAddr(config.Server.Host, config.Server.Port),
		pkgserver.WithHandler(router()),
		pkgserver.WithCertFile(config.Server.TLS.CertFile, config.Server.TLS.KeyFile),
		pkgserver.WithAutoTLS(config.Server.AutoTLS.Host, config.Server.AutoTLS.CacheDir),
	)

	switch {
	case config.Server.TLS.Enabled:
		return srv.StartTLS()
	case config.Server.AutoTLS.Enabled:
		return srv.StartAutoTLS()
	default:
		return srv.Start()
	}
}

func router() *gin.Engine {
	gin.SetMode(config.Server.GinMode)
	gin.DefaultWriter = log.WithHook(log.NoLevelDebugHook{})
	gin.DefaultErrorWriter = log.WithHook(log.NoLevelErrorHook{})

	r := gin.New()
	r.RemoveExtraSlash = true

	r.Use(gin.Recovery())

	r.GET("/", handler.Root)
	return r
}
