package config

import (
	"log"

	pkgConfig "github.com/foodarchive/truffls/pkg/config"
)

var (
	// Version is dynamically set by the toolchain or overridden by the Makefile.
	Version = "DEV"

	// BuildDate is dynamically set at build time in the Makefile.
	BuildDate = "" // YYYY-MM-DD
)

type Config struct {
	Debug bool
}

func New() Config {
	var c Config
	if err := pkgConfig.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	return c
}
