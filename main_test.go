package quran

import (
	"os"
	"testing"
)

var (
	quranTest, quranTestIndexed Quran
	transliterationTest         Transliteration
)

func TestMain(m *testing.M) {
	transliterationTest = NewDefaultTransliteration()

	quranTest = NewQuranSimpleClean()
	quranTestIndexed = transliterationTest.q

	os.Exit(m.Run())
}
