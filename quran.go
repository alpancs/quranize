package quran

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
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
	quran Quran
	root  *Node

	zeroLocs = make([]Location, 0, 0)
)

func buildIndex() {
	root = &Node{locations: zeroLocs}
	for _, sura := range quran.Suras {
		for _, aya := range sura.Ayas {
			indexAya([]rune(aya.Text), sura.Index, aya.Index, root)
		}
	}
}

func indexAya(harfs []rune, sura, aya int, node *Node) {
	sliceIndex := 0
	for i := range harfs {
		if i == 0 || harfs[i-1] == ' ' {
			buildTree(harfs[i:], Location{sura, aya, sliceIndex}, node)
			sliceIndex++
		}
	}
}

func buildTree(harfs []rune, location Location, node *Node) {
	for i, harf := range harfs {
		child := getChild(node.children, harf)
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

func getChild(children []Child, key rune) *Node {
	for _, child := range children {
		if child.key == key {
			return child.value
		}
	}
	return nil
}

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
