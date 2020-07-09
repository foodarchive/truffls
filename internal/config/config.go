package config

import (
	pkgConfig "github.com/foodarchive/truffls/pkg/config"
)

// Version is dynamically set by the toolchain or overridden by the Makefile.
var Version = "DEV"

// BuildDate is dynamically set at build time in the Makefile.
var BuildDate = "" // YYYY-MM-DD

type Config struct {
	Debug bool
}

func New() Config {
	var c Config
	_ = pkgConfig.Unmarshal(&c)
	return c
}
