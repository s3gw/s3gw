package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRootCommand(t *testing.T) {
	got := GetRootCommand()
	assert.Equal(t, "s3gw", got.Use)
}
