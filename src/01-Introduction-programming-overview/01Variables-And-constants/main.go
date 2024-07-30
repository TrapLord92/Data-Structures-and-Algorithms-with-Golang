package main

import "fmt"

func main() {

	//declaring variables
	var variable int
	var variable2 int
	variable = 100

	fmt.Println("Stored value : ", variable)
	fmt.Println("\n Stored value : ", variable2)

	//constant declaration

	const Pi = 3.14
	fmt.Println("\n Pi value : ", Pi)

	//short declaration
	variable3 := 100
	fmt.Println("\n Stored value (short declaration) : ", variable3)

}
