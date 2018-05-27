package quran

import (
	"strings"
)

// Quranize encodes arabic into alphabet.
type Quranize struct {
	t Transliteration
	q Quran
}

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
	q.BuildIndex()
	return Quranize{t, q}
}

// Encode returns arabic encodings of given string using Transliteration t.
func (q Quranize) Encode(s string) []string {
	var memo = make(map[string][]string)
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	results := []string{}
	for _, result := range q.quranize(s, memo) {
		if len(q.q.Locate(result)) > 0 {
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
				if q.q.exists(combination) {
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
