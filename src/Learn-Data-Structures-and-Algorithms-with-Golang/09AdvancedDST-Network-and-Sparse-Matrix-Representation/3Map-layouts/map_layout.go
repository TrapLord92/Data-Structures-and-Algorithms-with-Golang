// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*Map layouts
A map layout is a geographical visualization of locations that are linked together. The
nodes in the graph of a map consist of geo-based information. The node will have
information such as the name of the location, latitude, and longitude. Maps are laid out in
different scales. Cartographic design is referred to as map creation using geographic
information.*/

/*
A map layout is shown in the following code snippet. The Place class consists of Name,
Latitude, and Longitude properties:*/
// Place class
type Place struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// MapLayout class
/*The MapLayout class
The MapLayout class consists of GraphNodes and Links:*/
type MapLayout struct {
	GraphNodes map[Place]struct{}
	Links      map[Place]map[Place]struct{}
}

// NewMapLayout method
/*The NewMapLayout method
The NewMapLayout method creates a MapLayout. The MapLayout has GraphNodes and
links maps:*/
func NewMapLayout() *MapLayout {
	return &MapLayout{
		GraphNodes: make(map[Place]struct{}),
		Links:      make(map[Place]map[Place]struct{}),
	}
}

// AddPlace method
/*The AddPlace method
The AddPlace method of the MapLayout class takes place as a parameter and returns true
if the place exists. If the place does not exist, then a graph node with a new place key is
created:*/
func (mapLayout *MapLayout) AddPlace(place Place) bool {

	var exists bool
	if _, exists = mapLayout.GraphNodes[place]; exists {
		return true
	}
	mapLayout.GraphNodes[place] = struct{}{}
	return true
}

// Add Link
/*The AddLink method
The AddLink method of the MapLayout class takes the places as parameters and links them
together:*/
func (mapLayout *MapLayout) AddLink(place1 Place, place2 Place) {
	var exists bool
	if _, exists = mapLayout.GraphNodes[place1]; !exists {
		mapLayout.AddPlace(place1)
	}
	if _, exists = mapLayout.GraphNodes[place2]; !exists {
		mapLayout.AddPlace(place2)
	}

	if _, exists = mapLayout.Links[place1]; !exists {
		mapLayout.Links[place1] = make(map[Place]struct{})
	}
	mapLayout.Links[place1][place2] = struct{}{}

}

// PrintLinks method
/*The PrintLinks method
The PrintLinks method of MapLayout prints the places and the links:*/
func (mapLayout *MapLayout) PrintLinks() {
	var root Place
	root = Place{"Algeria", 3, 28}

	fmt.Printf("Printing all links adjacent to %s\n", root.Name)

	var node Place
	for node = range mapLayout.Links[root] {
		fmt.Printf("Link: %s -> %s\n", root.Name, node.Name)
	}

	var m map[Place]struct{}
	fmt.Println("Printing all links.")
	for root, m = range mapLayout.Links {
		var vertex Place
		for vertex = range m {
			fmt.Printf("Link: %s -> %s\n", root.Name, vertex.Name)
		}
	}
}

/*The main method
In the main method, the map layout is created by invoking the NewMapLayout method.
Places are instantiated and added to the map layout. Then, the links are added between
places:*/
// main method
func main() {

	var mapLayout *MapLayout

	mapLayout = NewMapLayout()

	var root Place = Place{"Algeria", 3, 28}
	var netherlands Place = Place{"Netherlands", 5.75, 52.5}

	var korea Place = Place{"Korea", 124.1, -8.36}
	var tunisia Place = Place{"Tunisia", 9, 34}

	mapLayout.AddPlace(root)
	mapLayout.AddPlace(netherlands)
	mapLayout.AddPlace(korea)
	mapLayout.AddPlace(tunisia)

	mapLayout.AddLink(root, netherlands)
	mapLayout.AddLink(root, korea)
	mapLayout.AddLink(root, tunisia)

	var singapore Place = Place{"Singapore", 103.8, 1.36}
	var uae Place = Place{"UAE", 54, 24}
	var japan Place = Place{"Japan", 139.75, 35.68}

	mapLayout.AddLink(korea, singapore)
	mapLayout.AddLink(korea, japan)
	mapLayout.AddLink(netherlands, uae)

	mapLayout.PrintLinks()
}
