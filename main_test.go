package quran

import (
	"os"
	"testing"
)

var (
	quranTest, quranTestIndexed Quran
	quranizeTest                Quranize
)

func TestMain(m *testing.M) {
	quranizeTest = NewDefaultQuranize()

	quranTest = NewQuranSimpleClean()
	quranTestIndexed = quranizeTest.q

	os.Exit(m.Run())
}
