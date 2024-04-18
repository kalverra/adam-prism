package main

import (
	"errors"
	"os"
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

func TestReadEnvFile(t *testing.T) {
	t.Parallel()
	if _, err := os.Stat("./.env"); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile("./.env", []byte("LATITUDE=1.0\nLONGITUDE=1.0\nNOTION_KEY=key"), 0644)
		require.NoError(t, err)
	}

	config, err := ReadConfig()
	require.NoError(t, err)
	require.NotNil(t, config)
}
