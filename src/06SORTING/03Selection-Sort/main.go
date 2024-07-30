package main

/*Selection-Sort
Selection-Sort searches the whole unsorted array and put the largest value at the end of it. This algorithm is having the
same Time Complexity, but performs better than both bubble and Insertion-Sort as less number of comparisons
required. The sorted list is created backward in Selection-Sort.*/

// func SelectionSort(arr []int) {
// 	size := len(arr)
// 	var i, j, max, temp int
// 	for i = 0; i < size-1; i++ {
// 		max = 0
// 		for j = 1; j < size-1-i; j++ {
// 			if arr[j] > arr[max] {
// 				max = j
// 			}
// 		}
// 		arr[size-1-i], arr[max] = arr[max], arr[size-1-i]
// 	}
// }

/*Analysis:
· The outer loop decide the number of times the inner loop will iterate. For an input of N elements, the inner loop
will iterate N number of times.
· In each iteration of the inner loop, the largest value is calculated and is placed at the end of the list.
· This is the final replacement of the maximum value to the proper location. The sorted list is created backward.
Complexity Analysis:
Worst Case Time Complexity O(n2)
Best Case Time Complexity O(n2)
Average case Time Complexity O(n2)
Space Complexity O(1)
Stable Sorting No
The same algorithm can be implemented by creating the sorted list in the front of the list.*/

// func SelectionSort2(arr []int) {
// 	size := len(arr)
// 	var i, j, min, temp int
// 	for i = 0; i < size-1; i++ {
// 		min = i
// 		for j = i + 1; j < size; j++ {
// 			if arr[j] < arr[min] {
// 				min = j
// 			}
// 		}
// 		arr[i], arr[min] = arr[min], arr[i]
// 	}
// }
