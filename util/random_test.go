package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOwner(t *testing.T) {
	s := RandomOwner()
	require.IsType(t, "string" , s)
	require.NotZero(t, s)
}

func TestRandomMoney(t *testing.T) {
	i := RandomMoney()
	require.IsType(t, int64(0), i)
}

func TestRandomCurrency(t *testing.T) {
	s := RandomCurrency()
	require.IsType(t, "string" , s)
	require.NotZero(t, s)
	require.Len(t, s, 3)
}