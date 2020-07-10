package config

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Flag = pflag.Flag

func Load(namespace, filename string) {
	if namespace == "" {
		log.Fatal("config namespace must be provided")
	}

	if filename != "" {
		viper.SetConfigFile(filename)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/" + namespace)
		viper.AddConfigPath("$HOME/." + namespace)
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix(namespace)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func Unmarshal(v interface{}) error {
	return viper.Unmarshal(v, func(c *mapstructure.DecoderConfig) {
		c.TagName = "config"
		c.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		)
	})
}

func BindFlags(flags ...*Flag) {
	for _, flag := range flags {
		_ = viper.BindPFlag(flag.Name, flag)
	}
}
