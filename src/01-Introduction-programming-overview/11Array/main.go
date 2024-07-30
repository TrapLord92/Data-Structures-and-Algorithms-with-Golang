package main

import "fmt"

/*Array
An array is a collection of variables of the same data type*/

func main() {
	var arr [10]int
	fmt.Println(arr)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	count := len(arr)
	fmt.Println("Length of array", count)
}

/*Output:
[0 0 0 0 0 0 0 0 0 0]
[0 1 2 3 4 5 6 7 8 9]
Length of array 10
Analysis:
· We had declared array arr of size 10. In Go language, the array size is a part of array. The array name is whole
array and not pointer to first element like that in C/ C++.
· By default, all the elements of array are initialized to their default value. For our example, the default value of int
type is 0.
· We can read and set each element of array.
· We can get size of array by using built-in function len().*/
