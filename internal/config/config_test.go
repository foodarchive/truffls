package config_test

import (
	"testing"

	. "github.com/foodarchive/truffls/internal/config"
	pkgConfig "github.com/foodarchive/truffls/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	pkgConfig.Load("truffls", "./testdata/config_test.yaml")
	c, err := New()

	assert.NoError(t, err)

	assert.True(t, c.Debug)
}
