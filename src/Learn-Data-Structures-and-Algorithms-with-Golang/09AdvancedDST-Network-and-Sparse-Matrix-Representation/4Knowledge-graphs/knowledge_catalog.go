// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*Knowledge graphs
A knowledge graph is a network representation of entities, items, and users as nodes. The
nodes interact with one another via links or edges. Knowledge graphs are widely used
because they are schema less. These data structures are used to represent knowledge in the
form of graphs, and the nodes have textual information. Knowledge graphs are created by
using item, entity, and user nodes and linking them with edges.
An ontology consists of a knowledge graph of information nodes. The reasoner derives
knowledge from knowledge graphs. A knowledge graph consists of classes, slots, and
facets, which are ontological terms. In the following code, a knowledge graph consisting of
a car's bill of materials is explained. The Class type consists of a name, which is a string:*/

// Class Type
type Class struct {
	Name string
}

// Knowledge Graph type
/*The KnowledgeGraph class
The KnowledgeGraph class consists of GraphNodes and links:*/
type KnowledgeGraph struct {
	GraphNodes map[Class]struct{}
	Links      map[Class]map[Class]struct{}
}

// NewKnowledgeGraph method
/*The NewKnowledgeGraph method
The NewKnowledgeGraph method creates a knowledge graph, which consists of
GraphNodes and Links maps:*/
func NewKnowledgeGraph() *KnowledgeGraph {
	return &KnowledgeGraph{
		GraphNodes: make(map[Class]struct{}),
		Links:      make(map[Class]map[Class]struct{}),
	}
}

// AddClass method
/*The AddClass method
The AddClass method of the KnowledgeGraph class takes class as a parameter and
returns true if the class exists. If the class does not exist, a GraphNode is created with
class as a key:*/
func (knowledgeGraph *KnowledgeGraph) AddClass(class Class) bool {

	var exists bool
	if _, exists = knowledgeGraph.GraphNodes[class]; exists {
		return true
	}
	knowledgeGraph.GraphNodes[class] = struct{}{}
	return true
}

// Add Link
/*The AddLink method
The AddLink method of the KnowledgeGraph class takes class1 and
class2 as parameters, and a link is created between these classes:*/
func (knowledgeGraph *KnowledgeGraph) AddLink(class1 Class, class2 Class) {
	var exists bool
	if _, exists = knowledgeGraph.GraphNodes[class1]; !exists {
		knowledgeGraph.AddClass(class1)
	}
	if _, exists = knowledgeGraph.GraphNodes[class2]; !exists {
		knowledgeGraph.AddClass(class2)
	}

	if _, exists = knowledgeGraph.Links[class1]; !exists {
		knowledgeGraph.Links[class1] = make(map[Class]struct{})
	}
	knowledgeGraph.Links[class1][class2] = struct{}{}

}

// Print Links method
/*The PrintLinks method
The PrintLinks method of the KnowledgeGraph class prints the links and nodes:*/
func (knowledgeGraph *KnowledgeGraph) PrintLinks() {
	var car Class
	car = Class{"Car"}

	fmt.Printf("Printing all links adjacent to %s\n", car.Name)

	var node Class
	for node = range knowledgeGraph.Links[car] {
		fmt.Printf("Link: %s -> %s\n", car.Name, node.Name)
	}

	var m map[Class]struct{}
	fmt.Println("Printing all links.")
	for car, m = range knowledgeGraph.Links {
		var vertex Class
		for vertex = range m {
			fmt.Printf("Link: %s -> %s\n", car.Name, vertex.Name)
		}
	}
}

// main method
/*The main method
The main method creates the knowledge graph, and the classes are instantiated. The links
between the classes are created and printed:*/
func main() {

	var knowledgeGraph *KnowledgeGraph

	knowledgeGraph = NewKnowledgeGraph()

	var car = Class{"Car"}
	var tyre = Class{"Tyre"}
	var door = Class{"Door"}
	var hood = Class{"Hood"}

	knowledgeGraph.AddClass(car)
	knowledgeGraph.AddClass(tyre)
	knowledgeGraph.AddClass(door)
	knowledgeGraph.AddClass(hood)

	knowledgeGraph.AddLink(car, tyre)
	knowledgeGraph.AddLink(car, door)
	knowledgeGraph.AddLink(car, hood)

	var tube = Class{"Tube"}
	var axle = Class{"Axle"}
	var handle = Class{"Handle"}
	var windowGlass = Class{"Window Glass"}

	knowledgeGraph.AddClass(tube)
	knowledgeGraph.AddClass(axle)
	knowledgeGraph.AddClass(handle)
	knowledgeGraph.AddClass(windowGlass)

	knowledgeGraph.AddLink(tyre, tube)
	knowledgeGraph.AddLink(tyre, axle)
	knowledgeGraph.AddLink(door, handle)
	knowledgeGraph.AddLink(door, windowGlass)

	knowledgeGraph.PrintLinks()
}
