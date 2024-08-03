// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

/*Slices
Go Slice is an abstraction over Go Array. Multiple data elements of the same type are
allowed by Go arrays. The definition of variables that can hold several data elements of the
same type are allowed by Go Array, but it does not have any provision of inbuilt methods
to increase its size in Go. This shortcoming is taken care of by Slices. A Go slice can be
appended to elements after the capacity has reached its size. Slices are dynamic and can
double the current capacity in order to add more elements.*/

// twiceValue method given slice of int type
func twiceValue(slice []int) {

	var i int
	var value int

	for i, value = range slice {

		slice[i] = 2 * value

	}

}

// main method
func main() {

	var slice = []int{1, 3, 5, 6}
	twiceValue(slice)

	var i int

	for i = 0; i < len(slice); i++ {

		fmt.Println("new slice value", slice[i])
	}
}
