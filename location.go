package quranize

// Location represents a location in Quran.
type Location struct {
	sura      uint8
	aya       uint16
	wordIndex uint8
}

// NewLocation returns new Location given sura number, aya number, and "word index"
// (assuming aya is splitted word by word using separator " ").
func NewLocation(sura, aya, wordIndex int) Location {
	return Location{uint8(sura), uint16(aya), uint8(wordIndex)}
}

// GetSura returns sura number of this location.
func (l Location) GetSura() int {
	return int(l.sura)
}

// GetAya returns aya number of this location.
func (l Location) GetAya() int {
	return int(l.aya)
}

// GetWordIndex returns word index of this location.
func (l Location) GetWordIndex() int {
	return int(l.wordIndex)
}
