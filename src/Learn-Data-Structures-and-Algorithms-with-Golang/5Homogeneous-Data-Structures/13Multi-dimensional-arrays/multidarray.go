// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	"math/rand"
)

/*Multi-dimensional arrays
An array is a homogeneous collection of data elements. An array's indexes range from
index 0 to index m-1, where m is the fixed length of the array. An array with multiple
dimensions is an array of an array. The following code initializes a multi-dimensional
array. A three-dimensional array is printed:*/
//main method
func main() {

	var threedarray [2][2][2]int

	var i int

	var j int

	var k int

	for i = 0; i < 2; i++ {

		for j = 0; j < 2; j++ {

			for k = 0; k < 2; k++ {

				threedarray[i][j][k] = rand.Intn(3)
			}
		}
	}

	fmt.Println(threedarray)
}
