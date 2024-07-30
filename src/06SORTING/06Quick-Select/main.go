package main

import "fmt"

/*Quick Select
Quick select algorithm is used to find the element, which will be at the Kth position when the list will be sorted without
actually sorting the whole list. Quick select is very similar to Quick-Sort in place of sorting the whole list we just
ignore the one-half of the list at each step of Quick-Sort and just focus on the region of list on which we are interested.
Example 6.8:*/
func QuickSelect(arr []int, key int) int {
	size := len(arr)
	QuickSelectUtil(arr, 0, size-1, key)
	return arr[key-1]
}
func QuickSelectUtil(arr []int, lower int, upper int, key int) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper
	for lower < upper {
		for arr[lower] <= pivot && lower < upper {
			lower++
		}
		for arr[upper] > pivot && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}
	swap(arr, upper, start)                   // upper is the pivot position
	QuickSelectUtil(arr, start, upper-1, key) // pivot -1 is the upper for left sub array.
	QuickSelectUtil(arr, upper+1, stop, key)  // pivot + 1 is the lower for right sub array.
}
func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}
func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Println(QuickSelect(data, 5))
}

/*Complexity Analysis:
Worst Case Time Complexity O(n2)
Best Case Time Complexity O(logn)
Average Time Complexity O(logn)
Space Complexity O(nlogn)*/
