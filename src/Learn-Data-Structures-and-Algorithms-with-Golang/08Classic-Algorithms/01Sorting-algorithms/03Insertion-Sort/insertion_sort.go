// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
	"math/rand"
	"time"
)

/*Insertion
Insertion sort is an algorithm that creates a final sorted array one element at a time. The
algorithm's performance is of the order O(n2). This algorithm is less efficient on large
collections than other algorithms, such as quick, heap, and merge sort. In real life, a good
example of insertion sort is the way cards are manually sorted by the players in a game of
bridge.
The implementation of the insertion sort algorithm is shown in the following code snippet.*/

/*The RandomSequence function takes the number of elements as a parameter and returns an
array of random integers:*/
// randomSequence method
func randomSequence(num int) []int {

	sequence := make([]int, num)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := 0; i < num; i++ {
		sequence[i] = r.Intn(999) - 999
	}
	return sequence
}

// InsertionSorter method
/*InsertionSorter method
The implementation of the InsertionSorter method is shown in the following snippet.
This method takes the array of integers as a parameter and sorts them:*/
func InsertionSorter(elements []int) {
	var n = len(elements)
	var i int

	for i = 1; i < n; i++ {
		var j int
		j = i
		for j > 0 {
			if elements[j-1] > elements[j] {
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j = j - 1
		}
	}
}

// main method
/*The main method
The main method initializes the sequence by invoking the randomSequence function, as
shown in the following code. The InsertionSorter function takes the sequence and
sorts it in ascending order:*/
func main() {

	sequence := randomSequence(24)
	fmt.Println("\n^^^^^^ Before Sorting ^^^ \n\n", sequence)
	InsertionSorter(sequence)
	fmt.Println("\n--- After Sorting ---\n\n", sequence, "\n")
}
