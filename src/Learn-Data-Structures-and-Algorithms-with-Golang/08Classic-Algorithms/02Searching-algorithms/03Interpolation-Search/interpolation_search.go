// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and os package
import (
	"fmt"
)

/*The interpolation search algorithm searches for the element in a sorted collection. The
algorithm finds the input element at an estimated position by diminishing the search space
before or after the estimated position. The time complexity of the search algorithm is of the
order O(log log n).
The following code snippet implements the interpolation search algorithm. */

/*The
InterpolationSearch function takes the array of integer elements and the integer
element to be found as parameters. The function finds the element in the collection and
returns the Boolean and the index for the found element:*/
//interpolation search method
func InterpolationSearch(elements []int, element int) (bool, int) {
	var mid int
	var low int
	low = 0
	var high int
	high = len(elements) - 1

	for elements[low] < element && elements[high] > element {
		mid = low + ((element-elements[low])*(high-low))/(elements[high]-elements[low])

		if elements[mid] < element {
			low = mid + 1
		} else if elements[mid] > element {
			high = mid - 1
		} else {
			return true, mid
		}
	}

	if elements[low] == element {
		return true, low
	} else if elements[high] == element {
		return true, high
	} else {
		return false, -1
	}

	return false, -1
}

/*The main method initializes the array of integer elements and invokes the
InterpolationSearch method with the elements array and the element parameters, as
follows*/
// main method
func main() {

	var elements []int

	elements = []int{2, 3, 5, 7, 9}

	var element int

	element = 7

	var found bool

	var index int
	found, index = InterpolationSearch(elements, element)

	fmt.Println(found, "found at", index)

}
