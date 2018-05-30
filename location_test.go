package quranize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationGetSura(t *testing.T) {
	expected := 2
	actual := NewLocation(2, 286, 0).GetSura()
	assert.Equal(t, expected, actual)
}

func TestLocationGetAya(t *testing.T) {
	expected := 286
	actual := NewLocation(2, 286, 0).GetAya()
	assert.Equal(t, expected, actual)
}

func TestLocationGetWordIndex(t *testing.T) {
	expected := 0
	actual := NewLocation(2, 286, 0).GetWordIndex()
	assert.Equal(t, expected, actual)
}
