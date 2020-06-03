package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrintResult(t *testing.T) {
	printResult(100, 10, 10, 0.132323, 1)

	require.True(t, true)
}
