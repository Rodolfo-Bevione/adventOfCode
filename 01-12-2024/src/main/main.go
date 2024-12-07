package main

import (
	"adventOfCode/01/crow8279-feat-ByIvan/src/utils"
	"fmt"
	"log"
	"path"
	"time"
)

func main() {
	start := time.Now()
	// leggere le due liste

	locationsA, locationsB := scanFile(
		path.Join("..", "input", "test.txt"),
	)

	sortedA := utils.MergeSort(locationsA)
	sortedB := utils.MergeSort(locationsB)

	resultPartOne := computePartOne(sortedA, sortedB)

	fmt.Println(fmt.Sprintf("Part 1  %d", resultPartOne))
	resultPartTwo := computePartTwo(sortedA, sortedB)
	fmt.Println(fmt.Sprintf("Part 2  %d", resultPartTwo))

	elapsed := time.Since(start)
	log.Printf("Time elapsed %s", elapsed)

}
