[![GoDoc](https://godoc.org/github.com/alpancs/quranize?status.svg)](https://godoc.org/github.com/alpancs/quranize)
[![Build Status](https://travis-ci.org/alpancs/quranize.svg?branch=master)](https://travis-ci.org/alpancs/quranize)
[![codecov](https://codecov.io/gh/alpancs/quranize/branch/master/graph/badge.svg)](https://codecov.io/gh/alpancs/quranize)

# Quranize

Quranize provides Alquran in Go representation.
Original source of Alquran is taken from http://tanzil.net in XML format.

Quranize can transform alphabet into arabic using fast and efficient algorithm:
suffix-tree for indexing and dynamic programming for parsing.

## Documentation

https://godoc.org/github.com/alpancs/quranize

## Example

Here is a minimal example.
```go
package main

import (
  "fmt"

  "github.com/alpancs/quranize"
)

func main() {
  q := quranize.NewDefaultQuranize()
  quran := quranize.NewQuranSimpleEnhanced()

  encodeds := q.Encode("alhamdulillah hirobbil 'alamin")
  fmt.Println(encodeds)
  // Output: [الحمد لله رب العالمين]

  locations := q.Locate(encodeds[0])
  fmt.Println(locations)
  // Output: [{1 2 0} {10 10 10} {39 75 13} {40 65 10}]

  suraName, _ := quran.GetSuraName(locations[0].Sura)
  fmt.Println(suraName)
  // Output: الفاتحة

  aya, _ := quran.GetAya(locations[0].Sura, locations[0].Aya)
  fmt.Println(aya)
  // Output: الْحَمْدُ لِلَّهِ رَبِّ الْعَالَمِينَ
}
```

## Related Project

https://github.com/alpancs/quranize-service
