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
	"net/http"
)

// Option is a function to configure Server.
type Option func(h *Server)

func (fn Option) apply(h *Server) {
	fn(h)
}

// WithAddr sets server address.
func WithAddr(host, port string) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort(host, port)
	}
}

// WithHandler assign server http handler.
func WithHandler(handler http.Handler) Option {
	return func(s *Server) {
		s.server.Handler = handler
	}
}

// WithCertFile sets the location of the certificate and matching private key files.
func WithCertFile(cert, key string) Option {
	return func(s *Server) {
		s.tls.CertFile = cert
		s.tls.KeyFile = key
	}
}

// WithCert sets the certificate and matching private key.
func WithCert(cert, key []byte) Option {
	return func(s *Server) {
		s.tls.Cert = cert
		s.tls.Key = key
	}
}

// WithAutoTLS sets host and cacheDir for auto-TLS.
func WithAutoTLS(host, cacheDir string) Option {
	return func(s *Server) {
		s.autoTLS.Host = host
		s.autoTLS.CacheDir = cacheDir
	}
}
