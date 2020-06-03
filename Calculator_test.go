package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculate(t *testing.T)  {
	t.Parallel()

	result := calculate(100000, 10, 10)

	require.Equal(t, 1321.5073688176203, result)
}

func TestCountPayments(t *testing.T)  {
	t.Parallel()

	result := countPayments(5)

	require.Equal(t, 60.0, result)
}

func TestAssumePayments(t *testing.T)  {
	t.Parallel()

	result := assumePayments(150)

	require.Equal(t, 12.5, result)
}