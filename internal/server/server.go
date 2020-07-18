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
	"net"

	"github.com/foodarchive/truffls/internal/config"
	"github.com/foodarchive/truffls/internal/handler"
	pkgServer "github.com/foodarchive/truffls/pkg/server"
	"github.com/gin-gonic/gin"
)

var (
	cfg config.Config
)

// Start starts HTTP server.
func Start() (err error) {
	if cfg, err = config.New(); err != nil {
		return
	}

	srv := pkgServer.New(
		pkgServer.WithAddr(net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)),
		pkgServer.WithHandler(router()),
	)

	return srv.Start()
}

func router() *gin.Engine {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()
	g.Use(gin.Recovery())

	g.GET("/", handler.Root)
	return g
}
