// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

/*For dynamic allocation, we use slice of slices. In the following code, slice of slices is
explained as two-dimensional*/
// main method
func main1() {

	var rows int
	var cols int

	rows = 7
	cols = 9
	var twodslices = make([][]int, rows)

	var i int

	for i = range twodslices {

		twodslices[i] = make([]int, cols)
	}

	fmt.Println(twodslices)
}
