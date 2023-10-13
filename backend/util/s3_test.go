package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestS3GetObjectFailed(t *testing.T) {
	c := require.New(t)
	_, err := GetObject("test", "test")
	c.Error(err)
}
