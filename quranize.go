// Package quranize provides Go representation of Alquran.
// Original source of Alquran is taken from http://tanzil.net in XML format.
//
// This package can transform alphabet into arabic using fast and efficient algorithm:
// suffix-tree for indexing and dynamic programming for parsing.
package quranize

import (
	"strings"
	"sync"
)

// Quranize encodes arabic into alphabet.
type Quranize struct {
	t    Transliteration
	q    Quran
	root *node
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
	q        Quranize
	once     sync.Once
	zeroLocs = make([]Location, 0, 0)
)

// NewDefaultQuranize returns new Quranize using default mapping and quran "quran-simple-clean.xml".
//
// Mapping: https://github.com/alpancs/quranize/blob/master/corpus/arabic_to_alphabet_clean.go#L3
//
// Quran: https://github.com/alpancs/quranize/blob/master/corpus/quran_simple_clean.go#L4
func NewDefaultQuranize() Quranize {
	once.Do(func() {
		q = NewQuranize(NewDefaultTransliteration(), NewQuranSimpleClean())
	})
	return q
}

// NewQuranize return new Quranize using Transliteration t and Quran q.
func NewQuranize(t Transliteration, q Quran) Quranize {
	quranize := Quranize{t: t, q: q}
	quranize.buildIndex()
	return quranize
}

// Encode returns arabic encodings of given string using Transliteration t.
func (q Quranize) Encode(s string) []string {
	s = strings.ToLower(strings.Replace(s, " ", "", -1))
	var memo = make(map[string][]string)
	dirtyResults := append(q.quranize(s, memo), q.quranize(trimLastNonVowel(s), memo)...)
	results := []string{}
	for _, result := range dirtyResults {
		if len(q.Locate(result)) > 0 {
			results = appendUniq(results, result)
		}
	}
	return results
}

func trimLastNonVowel(s string) string {
	if len(s) == 0 {
		return s
	}
	switch s[len(s)-1] {
	case 'a', 'e', 'i', 'o', 'u':
		return s
	}
	return s[:len(s)-1]
}

// Locate returns locations of s (quran kalima), matching the whole word.
func (q Quranize) Locate(s string) []Location {
	n := q.root
	if n == nil {
		return zeroLocs
	}

	for _, harf := range []rune(s) {
		n = n.getChild(harf)
		if n == nil {
			return zeroLocs
		}
	}
	return n.locations
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
			combinations = append(combinations, head+tail+"ى")
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

// exists returns existence of s.
func (q Quranize) exists(s string) bool {
	n := q.root
	for _, harf := range []rune(s) {
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
	for s, sura := range q.q.Suras {
		for a, aya := range sura.Ayas {
			q.indexAya([]rune(aya.Text), s+1, a+1)
		}
	}
}

func (q *Quranize) indexAya(harfs []rune, sura, aya int) {
	wordIndex := 0
	for i := range harfs {
		if i == 0 || harfs[i-1] == ' ' {
			q.buildTree(harfs[i:], NewLocation(sura, aya, wordIndex))
			wordIndex++
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
