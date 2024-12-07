package main

func computePartTwo(first, second []int, ch chan int) {
	if len(first) != len(second) {
		panic("Le lunghezze sono diverse!")
	}
	secondArrayMapped := make(map[int]int)

	for _, elementFromSecondArray := range second {
		howMany := secondArrayMapped[elementFromSecondArray]

		secondArrayMapped[elementFromSecondArray] = howMany + 1
	}

	result := 0
	for i := 0; i < len(first); i++ {
		func() {
			elementFromFirstArray := first[i]
			result = result + (elementFromFirstArray * secondArrayMapped[elementFromFirstArray])
		}()
	}
	ch <- result
	close(ch)
}
