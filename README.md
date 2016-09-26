# mbconv - multibyte converter for golang

[![Build Status](https://travis-ci.org/kaneshin/mbconv.svg?branch=master)](https://travis-ci.org/kaneshin/mbconv)
[![codecov](https://codecov.io/gh/kaneshin/mbconv/branch/master/graph/badge.svg)](https://codecov.io/gh/kaneshin/mbconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaneshin/mbconv)](https://goreportcard.com/report/github.com/kaneshin/mbconv)

## Installation

```shell
go get -u -d github.com/kaneshin/mbconv
```

## Usage

### Example

```go
package main

import (
	"fmt"

	"github.com/kaneshin/mbconv"
)

func main() {
	fmt.Println(mbconv.HiraganaToKatakana("ぁあぃいぅうぇえぉおかがきぎく"))
	// => ァアィイゥウェエォオカガキギク

	fmt.Println(mbconv.KatakanaToHiragana("ァアィイゥウェエォオカガキギク"))
	// => ぁあぃいぅうぇえぉおかがきぎく

	fmt.Println(mbconv.HalfWidthToFullWidth("Hello World"))
	// => Ｈｅｌｌｏ　Ｗｏｒｌｄ

	fmt.Println(mbconv.FullWidthToHalfWidth("Ｈｅｌｌｏ　Ｗｏｒｌｄ"))
	// => Hello World
}
```

## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
