// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

/*The len() function gives the current length of slice, and the capacity of slice can be
obtained using the cap() function. The following code sample shows the basic slice
creation and appending a slice*/
// main method
func main() {

	var slice = []int{1, 3, 5, 6}

	slice = append(slice, 8)

	fmt.Println("Capacity", cap(slice))

	fmt.Println("Length", len(slice))
}
