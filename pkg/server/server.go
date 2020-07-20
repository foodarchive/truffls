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
	"context"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

// Server an HTTP(s) server.
type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
	tls             struct {
		// TLS certificate, TLS.Key pair.
		Cert []byte
		// TLS private key, TLS.WithCert pair.
		Key []byte
		// TLS certificate file path, TLS.KeyFile pair.
		CertFile string
		// TLS private key file path, TLS.WithCertFile pair.
		KeyFile string
	}
	autoTLS struct {
		// Host allowed host for WithAutoTLS.
		Host string
		// CacheDir certificate caching directory.
		CacheDir string
	}
}

var (
	// ErrMissingCert thrown when starting TLS server without valid certificate.
	ErrMissingCert = errors.New("missing https certificate")
)

// New creates a new Server.
func New(opts ...Option) *Server {
	s := &Server{
		server: &http.Server{
			Addr: ":5000",
		},
	}

	for _, opt := range opts {
		opt.apply(s)
	}
	return s
}

// Start stars HTTP server.
func (s *Server) Start() error {
	return s.start()
}

// StartTLS starts HTTPS server.
//
// The certificate and matching private key must provide
// by setting the TLS option.
//
// Either pair of TLS.CertFile and TLS.KeyFile
// or TLS.Cert and TLS.Key must be provided.
func (s *Server) StartTLS() (err error) {
	var cert, key []byte
	cfg := &tls.Config{Certificates: make([]tls.Certificate, 1)}

	if s.tls.CertFile != "" && s.tls.KeyFile != "" {
		if cert, err = ioutil.ReadFile(s.tls.CertFile); err != nil {
			return
		}
		if key, err = ioutil.ReadFile(s.tls.KeyFile); err != nil {
			return
		}
	} else if s.tls.Cert != nil && s.tls.Key != nil {
		cert, key = s.tls.Cert, s.tls.Key
	} else {
		return ErrMissingCert
	}

	if cfg.Certificates[0], err = tls.X509KeyPair(cert, key); err != nil {
		return
	}

	s.server.TLSConfig = cfg
	return s.start()
}

// StartAutoTLS starts an HTTPS server using certificates
// automatically installed from https://letsencrypt.org.
func (s *Server) StartAutoTLS() error {
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(s.autoTLS.Host),
		Cache:      autocert.DirCache(s.autoTLS.CacheDir),
	}

	cfg := &tls.Config{GetCertificate: m.GetCertificate}
	cfg.NextProtos = append(cfg.NextProtos, acme.ALPNProto)

	s.server.TLSConfig = cfg
	return s.start()
}

func (s *Server) start() error {
	idleConnsClosed := make(chan struct{})

	var g errgroup.Group
	g.Go(func() error {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer func() {
			close(idleConnsClosed)
			cancel()
		}()
		if err := s.server.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		if err := s.listenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	g.Go(func() error {
		select {
		case <-idleConnsClosed:
		}
		return nil
	})
	return g.Wait()
}

func (s *Server) listenAndServe() error {
	if s.server.TLSConfig != nil {
		return s.server.ListenAndServeTLS("", "")
	}
	return s.server.ListenAndServe()
}
