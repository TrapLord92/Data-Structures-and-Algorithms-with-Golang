//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
/*Tuples are finite ordered sequences of objects. They can contain a mixture of other data
types and are used to group related data into a data structure. In a relational database, a
tuple is a row of a table. Tuples have a fixed size compared to lists, and are also faster. A
finite set of tuples in the relational database is referred to as a relation instance. A tuple can
be assigned in a single statement, which is useful for swapping values. Lists usually contain
values of the same data type, while tuples contain different data. For example, we can store
a name, age, and favorite color of a user in a tuple. Tuples were covered in Chapter 1, Data
Structures and Algorithms. The following sample shows a multi-valued expression from a
function's call*/
package main

// importing fmt package
import (
	"fmt"
)

// h function which returns the product of parameters x and y
func h(x int, y int) int {

	return x * y
}

// g function which returns x and y parameters after modification
func g(l int, m int) (x int, y int) {
	x = 2 * l
	y = 4 * m
	return
}

// main method
func main() {

	fmt.Println(h(g(1, 2)))
}
