package es

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestMinStreamID(t *testing.T) {

	const (
		id1 = "1766371548463-0"
		id2 = "1766371548463-2"
		id3 = "1766371548463-3"
		id4 = "1766371548466-0"
		id5 = "1766371548466-1"
		id6 = "1766371548466-2"
		id0 = "0-10"
	)

	s := []string{id1, id2, id3, id4, id5, id6, id0}
	slices.Sort(s)

	minStreamID, minOk := MinStreamID(s...)
	maxStreamID, maxOk := MaxStreamID(s...)
	assert.Equal(t, minStreamID, id0)
	assert.True(t, minOk)

	assert.Equal(t, maxStreamID, id6)
	assert.True(t, maxOk)

	t.Logf("min stream id: %s", minStreamID)
	t.Logf("max stream id: %s", maxStreamID)
}
