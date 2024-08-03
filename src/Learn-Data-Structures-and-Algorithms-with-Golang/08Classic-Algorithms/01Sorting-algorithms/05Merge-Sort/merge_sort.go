// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
	"math/rand"
	"time"
)

/*Merge
The merge sort algorithm is a comparison-based method that was invented by John Von
Neumann. Each element in the adjacent list is compared for sorting. The performance of the
algorithm is in the order of O(n log n). This algorithm is the best algorithm for sorting a
linked list.The following code snippet demonstrates the merge sort algorithm. */

// create array
/*The createArray
function takes num int as a parameter and returns an integer, array, that consists of
randomized elements:*/
func createArray(num int) []int {

	array := make([]int, num)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < num; i++ {
		array[i] = r.Intn(99999) - r.Intn(99999)
	}
	return array
}

// MergeSorter algorithm
/*MergeSorter method
The MergeSorter method takes an array of integer elements as a parameter, and two subarrays
of elements are recursively passed to the MergeSorter method. The resultant arrays
are joined and returned as the collection, as follows:*/
func MergeSorter(array []int) []int {

	if len(array) < 2 {
		return array
	}
	middle := len(array) / 2
	return JoinArrays(MergeSorter(array[:middle]), MergeSorter(array[middle:]))
}

// Join Arrays method
/*JoinArrays method
The JoinArrays function takes the leftArr and rightArr integer arrays as parameters.
The combined array is returned in the following code:*/
func JoinArrays(leftArr []int, rightArr []int) []int {

	num, i, j := len(leftArr)+len(rightArr), 0, 0
	array := make([]int, num)

	var k int
	for k = 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			array[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			array[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = rightArr[j]
			j++
		}
	}
	return array
}

// main method
func main() {
	/*The main method initializes the integer array of 40 elements, and the elements are printed
	before and after sorting, as follows:*/
	elements := createArray(40)
	fmt.Println("\n Before Sorting \n\n", elements)
	fmt.Println("\n-After Sorting\n\n", MergeSorter(elements))
}
