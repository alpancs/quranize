package quranize

import (
	"encoding/xml"
	"fmt"

	"github.com/alpancs/quranize/corpus"
)

// Quran stores information of every sura and aya.
// It has suffix-tree index.
type Quran struct {
	Suras []struct {
		Name string `xml:"name,attr"`
		Ayas []struct {
			Text string `xml:"text,attr"`
		} `xml:"aya"`
	} `xml:"sura"`
}

// NewQuranSimpleClean returns new Quran instance using corpus:
//  corpus.QuranSimpleCleanXML
// See https://github.com/alpancs/quranize/blob/master/corpus/quran_simple_clean.go#L4.
func NewQuranSimpleClean() Quran {
	q, _ := ParseQuran(corpus.QuranSimpleCleanXML)
	return q
}

// NewQuranSimpleEnhanced returns new Quran instance using corpus:
//  corpus.QuranSimpleEnhancedXML
// See https://github.com/alpancs/quranize/blob/master/corpus/quran_simple_enhanced.go#L4.
func NewQuranSimpleEnhanced() Quran {
	q, _ := ParseQuran(corpus.QuranSimpleEnhancedXML)
	return q
}

// NewIDIndonesian returns new Quran instance using corpus:
//  corpus.IDIndonesianXML
// See https://github.com/alpancs/quranize/blob/master/corpus/id_indonesian.go#L4.
func NewIDIndonesian() Quran {
	q, _ := ParseQuran(corpus.IDIndonesianXML)
	return q
}

// NewIDMuntakhab returns new Quran instance using corpus:
//  corpus.IDMuntakhabXML
// See https://github.com/alpancs/quranize/blob/master/corpus/id_muntakhab.go#L4.
func NewIDMuntakhab() Quran {
	q, _ := ParseQuran(corpus.IDMuntakhabXML)
	return q
}

// ParseQuran returns Quran from given raw.
func ParseQuran(raw string) (q Quran, err error) {
	err = xml.Unmarshal([]byte(raw), &q)
	return
}

// GetSuraName returns sura name from sura number in Quran q (number starting from 1).
func (q Quran) GetSuraName(sura int) (string, error) {
	if !(1 <= sura && sura <= len(q.Suras)) {
		return "", fmt.Errorf("invalid sura number %d", sura)
	}
	return q.Suras[sura-1].Name, nil
}

// GetAya returns aya text from sura number and aya number in Quran q (number starting from 1).
func (q Quran) GetAya(sura, aya int) (string, error) {
	if !(1 <= sura && sura <= len(q.Suras)) {
		return "", fmt.Errorf("invalid sura number %d", sura)
	}
	ayas := q.Suras[sura-1].Ayas
	if !(1 <= aya && aya <= len(ayas)) {
		return "", fmt.Errorf("invalid sura number %d and aya number %d", sura, aya)
	}
	return ayas[aya-1].Text, nil
}
