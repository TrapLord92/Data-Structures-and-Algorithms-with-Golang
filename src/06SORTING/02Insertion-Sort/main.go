package main

import "fmt"

// Insertion-Sort
// Insertion-Sort Time Complexity is O( ) which is same as Bubble-Sort but perform a bit better than it. It is the way we
// arrange our playing cards. We keep a sorted subarray. Each value is inserted into its proper position in the sorted subarray
// in the left of it.
// Example 6.3
func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}
func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	InsertionSort(data, more)
	fmt.Println(data)
}
func more(value1 int, value2 int) bool {
	return value1 > value2
}

// Analysis:
// · The outer loop is used to pick the value we want to insert into the sorted list in left.
// · The value we want to insert we have picked and saved in a temp variable.
// · The inner loop is doing the comparison using the more() function. The values are shifted to the right until we
// find the proper position of the temp value for which we are doing this iteration.
// · Finally, the value is placed into the proper position. In each iteration of the outer loop, the length of the sorted
// list increase by one. When we exit the outer loop, the whole list is sorted.
// Complexity Analysis:
// Worst case Time Complexity O(n2)
// Best case Time Complexity O(n)
// Average case Time Complexity O(n2)
// Space Complexity O(1)
// Stable sorting Yes
