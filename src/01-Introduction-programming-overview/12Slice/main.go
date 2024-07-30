package main

import "fmt"

/*Go Array is a collection of variables of same data type. Arrays are fixed in length and does not have any inbuilt method
to increase their size.
Go Slice is an abstraction over Array. It actually uses arrays as an underlying structure. The various operations over
slice are:
1. Inbuilt append() function is used to add the elements to a slice. If the size of underlying array is not enough then
automatically a new array is created and content of the old array is copied to it.
2. The len() function returns the number of elements presents in the slice.
3. The cap() function returns the capacity of the underlying array of the slice.
4. The copy() function, the contents of a source slice are copied to a destination slice.
5. Re-slices the slice, the syntax <Slice Name>[start : end] , It returns a slice object containing the elements of base
slice from index start to end-1. Length of the new slice will be (end - start), and capacity will be cap (base slice) -
start.
Slice provides these utility functions because of which it is widely used in Go programming.
To define a slice, you can declare it as an array without specifying its size. Alternatively, you can use make function to
create a slice.*/

func main() {
	var s []int
	for i := 1; i <= 17; i++ {
		s = append(s, i)
		PrintSlice(s)
	}
	/*Analysis:
	· First, we had created an empty slice s.
	· Using append() function we are adding elements to the slice.
	· The capacity of the underlying array is increasing from 1 to 2. Then from 2 to 4. Then 8 to 16 and finally 32.
	· As we keep adding elements to the slice, its capacity is managed automatically.*/

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PrintSlice(s2) // [1 2 3 4 5 6 7 8 9 10] :: len=10 cap=10
	a := make([]int, 10)
	PrintSlice(a) // [0 0 0 0 0 0 0 0 0 0] :: len=10 cap=10
	b := make([]int, 0, 10)
	PrintSlice(b) // [] :: len=0 cap=10
	c := s[0:4]
	PrintSlice(c) // [1 2 3 4] :: len=4 cap=10
	d := c[2:5]
	PrintSlice(d) // [3 4 5] :: len=3 cap=8
	/*Analysis:
	· First, we had created a slice s using array like initialization.
	· Then we had created slice “a” with 10 elements (Capacity is also 10) which are all initialized to 0
	· Then we had created slice “b” with 0 elements with capacity of 10 elements.
	· Then we are using re-slicing to create another slice “c”, which contain elements of slice “s” from index 0 to 3.
	Capacity of the slice “c” remain 10.
	· Then we are using re-slicing to create another slice “d”, which contain elements of slice “c” from index 2 to 5.
	Capacity of the slice “d” reduced to 8.*/
}
func PrintSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))
}
