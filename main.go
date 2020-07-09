package main

import (
	"fmt"

	"github.com/foodarchive/truffls/internal/config"
)

func main() {
	fmt.Println(config.Version)
	fmt.Println(config.BuildDate)
}
