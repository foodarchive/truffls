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
		server.WithAddr(":0"),
		server.WithHandler(http.NewServeMux()),
	)
	go func() {
		assert.NoError(t, srv.Start())
	}()

	time.Sleep(10 * time.Millisecond)
}

func TestStartTLS(t *testing.T) {
	cert, err := ioutil.ReadFile("./testdata/localhost.crt")
	require.NoError(t, err)
	key, err := ioutil.ReadFile("./testdata/localhost.key")
	require.NoError(t, err)

	testCases := []struct {
		tlsConfig server.TLS
		assertion assert.ErrorAssertionFunc
		name      string
	}{
		{
			tlsConfig: server.TLS{
				CertFile: "./testdata/localhost.crt",
				KeyFile:  "./testdata/localhost.key",
			},
			assertion: assert.NoError,
			name:      "ValidCertKeyFile",
		},
		{
			tlsConfig: server.TLS{
				Cert: cert,
				Key:  key,
			},
			assertion: assert.NoError,
			name:      "ValidCertKey",
		},
		{
			tlsConfig: server.TLS{},
			assertion: assert.Error,
			name:      "InvalidCert",
		},
		{
			tlsConfig: server.TLS{
				CertFile: "./testdata/invalid.crt",
				KeyFile:  "./testdata/localhost.key",
			},
			assertion: assert.Error,
			name:      "InvalidCertFile",
		},
		{
			tlsConfig: server.TLS{
				CertFile: "./testdata/localhost.crt",
				KeyFile:  "./testdata/invalid.key",
			},
			assertion: assert.Error,
			name:      "InvalidKeyFile",
		},
		{
			tlsConfig: server.TLS{
				Cert: []byte{1},
				Key:  []byte{2},
			},
			assertion: assert.Error,
			name:      "InvalidKeyPair",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			srv := server.New(
				server.WithHandler(http.NewServeMux()),
				server.WithAddr(":0"),
				server.WithCertFile(tc.tlsConfig.CertFile, tc.tlsConfig.KeyFile),
				server.WithCert(tc.tlsConfig.Cert, tc.tlsConfig.Key),
			)
			go func() {
				err := srv.StartTLS()
				tc.assertion(t, err)
			}()

			time.Sleep(10 * time.Millisecond)
		})
	}
}

func TestStartAutoTLS(t *testing.T) {
	srv := server.New(
		server.WithHandler(http.NewServeMux()),
		server.WithAddr(":0"),
		server.WithAutoTLS("", "./testdata"),
	)
	go func() {
		assert.NoError(t, srv.StartAutoTLS())
	}()

	time.Sleep(10 * time.Millisecond)
}
