// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
)

/*The shell sort algorithm sorts a pair of elements that are not in sequence in a collection. The
distance between the elements to be compared is decreased sequentially. This algorithm
performs more operations and has a greater cache miss ratio than the quick sort algorithm.
In the following code, we can see the implementation of the shell sort algorithm.*/

/*The
ShellSorter function takes an integer array as a parameter and sorts it:*/
// shell sorter method
func ShellSorter(elements []int) {
	var (
		n         = len(elements)
		intervals = []int{1}
		k         = 1
	)

	for {
		var interval int
		interval = power(2, k) + 1
		if interval > n-1 {
			break
		}
		intervals = append([]int{interval}, intervals...)
		k++
	}
	var interval int
	for _, interval = range intervals {
		var i int
		for i = interval; i < n; i += interval {
			var j int
			j = i
			for j > 0 {
				if elements[j-interval] > elements[j] {
					elements[j-interval], elements[j] = elements[j], elements[j-interval]
				}
				j = j - interval
			}
		}
	}
}

// power function
/*The power method
The power method takes exponent and index as parameters and returns the power of the
exponent to the index, as follows:*/
func power(exponent int, index int) int {
	var power int
	power = 1
	for index > 0 {
		if index&1 != 0 {
			power *= exponent
		}
		index >>= 1
		exponent *= exponent
	}
	return power
}

// main method
func main() {
	var elements []int
	elements = []int{34, 202, 13, 19, 6, 5, 1, 43, 506, 12, 20, 28, 17, 100, 25, 4, 5, 97, 1000, 27}
	ShellSorter(elements)
	fmt.Println(elements)
}
