// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Selection Sorter method
/*Selection
Selection sort is an algorithm that divides the input collection into two fragments. This
sublist of elements is sorted by swapping the smallest or largest element from the left of the
list to the right. The algorithm is of the order O(n2). This algorithm is inefficient for large
collections, and it performs worse than the insertion sort algorithm.
The following code shows the implementation of the SelectionSorter function, which
takes the collection to be sorted:*/

func SelectionSorter(elements []int) {

	var i int
	for i = 0; i < len(elements)-1; i++ {
		var min int
		min = i
		var j int
		for j = i + 1; j <= len(elements)-1; j++ {
			if elements[j] < elements[min] {
				min = j
			}
		}
		swap(elements, i, min)
	}
}

// swap method
/*The swap method
The swap method takes the elements array and the i and j indices as parameters. The
method swaps the element at position i with the element at position j, as shown here:*/
func swap(elements []int, i int, j int) {
	var temp int
	temp = elements[j]
	elements[j] = elements[i]
	elements[i] = temp
}

// main method
func main() {
	var elements []int
	elements = []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("Before Sorting ", elements)
	SelectionSorter(elements)
	fmt.Println("After Sorting", elements)
}
