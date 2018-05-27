package quran

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransliterationEmpty(t *testing.T) {
	input := ""
	expected := Transliteration{make(map[string][]string), 0}
	actual := NewTransliteration(input)
	assert.Equal(t, expected, actual)
}
