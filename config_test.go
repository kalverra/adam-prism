package main

import (
	"testing"

	"github.com/test-go/testify/require"
)

func TestReadConfig(t *testing.T) {
	t.Parallel()

	t.Setenv("LATITUDE", "1.0")
	t.Setenv("LONGITUDE", "1.0")
	t.Setenv("NOTION_KEY", "key")
	config, err := ReadConfig()
	require.NoError(t, err)
	require.NotNil(t, config)
	require.Equal(t, 1.0, config.Latitude)
	require.Equal(t, 1.0, config.Longitude)
	require.Equal(t, "key", config.NotionKey)
}

func TestReadConfigError(t *testing.T) {
	t.Parallel()

	t.Setenv("LATITUDE", "invalid")
	t.Setenv("LONGITUDE", "invalid")
	t.Setenv("NOTION_KEY", "key")
	config, err := ReadConfig()
	require.Error(t, err)
	require.Nil(t, config)
}
