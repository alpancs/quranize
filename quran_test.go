package quran

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	quranTest        = NewQuranSimpleClean()
	quranTestIndexed = NewQuranSimpleClean()
)

func TestMain(m *testing.M) {
	quranTestIndexed.BuildIndex()
	os.Exit(m.Run())
}

func TestGetAyaFound(t *testing.T) {
	text, err := quranTest.GetAya(1, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, text)
}

func TestGetAyaSuraNotFound(t *testing.T) {
	_, err := quranTest.GetAya(0, 0)
	assert.Error(t, err)
}

func TestGetAyaAyaNotFound(t *testing.T) {
	_, err := quranTest.GetAya(1, 0)
	assert.Error(t, err)
}

func TestGetSuraNameFound(t *testing.T) {
	text, err := quranTest.GetSuraName(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, text)
}

func TestGetSuraNameNotFound(t *testing.T) {
	_, err := quranTest.GetSuraName(0)
	assert.Error(t, err)
}

func TestLocateEmptyString(t *testing.T) {
	input := ""
	expected := zeroLocs
	actual := quranTestIndexed.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateNonAlquran(t *testing.T) {
	input := "alfan"
	expected := zeroLocs
	actual := quranTestIndexed.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateAlquran(t *testing.T) {
	input := "بسم الله الرحمن الرحيم"
	expected := []Location{Location{1, 1, 0}, Location{27, 30, 4}}
	actual := quranTestIndexed.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateAlquranBeforeBuildIndex(t *testing.T) {
	input := "بسم الله الرحمن الرحيم"
	expected := zeroLocs
	actual := quranTest.Locate(input)
	assert.Equal(t, expected, actual)
}
