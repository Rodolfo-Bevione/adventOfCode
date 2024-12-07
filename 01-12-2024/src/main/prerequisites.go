package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanFile(path string) (locationsA []int, locationsB []int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text() // this is a whole line
		left := 0
		right := 0
		_, err := fmt.Sscanf(text, "%d   %d", &left, &right)
		if err != nil {
			panic(err)
		}
		locationsA = append(locationsA, left)
		locationsB = append(locationsB, right)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return locationsA, locationsB
}
