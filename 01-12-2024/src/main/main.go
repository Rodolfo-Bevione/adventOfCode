package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func computeDifferences(first, second []int) []int {
	if len(first) != len(second) {
		panic("Le lunghezze sono diverse!")
	}
	result := make([]int, len(first))
	for i, _ := range first {
		difference := first[i] - second[i]
		if difference < 0 {
			difference = -1 * difference
		}
		result[i] = difference
	}
	return result
}

func mergeSortInternal(arr []int, left, right int) {
	if left < right {
		center := (left + right) / 2
		mergeSortInternal(arr, left, center)
		mergeSortInternal(arr, center+1, right)
		merge(arr, left, center, right)
	}
}

func merge(arr []int, left, center, right int) {
	leftIndex := left
	rightIndex := center + 1
	tempArrayIndex := 0
	tempArray := make([]int, right-left+1)

	for leftIndex <= center && rightIndex <= right {
		if arr[leftIndex] < arr[rightIndex] {
			tempArray[tempArrayIndex] = arr[leftIndex]
			leftIndex++
			tempArrayIndex++
		} else {
			tempArray[tempArrayIndex] = arr[rightIndex]
			rightIndex++
			tempArrayIndex++
		}
	}

	for leftIndex <= center {
		tempArray[tempArrayIndex] = arr[leftIndex]
		leftIndex++
		tempArrayIndex++
	}

	for rightIndex <= right {
		tempArray[tempArrayIndex] = arr[rightIndex]
		rightIndex++
		tempArrayIndex++
	}

	for z := left; z <= right; z++ {
		arr[z] = tempArray[z-left]
	}
}

func mergeSort(arr []int) []int {
	mergeSortInternal(arr, 0, len(arr)-1)
	return arr
}

func main() {
	start := time.Now()
	// leggere le due liste
	file, err := os.Open("./test1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	locationsA := []int{}
	locationsB := []int{}
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

	sortedA := mergeSort(locationsA)
	sortedB := mergeSort(locationsB)
	result := computeDifferences(sortedA, sortedB)
	total := 0
	for _, n := range result {
		total += n
	}
	fmt.Println(total)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
