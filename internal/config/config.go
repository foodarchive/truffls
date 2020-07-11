package config

import (
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

func New() (Config, error) {
	var c Config
	err := pkgConfig.Unmarshal(&c)
	return c, err
}
