package main

import (
	"adventOfCode/01/crow8279-feat-ByIvan/src/bufo"
	"adventOfCode/01/crow8279-feat-ByIvan/src/serio"
	"fmt"
)

var locationsA = [6]int{
	3,
	4,
	2,
	1,
	3,
	3,
}
var locationsB = [6]int{
	4,
	3,
	5,
	3,
	9,
	3,
}

func main() {
	fmt.Println("hello world")
	bufo.Bafo()
	serio.Serio()
}
