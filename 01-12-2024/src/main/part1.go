package main

func computePartOne(first, second []int, ch chan int) {
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

	total := 0
	for _, n := range result {
		total += n
	}
	ch <- total
	close(ch)
}
