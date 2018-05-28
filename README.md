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
  quranize := quranize.NewDefaultQuranize()
  fmt.Println(quranize.Encode("alhamdulillah"))
  // Output: [الحمد لله]
}
```

## Related Project

https://github.com/alpancs/quranize-service
