// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

// To get started, a list can be used in Go, as shown in the following example; elements are
// added through the PushBack method on the list, which is in the container/list
// package:
// main method
func main() {
	var intList list.List // declaring var list
	intList.PushBack(11)  //adding elements into list
	intList.PushBack(23)
	intList.PushBack(34)
	intList.PushBack(45)
	//looping into list
	for element := intList.Front(); //step1
	element != nil;                 //step2
	element = element.Next() {
		fmt.Println(element.Value.(int))
	} //step3
}
