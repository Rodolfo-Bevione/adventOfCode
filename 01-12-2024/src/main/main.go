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

	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go computePartOne(sortedA, sortedB, ch1)
	go computePartTwo(sortedA, sortedB, ch2)

	i := 2
	for i > 0 {
		select {
		case result1 := <-ch1:
			fmt.Println("First:")
			fmt.Println(result1)
			i--
		case result2 := <-ch2:
			fmt.Println("Second:")
			fmt.Println(result2)
			i--
		}
	}

	elapsed := time.Since(start)
	log.Printf("Time elapsed %s", elapsed)

}
