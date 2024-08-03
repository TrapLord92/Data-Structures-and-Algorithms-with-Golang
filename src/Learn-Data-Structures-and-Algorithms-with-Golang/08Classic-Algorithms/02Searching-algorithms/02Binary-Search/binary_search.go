// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	"sort"
)

/*Binary
The binary search algorithm compares the input value to the middle element of the sorted
collection. If the values are not equal, the half in which the element is not found is
eliminated. The search continues on the remaining half of the collection. The time
complexity of this algorithm is in the order of O(log n).
The following code snippet shows an implementation of the binary search algorithm using
the sort.Search function from the sort package. The main method initializes the
elements array and invokes the sort.Search function to find an integer element:*/
// main method
func main() {
	var elements []int
	elements = []int{1, 3, 16, 10, 45, 31, 28, 36, 45, 75}
	var element int
	element = 36

	var i int

	i = sort.Search(len(elements), func(i int) bool { return elements[i] >= element })
	if i < len(elements) && elements[i] == element {
		fmt.Printf("found element %d at index %d in %v\n", element, i, elements)
	} else {
		fmt.Printf("element %d not found in %v\n", element, elements)
	}
}
