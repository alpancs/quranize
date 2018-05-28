package quranize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNewQuran(t *testing.T, q Quran) {
	expected := 114
	actual := len(q.Suras)
	assert.Equal(t, expected, actual)
}

func TestNewQuranSimpleClean(t *testing.T) {
	testNewQuran(t, NewQuranSimpleClean())
}

func TestNewQuranSimpleEnhanced(t *testing.T) {
	testNewQuran(t, NewQuranSimpleEnhanced())
}

func TestNewIDIndonesian(t *testing.T) {
	testNewQuran(t, NewIDIndonesian())
}

func TestNewIDMuntakhab(t *testing.T) {
	testNewQuran(t, NewIDMuntakhab())
}

func TestGetAyaFound(t *testing.T) {
	text, err := NewQuranSimpleClean().GetAya(1, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, text)
}

func TestGetAyaSuraNotFound(t *testing.T) {
	_, err := NewQuranSimpleClean().GetAya(0, 0)
	assert.Error(t, err)
}

func TestGetAyaAyaNotFound(t *testing.T) {
	_, err := NewQuranSimpleClean().GetAya(1, 0)
	assert.Error(t, err)
}

func TestGetSuraNameFound(t *testing.T) {
	text, err := NewQuranSimpleClean().GetSuraName(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, text)
}

func TestGetSuraNameNotFound(t *testing.T) {
	_, err := NewQuranSimpleClean().GetSuraName(0)
	assert.Error(t, err)
}
