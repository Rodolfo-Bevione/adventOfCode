package utils

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

func MergeSort(arr []int) []int {
	mergeSortInternal(arr, 0, len(arr)-1)
	return arr
}
