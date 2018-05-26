package quran

import (
	"strings"
)

var (
	hijaiyas       map[string][]string
	alphabetMaxLen int

	base = []string{""}
)

// Returns arabic encodings of given string.
func Encode(s string) []string {
	var memo = make(map[string][]string)
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	results := []string{}
	for _, result := range quranize(s, memo) {
		if len(Locate(result)) > 0 {
			results = appendUniq(results, result)
		}
	}
	return results
}

// Returns locations of s (quran kalima), matching the whole word.
func Locate(s string) []Location {
	harfs := []rune(s)
	node := root
	for _, harf := range harfs {
		node = getChild(node.children, harf)
		if node == nil {
			return zeroLocs
		}
	}
	return node.locations
}

func quranize(s string, memo map[string][]string) []string {
	if s == "" {
		return base
	}

	if cache, ok := memo[s]; ok {
		return cache
	}

	kalimas := []string{}
	l := len(s)
	for width := 1; width <= alphabetMaxLen && width <= l; width++ {
		if tails, ok := hijaiyas[s[l-width:]]; ok {
			heads := quranize(s[:l-width], memo)
			for _, combination := range combine(heads, tails) {
				if exists(combination) {
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

func exists(s string) bool {
	harfs := []rune(s)
	node := root
	for _, harf := range harfs {
		node = getChild(node.children, harf)
		if node == nil {
			return false
		}
	}
	return true
}

func appendUniq(results []string, newResult string) []string {
	for _, result := range results {
		if result == newResult {
			return results
		}
	}
	return append(results, newResult)
}
