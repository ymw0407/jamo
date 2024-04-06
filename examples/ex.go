package main

import (
	"fmt"

	"github.com/ymw0407/golang-jamo/pkg/jamo"
	"github.com/ymw0407/golang-jamo/pkg/options"
)

func main() {
	fmt.Println(jamo.DecomposeHangeul("얘를롦놈", options.Jamo()))
}
