# Complexity analysis of algorithms
The complexity of an algorithm is measured by the speed of the algorithm. Typically, the
algorithm will perform differently based on processor speed, disk speed, memory, and
other hardware parameters. Hence, asymptotical complexity is used to measure the
complexity of an algorithm. An algorithm is a set of steps to be processed by different
operations to achieve a task. The time taken for an algorithm to complete is based on the
number of steps taken.
Let's say an algorithm iterates through an array, m, of size 10 and update the elements to
the sum of index and 200. The computational time will be 10*t, where t is the time taken to
add two integers and update them to an array. The next step will be printing them after
iterating over an array. The t time parameter will vary with the hardware of the computer
used. Asymptotically, the computational time grows as a factor of 10, as shown in the
following code:
//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main
// importing fmt package
import (
"fmt"
)
// main method
func main() {
var m [10]int
var k int
for k = 0; k < 10; k++ {
m[k] = k + 200
fmt.Printf("Element[%d] = %d\n", k, m[k] )
}
}
Run the following commands:
go run complexity.go
Data Structures and Algorithms Chapter 1
[ 38 ]
The following screenshot displays the output: