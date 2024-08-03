// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and sort package
import (
	"fmt"
	"sort"
)

type Mass float64
type Miles float64

// Thing class
/*The Thing class
A Thing class is defined in the following code with name, mass, distance,
meltingpoint, and freezingpoint properties:*/
type Thing struct {
	name          string
	mass          Mass
	distance      Miles
	meltingpoint  int
	freezingpoint int
}

// ByFactor Type
/*The ByFactor function type
ByFactor is a type of less function. The following code snippet shows the ByFactor
type:*/
type ByFactor func(Thing1 *Thing, Thing2 *Thing) bool

// Sort method
/*The Sort method
The Sort method is a function with the byFactor parameter, as shown here:*/
func (byFactor ByFactor) Sort(Things []Thing) {
	var sortedThings *ThingSorter
	sortedThings = &ThingSorter{
		Things:   Things,
		byFactor: byFactor,
	}
	sort.Sort(sortedThings)
}

// ThingSorter class
/*Thing sorter class
The Thing sorter sorts the elements by their properties. The ThingSorter class has an
array of things and a byFactor method:*/
type ThingSorter struct {
	Things   []Thing
	byFactor func(Thing1 *Thing, Thing2 *Thing) bool
}

/*The len, swap, and less methods
The sort.Interface has the len, swap, and less methods, as shown in the following
code:*/
// Len method
func (ThingSorter *ThingSorter) Len() int {
	return len(ThingSorter.Things)
}

// Swap method
func (ThingSorter *ThingSorter) Swap(i int, j int) {
	ThingSorter.Things[i], ThingSorter.Things[j] = ThingSorter.Things[j], ThingSorter.Things[i]
}

// Less method
func (ThingSorter *ThingSorter) Less(i int, j int) bool {
	return ThingSorter.byFactor(&ThingSorter.Things[i], &ThingSorter.Things[j])
}

// main method.
/*The main method
The main method creates things and initializes them with values. This method shows
things that are sorted by mass, distance, and name in decreasing order of distance:*/
func main() {
	var Things = []Thing{
		{"IronRod", 0.055, 0.4, 3000, -180},
		{"SteelChair", 0.815, 0.7, 4000, -209},
		{"CopperBowl", 1.0, 1.0, 60, -30},
		{"BrassPot", 0.107, 1.5, 10000, -456},
	}

	var name func(*Thing, *Thing) bool
	name = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.name < Thing2.name
	}
	var mass func(*Thing, *Thing) bool
	mass = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.mass < Thing2.mass
	}
	var distance func(*Thing, *Thing) bool
	distance = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.distance < Thing2.distance
	}
	var decreasingDistance func(*Thing, *Thing) bool
	decreasingDistance = func(p1, p2 *Thing) bool {
		return distance(p2, p1)
	}

	ByFactor(name).Sort(Things)
	fmt.Println("By name:", Things)

	ByFactor(mass).Sort(Things)
	fmt.Println("By mass:", Things)

	ByFactor(distance).Sort(Things)
	fmt.Println("By distance:", Things)

	ByFactor(decreasingDistance).Sort(Things)
	fmt.Println("By decreasing distance:", Things)

}
