package main

import (
	"fmt"
	"sort"

)

/*
Problems in Searching
Print Duplicates in List
Given a list of n numbers, print the duplicate elements in the list.
First approach: Exhaustive search or Brute force, for each element in list find if there is some other element with the
same value. This is done using two for loop, first loop to select the element and second loop to find its duplicate entry.
Example 5.5
*/
func printRepeating(data []int) {
	size := len(data)
	fmt.Print("Repeating elements are : ")
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				fmt.Print(" ", data[i])
			}
		}
	}
	fmt.Println()
}

/*
The Time Complexity is O(n2) and Space Complexity is O(1)
Second approach: Sorting, Sort all the elements in the list and after this in a single scan, we can find the duplicates.
Example 5.6
*/
func printRepeating2(data []int) {
	size := len(data)
	sort.Ints(data) // Sort(data,size)
	fmt.Print("Repeating elements are : ")
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			fmt.Print(" ", data[i])
		}
	}
	fmt.Println()
}

/*
Sorting algorithms take O(n log n) time and single scan take O(n) time.
The Time Complexity of an algorithm is O(n log n) and Space Complexity is O(1)
Third approach: Hash-Table, using Hash-Table, we can keep track of the elements we have already seen and we can
find the duplicates in just one scan.
Example 5.7
*/

// func printRepeating3(data []int) {
// 	s := make(Set)
// 	size := len(data)
// 	fmt.Print("Repeating elements are : ")
// 	for i := 0; i < size; i++ {
// 		if s.Find(data[i]) {
// 			fmt.Print(" ", data[i])
// 		} else {
// 			s.Add(data[i])
// 		}
// 	}
// 	fmt.Println()
// }

/*Hash-Table insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time. Space
Complexity is also O(n)
Forth approach: Counting, this approach is only possible if we know the range of the input. If we know that, the
elements in the list are in the range 0 to n-1. We can reserve a list of length n and when we see an element, we can
increase its count. In just one single scan, we know the duplicates. If we know the range of the elements, then this is the
fastest way to find the duplicates.*/

func printRepeating4(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)
	for i := 0; i < size; i++ {
		count[i] = 0
	}
	fmt.Println("Repeating elements are : ")
	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Println(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
	fmt.Println()
}

/*Counting approach just uses a list so insert and find take constant time O(1) so the total Time Complexity of the
algorithm is O(n) time. Space Complexity for creating count list is also O(n)*/
