// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*Network representation using graphs
A graph is a representation of a set of objects that's connected by links. The links connect
vertices, which are points. The basic operations on a graph are the addition and removal of
links and vertices. These are some different types of graphs:
Directed graph
Non-directed graph
Connected graph
Non-connected graph
Simple graph
Multi-graph
An adjacency list consists of adjacent vertices of a graph that have objects or records. An
adjacency matrix consists of source and destination vertices. An incidence matrix is a twodimensional
Boolean matrix. The matrix has rows of vertices and columns that represent
the links (edges).
Network representation using a graph is shown in the following code.*/

/* A social graph
consists of an array of links:*/
// Social Graph
type SocialGraph struct {
	Size  int
	Links [][]Link
}

// Link class
/*The Link class
The Link class consists of the vertex1 and vertex2 vertices and the LinkWeight integer
property:*/
type Link struct {
	Vertex1    int
	Vertex2    int
	LinkWeight int
}

// NewSocialGraph method
/*The NewSocialGraph method
The NewSocialGraph function creates a social graph given num, which is the size of the
graph. Size is the number of links in the graph:*/
func NewSocialGraph(num int) *SocialGraph {
	return &SocialGraph{
		Size:  num,
		Links: make([][]Link, num),
	}
}

// AddLink method
/*The AddLink method
The AddLink method adds the link between two vertices. The AddLink method of a social
graph takes vertex1, vertex2, and weight as parameters. The method adds the link from
vertex1 to vertex2, as shown in the following code:*/
func (socialGraph *SocialGraph) AddLink(vertex1 int, vertex2 int, weight int) {
	socialGraph.Links[vertex1] = append(socialGraph.Links[vertex1], Link{Vertex1: vertex1, Vertex2: vertex2, LinkWeight: weight})
}

// Print Links Example
/*The PrintLinks method of the SocialGraph class prints the links from vertex = 0 and
all the links in the graph:*/
func (socialGraph *SocialGraph) PrintLinks() {

	var vertex int
	vertex = 0

	fmt.Printf("Printing all links from %d\n", vertex)
	var link Link
	for _, link = range socialGraph.Links[vertex] {
		fmt.Printf("Link: %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LinkWeight)
	}

	fmt.Println("Printing all links in graph.")
	var adjacent []Link
	for _, adjacent = range socialGraph.Links {
		for _, link = range adjacent {
			fmt.Printf("Link: %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LinkWeight)
		}
	}
}

// main method
/*The main method
The main method creates a social graph by invoking the NewSocialGraph method. The
links from 0 to 1, 0 to 2, 1 to 3, and 2 to 4 are added to the social graph. The links are
printed using the printLinks method:*/
func main() {

	var socialGraph *SocialGraph

	socialGraph = NewSocialGraph(4)

	socialGraph.AddLink(0, 1, 1)
	socialGraph.AddLink(0, 2, 1)
	socialGraph.AddLink(1, 3, 1)
	socialGraph.AddLink(2, 4, 1)

	socialGraph.PrintLinks()

}
