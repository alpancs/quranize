package quran

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/alpancs/quran/corpus"
)

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
	root *Node
}

type Node struct {
	locations []Location
	children  []Child
}

type Location struct {
	Sura      int // sura number
	Aya       int // aya number
	WordIndex int // assuming aya is splitted word by word
}

type Child struct {
	key   rune
	value *Node
}

var (
	zeroLocs = make([]Location, 0, 0)
)

// Returns Quran from given file path.
func LoadQuran(filePath string) (q Quran, err error) {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	return ParseQuran(raw)
}

// Returns Quran from given raw.
func ParseQuran(raw []byte) (q Quran, err error) {
	err = xml.Unmarshal(raw, &q)
	return
}

// Return new instance from quran-simple-clean.xml.
// See corpus/QuranSimpleClean.
func NewQuranSimpleClean() Quran {
	q, _ := ParseQuran([]byte(corpus.QuranSimpleCleanXML))
	return q
}

// Returns sura name from sura number in Quran q (number starting from 1).
func (q Quran) GetSuraName(sura int) (string, error) {
	if !(1 <= sura && sura <= len(q.Suras)) {
		return "", errors.New(fmt.Sprintf("invalid sura number %d", sura))
	}
	return q.Suras[sura-1].Name, nil
}

// Returns aya text from sura number and aya number in Quran q (number starting from 1).
func (q Quran) GetAya(sura, aya int) (string, error) {
	if !(1 <= sura && sura <= len(q.Suras)) {
		return "", errors.New(fmt.Sprintf("invalid sura number %d", sura))
	}
	ayas := q.Suras[sura-1].Ayas
	if !(1 <= aya && aya <= len(ayas)) {
		return "", errors.New(fmt.Sprintf("invalid sura number %d and aya number %d", sura, aya))
	}
	return ayas[aya-1].Text, nil
}

// Returns locations of s (quran kalima) in Quran q, matching the whole word.
func (q Quran) Locate(s string) []Location {
	if q.root == nil {
		return zeroLocs
	}

	harfs := []rune(s)
	node := q.root
	for _, harf := range harfs {
		node = node.getChild(harf)
		if node == nil {
			return zeroLocs
		}
	}
	return node.locations
}

// Returns locations of s (quran kalima) in Quran q, matching the whole word.
func (q Quran) exists(s string) bool {
	harfs := []rune(s)
	node := q.root
	for _, harf := range harfs {
		node = node.getChild(harf)
		if node == nil {
			return false
		}
	}
	return true
}

// Build index for Quran q.
// Without index, `Locate` function won't work.
func (q *Quran) BuildIndex() {
	q.root = &Node{locations: zeroLocs}
	for _, sura := range q.Suras {
		for _, aya := range sura.Ayas {
			q.indexAya([]rune(aya.Text), sura.Index, aya.Index)
		}
	}
}

func (q *Quran) indexAya(harfs []rune, sura, aya int) {
	sliceIndex := 0
	for i := range harfs {
		if i == 0 || harfs[i-1] == ' ' {
			q.buildTree(harfs[i:], Location{sura, aya, sliceIndex})
			sliceIndex++
		}
	}
}

func (q *Quran) buildTree(harfs []rune, location Location) {
	node := q.root
	for i, harf := range harfs {
		child := node.getChild(harf)
		if child == nil {
			child = &Node{}
			node.children = append(node.children, Child{harf, child})
		}
		node = child
		if i == len(harfs)-1 || harfs[i+1] == ' ' {
			node.locations = append(node.locations, location)
		}
	}
}

func (n *Node) getChild(key rune) *Node {
	for _, child := range n.children {
		if child.key == key {
			return child.value
		}
	}
	return nil
}
