package main

import "fmt"

func main() {
	data := 10
	ptr := &data //linking variable to memory adress data

	fmt.Println("Value stored at variable var is ", data)
	fmt.Println("Value stored at variable var is ", *ptr) //acessing memory adress

	fmt.Println("The address of variable var is ", &data) //printing the memory address
	fmt.Println("The address of variable var is ", ptr)   //printing the memory address
}

/*
Value stored at variable var is  10
Value stored at variable var is  10
The address of variable var is  0xc04200e210
The address of variable var is  0xc04200e210
*/
