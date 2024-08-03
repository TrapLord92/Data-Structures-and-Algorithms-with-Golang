// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	"math/rand"
)

/*Tensors
A tensor is a multi-dimensional array of components that are spatial coordinates. Tensors
are used extensively in physics and biological studies in topics such as electromagnetism
and diffusion tensor imaging. William Rowan Hamilton was the first to come up with the
term tensor. Tensors play a basic role in abstract algebra and algebraic topology.
The tensor order is the sum of the order of its arguments, plus the order of the result tensor.
For example, an inertia matrix is a second-order tensor. Spinors are also multi-dimensional
arrays, but the values of their elements change via coordinate transformations.
The initialization of a tensor is shown here. The array is initialized with integer values
ranging from 0 to 3:*/

// main method
func main() {

	var array [3][3][3]int
	var i int
	var j int
	var k int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			for k = 0; k < 3; k++ {

				array[i][j][k] = rand.Intn(3)
			}
		}
	}
	fmt.Println(array)
	/*Unfolding a tensor is done along the first dimension. Rearranging the tensor mode's n
	vectors is referred to as mode n-unfolding of a tensor. 0-mode unfolding of a tensor array is
	shown here:*/

	fmt.Println("zero mode unfold")
	for j = 0; j < 3; j++ {
		for k = 0; k < 3; k++ {
			fmt.Printf("%d ", array[0][j][k])
		}
		fmt.Printf("\n")
	}
	/*1-mode unfolding of a tensor array is shown here. The array's first dimension index is set to
	1:*/
	fmt.Println("1-mode unfold")
	for j = 0; j < 3; j++ {
		for k = 0; k < 3; k++ {
			fmt.Printf("%d ", array[1][j][k])
		}
		fmt.Printf("\n")
	}
	/*The 2-mode unfolding of a tensor array is shown here. The array's first dimension row
	index is set to 2:*/
	fmt.Println("2-mode unfold")
	for j = 0; j < 3; j++ {
		for k = 0; k < 3; k++ {
			fmt.Printf("%d ", array[2][j][k])
		}
		fmt.Printf("\n")
	}

}
