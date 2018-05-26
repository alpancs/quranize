package quran

import (
	"strings"

	"github.com/alpancs/quran/corpus"
)

// Transliteration is used to encode arabic into alphabet.
type Transliteration struct {
	hijaiyas       map[string][]string
	alphabetMaxLen int
	q              Quran
}

var (
	base = []string{""}
)

// NewDefaultTransliteration returns new default Transliteration.
//
// Default map: https://github.com/alpancs/quran/blob/master/corpus/arabic_to_alphabet.go#L3.
//
// Default Quran: https://github.com/alpancs/quran/blob/master/corpus/quran_simple_clean.go#L4.
func NewDefaultTransliteration() Transliteration {
	return NewTransliteration(corpus.ArabicToAlphabet, NewQuranSimpleClean())
}

// NewTransliteration returns new Transliteration.
func NewTransliteration(raw string, q Quran) Transliteration {
	hijaiyas := make(map[string][]string)
	alphabetMaxLen := 0

	trimmed := strings.TrimSpace(raw)
	for _, line := range strings.Split(trimmed, "\n") {
		components := strings.Split(line, " ")
		arabic := components[0]
		for _, alphabet := range components[1:] {
			hijaiyas[alphabet] = append(hijaiyas[alphabet], arabic)

			length := len(alphabet)
			ending := alphabet[length-1]
			if ending == 'a' || ending == 'i' || ending == 'o' || ending == 'u' {
				alphabet = alphabet[:length-1] + alphabet[:length-1] + alphabet[length-1:]
			} else {
				alphabet += alphabet
			}
			hijaiyas[alphabet] = append(hijaiyas[alphabet], arabic)
			length = len(alphabet)
			if length > alphabetMaxLen {
				alphabetMaxLen = length
			}
		}
	}

	q.BuildIndex()
	return Transliteration{hijaiyas, alphabetMaxLen, q}
}

// Encode returns arabic encodings of given string using Transliteration t.
func (t Transliteration) Encode(s string) []string {
	var memo = make(map[string][]string)
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	results := []string{}
	for _, result := range t.quranize(s, memo) {
		if len(t.q.Locate(result)) > 0 {
			results = appendUniq(results, result)
		}
	}
	return results
}

func (t Transliteration) quranize(s string, memo map[string][]string) []string {
	if s == "" {
		return base
	}

	if cache, ok := memo[s]; ok {
		return cache
	}

	kalimas := []string{}
	l := len(s)
	for width := 1; width <= t.alphabetMaxLen && width <= l; width++ {
		if tails, ok := t.hijaiyas[s[l-width:]]; ok {
			heads := t.quranize(s[:l-width], memo)
			for _, combination := range combine(heads, tails) {
				if t.q.exists(combination) {
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
