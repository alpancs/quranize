// Package quran provides Go representation of Alquran.
// Original source of Alquran is taken from http://tanzil.net in XML format.
//
// This package can transform alphabet into arabic using fast and efficient algorithm:
// suffix-tree for indexing and dynamic programming for parsing.
package quran

import (
	"encoding/xml"
	"fmt"

	"github.com/alpancs/quran/corpus"
)

// Quran stores information of every sura and aya.
// It has suffix-tree index.
type Quran struct {
	Suras []struct {
		Index int    `xml:"index,attr"`
		Name  string `xml:"name,attr"`
		Ayas  []struct {
			Index     int    `xml:"index,attr"`
			Text      string `xml:"text,attr"`
			Bismillah string `xml:"bismillah,attr"`
		} `xml:"aya"`
	} `xml:"sura"`
}

// NewQuranSimpleClean returns new instance from quran-simple-clean.xml.
//
// See https://github.com/alpancs/quran/blob/master/corpus/quran_simple_clean.go#L4.
func NewQuranSimpleClean() Quran {
	q, _ := ParseQuran(corpus.QuranSimpleCleanXML)
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
