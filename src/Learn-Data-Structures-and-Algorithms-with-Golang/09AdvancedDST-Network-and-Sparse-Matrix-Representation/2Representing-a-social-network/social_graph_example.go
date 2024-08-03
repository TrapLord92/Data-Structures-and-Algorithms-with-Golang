// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*Representing a social network
A social network consists of social links that contain social entities such as people, friends,
discussions, shares, beliefs, trust, and likes. This graph is used to represent the social
network.
Metrics related to the proximity of entities can be calculated based on the graph. Social
graphs consist of graph nodes and links, which are maps with a key name and multiple
keys names, respectively:*/
// Name type
type Name string

// Social Graph class
type SocialGraph struct {
	GraphNodes map[Name]struct{}
	Links      map[Name]map[Name]struct{}
}

// NewSocialGraph method
/*The NewSocialGraph method
The NewSocialGraph method returns a social graph consisting of nil-valued GraphNodes
and Links:*/
func NewSocialGraph() *SocialGraph {
	return &SocialGraph{
		GraphNodes: make(map[Name]struct{}),
		Links:      make(map[Name]map[Name]struct{}),
	}
}

// AddEntity method
/*The AddEntity method
The AddEntity method adds the entity to the social graph. The AddEntity method of
the SocialGraph class takes name as a parameter and returns true if it is added to the
social graph:*/
func (socialGraph *SocialGraph) AddEntity(name Name) bool {

	var exists bool
	if _, exists = socialGraph.GraphNodes[name]; exists {
		return true
	}
	socialGraph.GraphNodes[name] = struct{}{}
	return true
}

// Add Link
/*The AddLink method
The AddLink method of the SocialGraph class takes name1 and name2 as parameters.
This method creates the entities if the named entities do not exist and creates a link between
the entities:*/
func (socialGraph *SocialGraph) AddLink(name1 Name, name2 Name) {
	var exists bool
	if _, exists = socialGraph.GraphNodes[name1]; !exists {
		socialGraph.AddEntity(name1)
	}
	if _, exists = socialGraph.GraphNodes[name2]; !exists {
		socialGraph.AddEntity(name2)
	}

	if _, exists = socialGraph.Links[name1]; !exists {
		socialGraph.Links[name1] = make(map[Name]struct{})
	}
	socialGraph.Links[name1][name2] = struct{}{}

}

/*
The PrintLinks method
The PrintLinks method of the SocialGraph class prints the links adjacent to the
root and all the links, as shown in the following code snippet:
*/
func (socialGraph *SocialGraph) PrintLinks() {
	var root Name
	root = Name("Root")

	fmt.Printf("Printing all links adjacent to %d\n", root)

	var node Name
	for node = range socialGraph.Links[root] {
		fmt.Printf("Link: %d -> %d\n", root, node)
	}

	var m map[Name]struct{}
	fmt.Println("Printing all links.")
	for root, m = range socialGraph.Links {
		var vertex Name
		for vertex = range m {
			fmt.Printf("Link: %d -> %d\n", root, vertex)
		}
	}
}

// main method
/*The main method
The main method creates a social graph. The entities, such as john, per, and cynthia, are
created and linked with the root node. The friends, such as mayo, lorrie, and ellie, are
created and linked with john and per:*/
func main() {

	var socialGraph *SocialGraph

	socialGraph = NewSocialGraph()

	var root Name = Name("Root")
	var john Name = Name("John Smith")
	var per Name = Name("Per Jambeck")
	var cynthia Name = Name("Cynthia Gibas")

	socialGraph.AddEntity(root)
	socialGraph.AddEntity(john)
	socialGraph.AddEntity(per)
	socialGraph.AddEntity(cynthia)

	socialGraph.AddLink(root, john)
	socialGraph.AddLink(root, per)
	socialGraph.AddLink(root, cynthia)

	var mayo Name = Name("Mayo Smith")
	var lorrie Name = Name("Lorrie Jambeck")
	var ellie Name = Name("Ellie Vlocksen")

	socialGraph.AddLink(john, mayo)
	socialGraph.AddLink(john, lorrie)
	socialGraph.AddLink(per, ellie)

	socialGraph.PrintLinks()
}
