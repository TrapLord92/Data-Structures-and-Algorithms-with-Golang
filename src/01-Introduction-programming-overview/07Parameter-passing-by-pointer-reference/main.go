package main

import "fmt"

func incrementPassBy(ptr *int) {
	*ptr++
}

func main() {
	x := 5
	fmt.Println("Initial value of x:", x)

	incrementPassBy(&x)
	fmt.Println("Value of x after increment:", x)
}
