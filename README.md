# goconv - multibyte converter for golang

[![Build Status](https://travis-ci.org/kaneshin/goconv.svg?branch=master)](https://travis-ci.org/kaneshin/goconv)
[![codecov](https://codecov.io/gh/kaneshin/goconv/branch/master/graph/badge.svg)](https://codecov.io/gh/kaneshin/goconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaneshin/goconv)](https://goreportcard.com/report/github.com/kaneshin/goconv)

## Installation

```shell
go get -u -d github.com/kaneshin/goconv
```

## Usage

### Example

```go
package main

import (
	"fmt"

	"github.com/kaneshin/goconv"
)

func main() {
	fmt.Println(goconv.HiraganaToKatakana("ぁあぃいぅうぇえぉおかがきぎく"))
	// => ァアィイゥウェエォオカガキギク

	fmt.Println(goconv.KatakanaToHiragana("ァアィイゥウェエォオカガキギク"))
	// => ぁあぃいぅうぇえぉおかがきぎく

	fmt.Println(goconv.HalfWidthToFullWidth("Hello World"))
	// => Ｈｅｌｌｏ　Ｗｏｒｌｄ

	fmt.Println(goconv.FullWidthToHalfWidth("Ｈｅｌｌｏ　Ｗｏｒｌｄ"))
	// => Hello World
}
```

## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
