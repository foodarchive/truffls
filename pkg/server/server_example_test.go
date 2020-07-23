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
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/foodarchive/truffls/pkg/server"
)

func ExampleServer_Start() {
	srv := server.New(
		mux(),
		server.WithAddr("", "9876"),
	)

	defer func() {
		if err := srv.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(10 * time.Millisecond)

	resp, err := fetch("http://localhost:9876/")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
	// Output:
	// Protocol: HTTP/1.1, Host: localhost:9876, Method: GET, Path: /
}

func ExampleServer_StartTLS() {
	srv := server.New(
		mux(),
		server.WithAddr("", "8765"),
		server.WithCertFile(
			"./testdata/localhost.crt",
			"./testdata/localhost.key",
		),
	)

	defer func() {
		if err := srv.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := srv.StartTLS(); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(20 * time.Millisecond)

	resp, err := fetch("https://localhost:8765/")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
	// Output:
	// Protocol: HTTP/1.1, Host: localhost:8765, Method: GET, Path: /
}

func fetch(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}

	return client.Get(url)
}

func mux() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := "Protocol: %s, Host: %s, Method: %s, Path: %s"
		_, _ = fmt.Fprintf(w, tpl, r.Proto, r.Host, r.Method, r.URL.Path)
	})
	return r
}
