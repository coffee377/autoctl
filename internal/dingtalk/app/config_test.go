package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	configurations := ReadConfiguration()
	assert.Equal(t, 2, len(configurations.App))
}
