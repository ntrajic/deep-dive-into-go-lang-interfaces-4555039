package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpeN(t *testing.T) {
	_, err := Open("localhost")
	require.NoError(t, err)
}
