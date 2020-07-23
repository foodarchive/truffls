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

package server_test

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/foodarchive/truffls/pkg/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStart(t *testing.T) {
	srv := server.New(
		http.NewServeMux(),
		server.WithAddr("", "0"),
	)
	go func() {
		assert.NoError(t, srv.Start())
	}()

	time.Sleep(10 * time.Millisecond)
	srv.Stop()
}

func TestStartTLS(t *testing.T) {
	cert, err := ioutil.ReadFile("./testdata/localhost.crt")
	require.NoError(t, err)
	key, err := ioutil.ReadFile("./testdata/localhost.key")
	require.NoError(t, err)

	testCases := []struct {
		certFile  string
		keyFile   string
		cert      []byte
		key       []byte
		assertion assert.ErrorAssertionFunc
		name      string
	}{
		{
			certFile:  "./testdata/localhost.crt",
			keyFile:   "./testdata/localhost.key",
			assertion: assert.NoError,
			name:      "ValidCertKeyFile",
		},
		{

			cert:      cert,
			key:       key,
			assertion: assert.NoError,
			name:      "ValidCertKey",
		},
		{
			assertion: assert.Error,
			name:      "InvalidCert",
		},
		{

			certFile:  "./testdata/invalid.crt",
			keyFile:   "./testdata/localhost.key",
			assertion: assert.Error,
			name:      "InvalidCertFile",
		},
		{
			certFile:  "./testdata/localhost.crt",
			keyFile:   "./testdata/invalid.key",
			assertion: assert.Error,
			name:      "InvalidKeyFile",
		},
		{
			cert:      []byte{1},
			key:       []byte{2},
			assertion: assert.Error,
			name:      "InvalidKeyPair",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			srv := server.New(
				http.NewServeMux(),
				server.WithAddr("", "0"),
				server.WithCertFile(tc.certFile, tc.keyFile),
				server.WithCert(tc.cert, tc.key),
			)
			go func() {
				err := srv.StartTLS()
				tc.assertion(t, err)
			}()

			time.Sleep(10 * time.Millisecond)
			srv.Stop()
		})
	}
}

func TestStartAutoTLS(t *testing.T) {
	srv := server.New(
		http.NewServeMux(),
		server.WithAddr("", "0"),
		server.WithAutoTLS("", "./testdata"),
		server.WithConfig(server.Config{
			ReadTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       5 * time.Second,
			ShutdownTimeout:   5 * time.Second,
			ReadHeaderTimeout: 500 * time.Millisecond,
			MaxHeaderBytes:    200,
		}),
	)
	go func() {
		assert.NoError(t, srv.StartAutoTLS())
	}()

	time.Sleep(10 * time.Millisecond)
	srv.Stop()
}
