package quran

import (
	"strings"
)

// Quranize encodes arabic into alphabet.
type Quranize struct {
	t    Transliteration
	q    Quran
	root *node
}

// Location represents a location in Quran.
type Location struct {
	Sura      int // sura number
	Aya       int // aya number
	WordIndex int // assuming aya is splitted word by word using separator " "
}

type node struct {
	locations []Location
	children  []child
}

type child struct {
	key   rune
	value *node
}

var (
	zeroLocs = make([]Location, 0, 0)
)

// NewDefaultQuranize returns new Quranize using default mapping and quran "quran-simple-clean.yml".
//
// Mapping: https://github.com/alpancs/quran/blob/master/corpus/arabic_to_alphabet.go#L3
//
// Quran: https://github.com/alpancs/quran/blob/master/corpus/quran_simple_clean.go#L4
func NewDefaultQuranize() Quranize {
	return NewQuranize(NewDefaultTransliteration(), NewQuranSimpleClean())
}

// NewQuranize return new Quranize using Transliteration t and Quran q.
func NewQuranize(t Transliteration, q Quran) Quranize {
	quranize := Quranize{t, q, nil}
	quranize.buildIndex()
	return quranize
}

// Encode returns arabic encodings of given string using Transliteration t.
func (q Quranize) Encode(s string) []string {
	var memo = make(map[string][]string)
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	results := []string{}
	for _, result := range q.quranize(s, memo) {
		if len(q.Locate(result)) > 0 {
			results = appendUniq(results, result)
		}
	}
	return results
}

func (q Quranize) quranize(s string, memo map[string][]string) []string {
	if s == "" {
		return base
	}

	if cache, ok := memo[s]; ok {
		return cache
	}

	kalimas := []string{}
	l := len(s)
	for width := 1; width <= q.t.alphabetMaxLen && width <= l; width++ {
		if tails, ok := q.t.hijaiyas[s[l-width:]]; ok {
			heads := q.quranize(s[:l-width], memo)
			for _, combination := range combine(heads, tails) {
				if q.exists(combination) {
					kalimas = appendUniq(kalimas, combination)
				}
			}
		}
	}

	memo[s] = kalimas
	return kalimas
}

func combine(heads, tails []string) []string {
	combinations := []string{}
	for _, head := range heads {
		for _, tail := range tails {
			combinations = append(combinations, head+tail)
			combinations = append(combinations, head+" "+tail)
			combinations = append(combinations, head+"ا"+tail)
			combinations = append(combinations, head+"ال"+tail)
			combinations = append(combinations, head+" ال"+tail)
			if tail == "و" {
				combinations = append(combinations, head+tail+"ا")
			}
		}
	}
	return combinations
}

func appendUniq(results []string, newResult string) []string {
	for _, result := range results {
		if result == newResult {
			return results
		}
	}
	return append(results, newResult)
}

// Locate returns locations of s (quran kalima), matching the whole word.
func (q Quranize) Locate(s string) []Location {
	if q.root == nil {
		return zeroLocs
	}

	harfs := []rune(s)
	n := q.root
	for _, harf := range harfs {
		n = n.getChild(harf)
		if n == nil {
			return zeroLocs
		}
	}
	return n.locations
}

// exists returns existence of s.
func (q Quranize) exists(s string) bool {
	harfs := []rune(s)
	n := q.root
	for _, harf := range harfs {
		n = n.getChild(harf)
		if n == nil {
			return false
		}
	}
	return true
}

// buildIndex build index for Quranize q.
//
// Without index,
//   q.Locate
// won't work.
func (q *Quranize) buildIndex() {
	q.root = &node{locations: zeroLocs}
	for _, sura := range q.q.Suras {
		for _, aya := range sura.Ayas {
			q.indexAya([]rune(aya.Text), sura.Index, aya.Index)
		}
	}
}

func (q *Quranize) indexAya(harfs []rune, sura, aya int) {
	sliceIndex := 0
	for i := range harfs {
		if i == 0 || harfs[i-1] == ' ' {
			q.buildTree(harfs[i:], Location{sura, aya, sliceIndex})
			sliceIndex++
		}
	}
}

func (q *Quranize) buildTree(harfs []rune, location Location) {
	n := q.root
	for i, harf := range harfs {
		c := n.getChild(harf)
		if c == nil {
			c = &node{}
			n.children = append(n.children, child{harf, c})
		}
		n = c
		if i == len(harfs)-1 || harfs[i+1] == ' ' {
			n.locations = append(n.locations, location)
		}
	}
}

func (n *node) getChild(key rune) *node {
	for _, c := range n.children {
		if c.key == key {
			return c.value
		}
	}
	return nil
}
