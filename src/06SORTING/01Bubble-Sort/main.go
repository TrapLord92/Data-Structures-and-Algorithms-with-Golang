package main

import "fmt"

/*Bubble-Sort
Bubble-Sort is the slowest algorithm for sorting. It is easy to implement and used when data is small.
In Bubble-Sort, we compare each pair of adjacent values. We want to sort values in increasing order so if the second
value is less than the first value then we swap these two values. Otherwise, we will go to the next pair. Thus, largest
values bubble to the end of the list.
After the first pass, the largest value will be in the rightmost position. We will have N number of passes to get the list
completely sorted.
Example 6.1*/
func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				/* Swapping */
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data, more)
	fmt.Println(data)
}
func more(value1 int, value2 int) bool {
	return value1 > value2
}

/*Analysis:
· The outer for loops represents the number of swaps that are done for comparison of data.
· The inner loop is actually used to do the comparison of data. At the end of each inner loop iteration, the largest
value is moved to the end of the list. In the first iteration the largest value, in the second iteration the second largest
and so on.
· more() function is used for comparison which means when the value of the first argument is greater than the value
of the second argument then perform a swap. By this we are sorting in increasing order if we have, the less()
function in place of more() then list will be sorted in decreasing order.
Complexity Analysis:
Each time the inner loop execute for (n-1), (n-2), (n-3)…(n-1) + (n-2) + (n-3) + ..... + 3 + 2 + 1 = n(n-1)/2
Worst case performance O(n2)
Average case performance O(n2)
Space Complexity O(1) as we need only one temp variable
Stable Sorting*/

// Modified (improved) Bubble-Sort
// When there is no more swap in one pass of the outer loop. It indicates that all the elements are already in order so we
// should stop sorting. This sorting improvement in Bubble-Sort is extremely useful when we know that, except few
// elements rest of the list is already sorted.
// Example 6.2
func BubbleSort2(arr []int, comp func(int, int) bool) {
	size := len(arr)
	swapped := 1
	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
}

// By applying this improvement, best case performance of this algorithm when a list is nearly sorted, is improved. In this
// case we just need one single pass and best case complexity is O(n)
// Complexity Analysis:
// Worst case performance O(n2)
// Average case performance O(n2)
// Space Complexity O(1)
// Adaptive: When list is nearly sorted O(n)
// Stable Sorting Yes
