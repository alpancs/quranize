package quranize

import (
	"fmt"
)

// Location represents a location in Quran.
type Location struct {
	sum int32
}

const (
	ayaMultiplier  = int32(1000)
	suraMultiplier = int32(1000 * ayaMultiplier)
)

// NewLocation returns new Location given sura number, aya number, and "word index"
// (assuming aya is splitted word by word using separator " ").
func NewLocation(sura, aya, wordIndex int) Location {
	return Location{int32(sura)*suraMultiplier + int32(aya)*ayaMultiplier + int32(wordIndex)}
}

// GetSura returns sura number of this location.
func (l Location) GetSura() int {
	return int(l.sum / suraMultiplier)
}

// GetAya returns aya number of this location.
func (l Location) GetAya() int {
	return int((l.sum % suraMultiplier) / ayaMultiplier)
}

// GetWordIndex returns word index of this location.
func (l Location) GetWordIndex() int {
	return int(l.sum % ayaMultiplier)
}

// String returns string representation of this location.
// Format: "<sura>.<aya>.<word index>".
func (l Location) String() string {
	return fmt.Sprintf("%d.%d.%d", l.GetSura(), l.GetAya(), l.GetWordIndex())
}
