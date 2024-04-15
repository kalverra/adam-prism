package main

import (
	"testing"

	"github.com/test-go/testify/require"
)

func TestReadConfig(t *testing.T) {
	t.Setenv("LATITUDE", "1.0")
	t.Setenv("LONGITUDE", "1.0")
	t.Setenv("NOTION_KEY", "key")
	config, err := ReadConfig()
	require.NoError(t, err)
	require.NotNil(t, config)
}

func TestReadConfigError(t *testing.T) {
	t.Setenv("LATITUDE", "invalid")
	t.Setenv("LONGITUDE", "invalid")
	t.Setenv("NOTION_KEY", "key")
	config, err := ReadConfig()
	require.Error(t, err)
	require.Equal(t, config.Latitude, 0.0)
}
