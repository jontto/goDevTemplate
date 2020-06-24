package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	err := echo([]string{"bin-name", "hello", "world!"})
	require.NoError(t, err)
}

func TestEchoErrorNoArgas(t *testing.T) {
	err := echo([]string{})
	require.Error(t, err)
}
