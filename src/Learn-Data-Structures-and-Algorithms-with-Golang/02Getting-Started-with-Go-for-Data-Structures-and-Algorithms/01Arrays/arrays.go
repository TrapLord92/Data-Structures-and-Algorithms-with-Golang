// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

/*Arrays are the most famous data structures in different programming languages. Different
data types can be handled as elements in arrays such as int, float32, double, and others.
The following code snippet shows the initialization of an array (arrays.go):*/
// main method
func main() {

	var arr = [5]int{1, 2, 4, 5, 6}
	/*An array's size can be found with the len() function. A for loop is used for accessing all
	the elements in an array, as follows:*/
	var i int
	for i = 0; i < len(arr); i++ {

		fmt.Println("printing elements ", arr[i])

	}
	/*The range
	keyword can be used to access the index and value for each element:*/
	var value int
	for i, value = range arr {

		fmt.Println("The index is : ", i, "and the  value is : ", value)

	}
	/*The _ blank identifier is used if the index is ignored. The following code shows how a _
	blank identifier can be used*/
	for _, value = range arr {

		fmt.Println("blank range", value)

	}

}
