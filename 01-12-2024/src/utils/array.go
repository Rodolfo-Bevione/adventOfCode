package utils

func Quicksort(arr) {

}

func quicksortIteration(arr []int, low int, high int) {
	// find a pivot
	var pivot = arr[len(arr)-1]

	var i int = 0;
	var j int = -1;
	var temp int = nil

	for i := 0; i < len(arr); i++ {

	}
}

func quicksortPartition(arr, low, high) {

}

func Quicksort(arr1) {
function quicksort(arr, low, high)
    if low < high:
        pi = partition(arr, low, high)
        quicksort(arr, low, pi - 1)
        quicksort(arr, pi + 1, high)

function partition(arr, low, high):
    pivot = arr[high]
    i = low - 1
    for j = low to high - 1:
        if arr[j] < pivot:
            i = i + 1
            swap arr[i] with arr[j]
    swap arr[i + 1] with arr[high]
    return (i + 1)
}