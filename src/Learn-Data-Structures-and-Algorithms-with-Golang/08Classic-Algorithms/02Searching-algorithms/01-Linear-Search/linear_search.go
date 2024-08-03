// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*Linear
The linear search method finds a given value within a collection by sequentially checking
every element in the collection. The time complexity of the linear search algorithm is O(n).
The binary search algorithm and hash tables perform better than this search algorithm.
The implementation of the linear search method is shown in the following code snippet.*/

/*The LinearSearch function takes an array of integer elements and findElement int as parameters.
The function returns a Boolean true if the findElement is found; otherwise, it
returns false:*/
// Linear Search method
func LinearSearch(elements []int, findElement int) bool {
	var element int
	for _, element = range elements {
		if element == findElement {
			return true
		}
	}
	return false
}

// main method
/*The main method initializes the array of integer elements and invokes the LinearSearch
method by passing an integer that needs to be found, as follows:*/
func main() {
	var elements []int
	elements = []int{15, 48, 26, 18, 41, 86, 29, 51, 20}
	fmt.Println(LinearSearch(elements, 48))
}
