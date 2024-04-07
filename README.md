# golang-jamo

Hangeul character syllable decomposing/composing Go library

<a href="/LICENSE"><img src="https://img.shields.io/github/license/ymw0407/jamo" alt="License" /></a>
<a href="https://github.com/ymw0407/jamo/graphs/contributors" target="_blank"><img src="https://img.shields.io/github/contributors-anon/ymw0407/jamo" alt="contributors" /></a>
<a href="https://github/ymw0407/jamo"><img src="https://img.shields.io/github/last-commit/ymw0407/jamo" alt="your repo's last-commit" /></a>
<a href="https://github/ymw0407/jamo"><img src="https://img.shields.io/github/issues/ymw0407/jamo" alt="your repo's issues" /></a>
<a href="https://codecov.io/gh/ymw0407/jamo" ><img src="https://codecov.io/gh/ymw0407/jamo/graph/badge.svg?token=1PM1ZYQOMA"/></a>
<a href="https://github/ymw0407/jamo"><img src="https://img.shields.io/github/stars/ymw0407/jamo" alt="your repo's stars" /></a>
<a href="https://github/ymw0407/jamo"><img src="https://img.shields.io/github/forks/ymw0407/jamo" alt="your repo's forks" /></a>
<a href="https://github/ymw0407/jamo"><img src="https://img.shields.io/github/watchers/ymw0407/jamo" alt="your repo's watchers" /></a>

## Description

Golang-jamo is a Go Hangeul syllable decomposition/composition library for working with Hangul characters and jamo. <br><br>
This library support Hangeul syllable decomposition with several options. (e.g. Detailed Jamo decomposition, Qwerty keyboard layout decomposition) <br>

## Installation

```bash
go get -u github.com/ymw0407/jamo
```

## Usage/Examples

```go
package main

import (
	"fmt"

	"github.com/ymw0407/jamo/pkg/jamo"
	"github.com/ymw0407/jamo/pkg/options"
)

func main() {
	fmt.Println(
		jamo.DecomposeHangeul("한글 테스트 예시!"),
	) // "ㅎㅏㄴㄱㅡㄹ ㅌㅔㅅㅡㅌㅡ ㅇㅖㅅㅣ!"
	fmt.Println(
		jamo.DecomposeHangeul("한글 테스트 예시!", options.Jamo().SetComplexConsonants(true).SetDiphthong(true).SetTenseConsonants(true)),
	) // "ㅎㅏㄴㄱㅡㄹ ㅌㅓㅣㅅㅡㅌㅡ ㅇㅕㅣㅅㅣ!"
	fmt.Println(
		jamo.DecomposeHangeul("한글 테스트 예시!", options.Qwerty().SetShiftOption(options.QwertyShiftOption1)),
	) // "ㅎㅏㄴㄱㅡㄹ ㅌㅔㅅㅡㅌㅡ ㅇㅕㅣㅅㅣ!"
	fmt.Println(
		jamo.DecomposeHangeul("한글 테스트 예시!", *options.Qwerty().SetShiftOption(options.QwertyShiftOption2)),
	) // "ㅎㅏㄴㄱㅡㄹ ㅌㅔㅅㅡㅌㅡ ㅇㅔㅔㅅㅣ!"

	fmt.Println(
		jamo.ComposeHangeul("ㅎㅏㄴㄱㅡㄹ"),
	) // ["한글"], nil
	fmt.Println(
		jamo.ComposeHangeul("ㅎㅏㄴ ㄱㅡㄹ"), // any other characters except hangeul syllable is unavailable
	) // [], JamoError.ErrImpossibleToCompose
	fmt.Println(
		jamo.ComposeHangeul("english"), // any other characters except hangeul syllable is unavailable
	) // [], JamoError.ErrImpossibleToCompose
	fmt.Println(
		jamo.ComposeHangeul(
			jamo.DecomposeHangeul("한글테스트예시", *options.Qwerty().SetShiftOption(options.QwertyShiftOption2)),
		), // qwerty decomposition, jamo decomposition available to compose
	) // ["한글테스트예시"], nil
	fmt.Println(
		jamo.ComposeHangeul(
			jamo.DecomposeHangeul("ㄱㅡㄹㄱㄱㅣ", *options.Qwerty().SetShiftOption(options.QwertyShiftOption2)),
		), // if it can be combined into several other characters, it returns all of them.
	) // ["긁기", "글끼"], nil
}
```

## Contributers

<a href="https://github.com/ymw0407/jamo/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ymw0407/jamo" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

## License

jamo is released under Apache License 2.0.
See the [LICENSE file]("./LICENSE") for details.
