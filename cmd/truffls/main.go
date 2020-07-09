/*
Copyright 2020 @truffls contributors.
Licensed under the GNU GENERAL PUBLIC LICENSE, Version 3.0 (the "License");
you may not use this file except in compliance with the License.
*/

package main

import (
	"fmt"

	"github.com/foodarchive/truffls/internal/config"
)

func main() {
	fmt.Println(config.Version)
	fmt.Println(config.BuildDate)
}
