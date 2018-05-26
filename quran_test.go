package quran

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAyaFound(t *testing.T) {
	text, err := quran.GetAya(1, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, text)
}

func TestGetAyaSuraNotFound(t *testing.T) {
	_, err := quran.GetAya(0, 0)
	assert.Error(t, err)
}

func TestGetAyaAyaNotFound(t *testing.T) {
	_, err := quran.GetAya(1, 0)
	assert.Error(t, err)
}

func TestGetSuraNameFound(t *testing.T) {
	text, err := quran.GetSuraName(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, text)
}

func TestGetSuraNameNotFound(t *testing.T) {
	_, err := quran.GetSuraName(0)
	assert.Error(t, err)
}

func TestLocateEmptyString(t *testing.T) {
	input := ""
	expected := []Location{}
	actual := Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateNonAlquran(t *testing.T) {
	input := "alfan"
	expected := []Location{}
	actual := Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateAlquran(t *testing.T) {
	input := "بسم الله الرحمن الرحيم"
	expected := []Location{Location{1, 1, 0}, Location{27, 30, 4}}
	actual := Locate(input)
	assert.Equal(t, expected, actual)
}
