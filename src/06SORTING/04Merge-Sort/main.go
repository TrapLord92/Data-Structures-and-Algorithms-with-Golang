package main

import "fmt"

// Merge-Sort
// Merge sort divide the input into half recursive in each step. It sort the two parts separately recursively and finally
// combine the result into final sorted output.
// Example 6.6:
func MergeSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	tempArray := make([]int, size)
	mergeSrt(arr, tempArray, 0, size-1, comp)
}
func merge(arr []int, tempArray []int, lowerIndex int, middleIndex int, upperIndex int, comp func(int, int) bool) {
	lowerStart := lowerIndex
	lowerStop := middleIndex
	upperStart := middleIndex + 1
	upperStop := upperIndex
	count := lowerIndex
	for lowerStart <= lowerStop && upperStart <= upperStop {
		if comp(arr[lowerStart], arr[upperStart]) == false {
			tempArray[count] = arr[lowerStart]
			lowerStart++
		} else {
			tempArray[count] = arr[upperStart]
			upperStart++
		}
		count++
	}
	for lowerStart <= lowerStop {
		tempArray[count] = arr[lowerStart]
		count++
		lowerStart++
	}
	for upperStart <= upperStop {
		tempArray[count] = arr[upperStart]
		count++
		upperStart++
	}
	for i := lowerIndex; i <= upperIndex; i++ {
		arr[i] = tempArray[i]
	}
}
func mergeSrt(arr []int, tempArray []int, lowerIndex int, upperIndex int, comp func(int, int) bool) {
	if lowerIndex >= upperIndex {
		return
	}
	middleIndex := (lowerIndex + upperIndex) / 2
	mergeSrt(arr, tempArray, lowerIndex, middleIndex, comp)
	mergeSrt(arr, tempArray, middleIndex+1, upperIndex, comp)
	merge(arr, tempArray, lowerIndex, middleIndex, upperIndex, comp)
}
func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	MergeSort(data, more)
	fmt.Println(data)
}
func more(value1 int, value2 int) bool {
	return value1 > value2
}

// · The Time Complexity of Merge-Sort is O(nlogn) in all 3 cases (best, average and worst) as Merge-Sort always
// divides the list into two halves and take linear time to merge two halves.
// · It requires the equal amount of additional space as the unsorted list. Hence, it is not at all recommended for
// searching large unsorted lists.
// · It is the best Sorting technique for sorting Linked Lists.
// Complexity Analysis:
// Worst Case Time Complexity O(nlogn)
// Best Case Time Complexity O(nlogn)
// Average Time Complexity O(nlogn)
// Space Complexity O(n)
// Stable Sorting Yes
