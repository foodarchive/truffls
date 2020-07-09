package config_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/foodarchive/truffls/internal/config"
	pkgConfig "github.com/foodarchive/truffls/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_ = os.Setenv("TRUFFLS_DEBUG", "true")
	defer func() {
		_ = os.Unsetenv("TRUFFLS_DEBUG")
	}()

	pkgConfig.Init("truffls", "")
	c := New()

	fmt.Println(c.Debug)
	assert.True(t, c.Debug)
}
