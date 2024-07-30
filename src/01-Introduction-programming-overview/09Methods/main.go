package main

import "fmt"

//structure
type Rect struct {
	width  float64
	height float64
}

//associating method with struct
func (r Rect /*here the r is the struct receiver*/) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func main() {
	r := Rect{width: 10, height: 10}
	fmt.Println("Area: ", r.Area())
	fmt.Println("Perimeter: ", r.Perimeter())
	ptr := &Rect{width: 10, height: 5}
	fmt.Println("Area: ", ptr.Area())
	fmt.Println("Perimeter: ", ptr.Perimeter())

	var data MyInt = 1
	fmt.Println("value before increment1() call :", data)
	data.increment1()
	fmt.Println("value after increment1() call :", data)
	data.increment2()
	fmt.Println("value after increment2() call :", data)

}

/*There are two ways of defining associated function of a data type.
1. Accessor function, which take receiver as value. Go passes a copy of the instance this function (Just like pass by
value.). Any change done over the object is not reflected in the calling object.
The syntax of accessor function:
func (r <Receiver Data type>) <Function Name>(<Parameter List>) (<Return List>)*/

/*2. Mutator function, which take receiver as pointer. Go passes a pointer to the instance this function (Just like pass by
reference). Any change done over the object is reflected in the calling object.
The syntax of mutator function:
func (r *<Receiver Data type>) <Function Name>(<Parameter List>) (<Return List>)*/
type MyInt int

func (data MyInt) increment1() {
	data = data + 1
}
func (data *MyInt) increment2() {
	*data = *data + 1
}

/*Analysis:
· Accessor function increment1() does changes on a local copy of the object. Therefore, changes done are lost.
· Modifier function increment2() does changes on the actual instance so changes done are preserved.*/
