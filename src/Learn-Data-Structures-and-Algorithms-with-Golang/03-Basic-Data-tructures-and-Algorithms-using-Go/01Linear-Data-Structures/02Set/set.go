// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

/*A Set is a linear data structure that has a collection of values that are not repeated. A set can
store unique values without any particular order. In the real world, sets can be used to
collect all tags for blog posts and conversation participants in a chat. The data can be of
Boolean, integer, float, characters, and other types. Static sets allow only query operations,
which means operations related to querying the elements. Dynamic and mutable sets allow
for the insertion and deletion of elements. Algebraic operations such as union, intersection,
difference, and subset can be defined on the sets. The following example shows
the Set integer with a map integer key and bool as a value:*/
//Set class
type Set struct {
	integerMap map[int]bool
}

// create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// adds the element to the set
/*The AddElement method adds the element to a set. In the following code snippet, the
AddElement method of the Set class adds the element to integerMap if the element is not
in the Set. The integerMap element has the key integer and value as bool, as shown in
the following code:*/
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

// deletes the element from the set
/*The DeleteElement method deletes the element from integerMap using the delete
method. This method removes the element from the integerMap of the set, as foll*/
func (set *Set) DeleteElement(element int) {

	delete(set.integerMap, element)
}

// checks if element is in the set
/*The ContainsElement method
The ContainsElement method of the Set class checks whether or not the element exists
in integerMap. The integerMap element is looked up with a key integer element, as
shown in the following code example:*/

func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]

	return exists
}

/*The InterSect method
In the following code, the InterSect method on the Set class returns an
intersectionSet that consists of the intersection of set and anotherSet. The set class
is traversed through integerMap and checked against another Set to see if any elements
exist:
*/

// Intersect method returns the set which intersects with anotherSet
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	var value int
	for value, _ = range set.integerMap {
		if anotherSet.ContainsElement(value) {

			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

/*The example output after invoking the intersect with the parameter of another Set is as
follows. A new intersectSet is created. The current set is iterated and every value is
checked to see if it is in another set. If the value is in another set, it is added to
the set intersect:*/

/*The Union method on the Set class returns a unionSet that consists of a union of set and
anotherSet. Both sets are traversed through integerMap keys, and union set is updated
with elements from the sets, as follows:*/

// Union method returns the set which is union of the set with anotherSet
func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	var value int
	for value, _ = range set.integerMap {
		unionSet.AddElement(value)
	}

	for value, _ = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}

	return unionSet
}

// main method
func main() {
	var set *Set
	set = &Set{}

	set.New()

	set.AddElement(1)
	set.AddElement(2)

	fmt.Println("initial set", set)
	fmt.Println(set.ContainsElement(1))

	var anotherSet *Set
	anotherSet = &Set{}

	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)

	fmt.Println("another set", set)

	fmt.Println("intersection of sets ", set.Intersect(anotherSet))

	fmt.Println("union of sets ", set.Union(anotherSet))

}
