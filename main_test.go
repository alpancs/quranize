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
	quranTest = NewQuranSimpleClean()
	quranizeTest = NewDefaultQuranize()
	os.Exit(m.Run())
}
