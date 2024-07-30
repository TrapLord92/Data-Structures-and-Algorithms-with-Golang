package main

import "fmt"

/*Quick-Sort
Quick sort is also a recursive algorithm.
· In each step, we select a pivot (let us say first element of list).
· Then we traverse the rest of the list and copy all the elements of the list which are smaller than the pivot to the left
side of list
· We copy all the elements of the list, which are greater than pivot to the right side of the list. Obviously, the pivot is
at its sorted position.
· Then we sort the left and right subarray separately.
· When the algorithm returns the whole list is sorted.
Example 6.7:*/
func QuickSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	quickSortUtil(arr, 0, size-1, comp)
}
func quickSortUtil(arr []int, lower int, upper int, comp func(int, int) bool) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper
	for lower < upper {
		for comp(arr[lower], pivot) == false && lower < upper {
			lower++
		}
		for comp(arr[upper], pivot) && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}
	swap(arr, upper, start)                  // upper is the pivot position
	quickSortUtil(arr, start, upper-1, comp) // pivot -1 is the upper for left sub array.
	quickSortUtil(arr, upper+1, stop, comp)  // pivot + 1 is the lower for right sub array.
}
func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}
func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	QuickSort(data, more)
	fmt.Println(data)
}
func more(value1 int, value2 int) bool {
	return value1 > value2
}

/*· The space required by Quick-Sort is very less, only O(nlogn) additional space is required.
· Quicksort is not a stable sorting technique. It can reorder elements with identical keys.
Complexity Analysis:
Worst Case Time Complexity O(n2)
Best Case Time Complexity O(nlogn)
Average Time Complexity O(nlogn)
Space Complexity O(nlogn)
Stable Sorting*/
