package main

func computePartTwo(first, second []int) int {
	if len(first) != len(second) {
		panic("Le lunghezze sono diverse!")
	}

	// run through every element of array A
	// if element X appears in list B, increment a counter
	// At the end, multiply X by the counter and increase the total score
	result := 0
	for i := 0; i < len(first); i++ {
		elementFromFirstArray := first[i]
		howMany := 0
		for _, elementFromSecondArray := range second {
			if elementFromFirstArray == elementFromSecondArray {
				howMany++
			}
		}
		result = result + (howMany * elementFromFirstArray)
	}
	return result
}
