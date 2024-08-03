// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*[ 223 ]
}
Run the following command to execute the merge_sort.go file:
go run merge_sort.go
The output is as follows:
Let's take a look at the quick sort algorithm in the following section.
Quick
Quick sort is an algorithm for sorting the elements of a collection in an organized
way. Parallelized quick sort is two to three times faster than merge sort and heap sort. The
algorithm's performance is of the order O(n log n). This algorithm is a space-optimized
version of the binary tree sort algorithm.
In the following code snippet, the quick sort algorithm is implemented. The */

//Quick Sorter method
/*QuickSorter
function takes an array of integer elements, upper int, and below int as parameters.
The function divides the array into parts, which are recursively divided and sorted:*/
func QuickSorter(elements []int, below int, upper int) {
	if below < upper {
		part := divideParts(elements, below, upper)
		QuickSorter(elements, below, part-1)
		QuickSorter(elements, part+1, upper)
	}
}

// divideParts method
/*The divideParts method
The divideParts method takes an array of integer elements, upper int, and below
int as parameters. The method sorts the elements in ascending order, as shown in the
following code:*/
func divideParts(elements []int, below int, upper int) int {

	center := elements[upper]
	var i int
	i = below
	var j int
	for j = below; j < upper; j++ {
		if elements[j] <= center {
			swap(&elements[i], &elements[j])
			i += 1
		}
	}
	swap(&elements[i], &elements[upper])
	return i
}

// swap method
/*The swap method
In the following code snippet, the swap method exchanges elements by interchanging the
values:*/
func swap(element1 *int, element2 *int) {
	var val int
	val = *element1
	*element1 = *element2
	*element2 = val
}

// main method
/*The main method
The main method asks the user to input the number of elements and the elements to be
read. The array is initialized and printed before and after sorting, as follows:*/
func main() {
	var num int

	fmt.Print("Enter Number of Elements: ")
	fmt.Scan(&num)

	var array = make([]int, num)

	var i int
	for i = 0; i < num; i++ {
		fmt.Print("array[", i, "]: ")
		fmt.Scan(&array[i])
	}

	fmt.Print("Elements: ", array, "\n")
	QuickSorter(array, 0, num-1)
	fmt.Print("Sorted Elements: ", array, "\n")
}
