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

package signal

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// Receiver represent a component (server, subsystem, etc) that can be stopped
// or queried about the status with a signal
type Receiver interface {
	Stop() error
}

// Handler handles signals, can be interrupted.
// On SIGINT or SIGTERM it will exit, on SIGQUIT it
// will dump goroutine stacks to the Logger.
type Handler interface {
	// Stop the handler.
	Stop()
	// Loop handles signals.
	Loop() error
}

type handler struct {
	receivers []Receiver
	quit      chan struct{}
}

// NewHandler create a new signal handler.
func NewHandler(receivers ...Receiver) Handler {
	return &handler{
		receivers: receivers,
		quit:      make(chan struct{}),
	}
}

// Stop the handler.
func (h *handler) Stop() {
	close(h.quit)
}

// Loop keep looping until it received SIGINT, SIGTERM or SIGQUIT.
// When SIGINT or SIGTERM called it will call stop method on the receivers.
func (h *handler) Loop() error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer signal.Stop(signals)

	for {
		select {
		case <-h.quit:
			log.Println("signal handler stopped")
			return nil
		case sig := <-signals:
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				log.Println("received SIGINT/SIGTERM, exiting")
				for _, receiver := range h.receivers {
					_ = receiver.Stop() // TODO: handle error
				}
				return nil
			case syscall.SIGQUIT:
				buf := make([]byte, 1<<20)
				stacklen := runtime.Stack(buf, true)
				log.Printf("received SIGQUIT goroutine dump....%s\n", buf[:stacklen])
				return nil
			}
		}
	}
}

// HandlerLoop blocks until it receives a SIGINT, SIGTERM or SIGQUIT.
// For SIGINT and SIGTERM, it exits; for SIGQUIT is print a goroutine stack
// dump.
func HandlerLoop(receivers ...Receiver) error {
	return NewHandler(receivers...).Loop()
}
