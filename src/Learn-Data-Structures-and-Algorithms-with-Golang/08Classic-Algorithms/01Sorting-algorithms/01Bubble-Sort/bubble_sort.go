// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
)

//bubble Sorter method
/*The bubble sort algorithm is a sorting algorithm that compares a pair of neighboring
elements and swaps them if they are in the wrong order. The algorithm has a complexity of
O(n2), where n is the number of elements to be sorted. The smallest or greatest value
bubbles up to the top of the collection, or the smallest or greatest sinks to the bottom
(depending on whether you're sorting into ascending or descending order).
The following code snippet shows the implementation of the bubble sort algorithm. The
bubbleSorter function takes an integer array and sorts the array's elements in ascending
order.*/
func bubbleSorter(integers [11]int) {

	var num int
	num = 11
	var isSwapped bool
	isSwapped = true
	for isSwapped {
		isSwapped = false
		var i int
		for i = 1; i < num; i++ {
			if integers[i-1] > integers[i] {

				var temp = integers[i]
				integers[i] = integers[i-1]
				integers[i-1] = temp
				isSwapped = true
			}
		}
	}
	fmt.Println(integers)
}

// main method
/*The main method initializes the array's integers and invokes the bubbleSorter function,
as follows:*/
func main() {
	var integers [11]int = [11]int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("Bubble Sorter")
	bubbleSorter(integers)

}
