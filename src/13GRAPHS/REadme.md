TODO

Introduction
In this chapter, we will study about Graphs. Graphs can be used to represent many interesting things in the real world.
Flights from cities to cities, rods connecting various town and cities. Even the sequence of steps that we take to become
ready for jobs daily, or even a sequence of classes that we take to become a graduate in computer science. Once we
have a good representation of the map, then we use a standard graph algorithms to solve many interesting problems of
real life.
The flight connection between major cities of India can also be represented by the below graph. Each node is a city and
each edge is a straight flight path from one city to another. You may want to go from Delhi to Chennai, if given this
data in good representation to a computer, through graph algorithms the computer may propose shortest, quickest or
cheapest path from soured to destination.
Google map that we use is also a big graph of lots of nodes and edges. That suggest shortest and quickest path to the
user.
Graph Definitions
A Graph is represented by G where G = (V, E), where V is a finite set of points called Vertices and E is a finite set of
Edges.
Each edge is a tuple (u, v) where u, v ∈ V. There can be a third component weight to the tuple. Weight is cost to go
from one vertex to another.
Edge in a graph can be directed or undirected. If the edges of graph are one way, it is called Directed graph or
Digraph. The graph whose edges are two ways are called Undirected graph or just graph.
A Path is a sequence of edges between two vertices. The length of a path is defined as the sum of the weight of all the
edges in the path.
Two vertices u and v are adjacent if there is an edge whose endpoints are u and v.
In the below graph:
V = {V1, V2, V3, V4, V5, V6, V7, V8, V9},
E =
The in-degree of a vertex v, denoted by indeg(v) is the number of incoming edges to the vertex v. The out-degree of a
vertex v, denoted by outdeg(v) is the number of outgoing edges of a vertex v. The degree of a vertex v, denoted by
deg(v) is the total number of edges whose one endpoint is v.
deg(v) = Indeg (v) + outdeg (v)
In the above graph
deg(V4) =3, indeg(V4) =2 and outdeg(V4) =1
A Cycle is a path that starts and ends at the same vertex and include at least one vertex.
An edge is a Self-Loop if two if its two endpoints coincide. This is a form of a cycle.
A vertex v is Reachable from vertex u or “u reaches v” if there is a path from u to v. In an undirected graph if v is
reachable from u then u is reachable from v. However, in a directed graph it is possible that u reaches v but there is no
path from v to u.
A graph is Connected if for any two vertices there is a path between them.
A Forest is a graph without cycles.
A Sub-Graph of a graph G is a graph whose vertices and edges are a subset of the vertices and edges of G.
A Spanning Sub-Graph of G is a graph that connects all the vertices of G.
A tree is an acyclic connected graph.
A Spanning tree of a graph is a tree that connects all the vertices of the graph. Since a Spanning-Tree is a tree, so it
should not have any cycle.

Graph Representation
In this section, we introduce the data structure for representing a graph. In the below representations we maintain a
collection to store edges and vertices of the graph.
Adjacency Matrix
One of the ways to represent a graph is to use two-dimensional matrix. Each combination of row and column represent
a vertex in the graph. The value stored at the location row v and column w is the edge from vertex v to vertex w. The
nodes that are connected by an edge are called adjacent nodes. This matrix is used to store adjacent relation so it is
called the Adjacency Matrix. In the below diagram, we have a graph and its Adjacency matrix.
In the above graph, each node has weight 1 so the adjacency matrix has just 1s or 0s. If the edges are of different,
weights that that weight will be filled in the matrix.
Pros: Adjacency matrix implementation is simple. Adding/Removing an edge between two vertices is just O(1). Query
if there is an edge between two vertices is also O(1)
Cons: It always consumes O(V2) space, which is an inefficient way to store when a graph is a sparse.
Sparse Matrix: In a huge graph, each node is connected with fewer nodes. So most of the places in adjacency matrix are
empty. Such matrix is called sparse matrix. In most of the real world problems adjacency matrix is not a good choice
for sore graph data.
Adjacency List
A more space efficient way of storing graph is adjacency list. In adjacency list of references to a linked list node. Each
reference corresponds to vertices in a graph. Each reference will then point to the vertices that are connected to it and
store this as a list.
In the below diagram node 2 is connected to 1, 3 and 4. Therefore, the reference at location 2 is pointing to a list that
contain 1, 3 and 4.
The adjacency list helps us to represent a sparse graph. An adjacency list representation also allows us to find all the
vertices that are directly connected to any vertices by just one link list scan. In all our programs, we are going to use the
adjacency list to store the graph.
Example 13.1: adjacency list representation of an undirected graph
type Edge struct {
source int
destination int
cost int
next *Edge
}
type EdgesList struct {
head *Edge
}
type Graph struct {
count int
VertexList [](*EdgesList)
}
func (g *Graph) Init(cnt int) {
g.count = cnt
g.VertexList = make([]*EdgesList, cnt)
for i := 0; i < cnt; i++ {
g.VertexList[i] = new(EdgesList)
g.VertexList[i].head = nil
}
}
func (g *Graph) AddEdge(source int, destination int, cost int) {
edge := &Edge{source, destination, cost, g.VertexList[source].head}
g.VertexList[source].head = edge
}
func (g *Graph) AddEdgeUnweighted(source int, destination int) {
g.AddEdge(source, destination, 1)
}
func (g *Graph) AddBiEdge(source int, destination int, cost int) {
g.AddEdge(source, destination, cost)
g.AddEdge(destination, source, cost)
}
func (g *Graph) AddBiEdgeUnweghted(source int, destination int) {
g.AddBiEdge(source, destination, 1)
}
func (g *Graph) Print() {
for i := 0; i < g.count; i++ {
ad := g.VertexList[i].head
if ad != nil {
fmt.Print("Vertex ", i, " is connected to : ")
for ad != nil {
fmt.Print("[", ad.destination, ad.cost, "] ")
ad = ad.next
}
fmt.Println()
}
}
}
Graph traversals
The Depth first search (DFS) and Breadth first search (BFS) are the two algorithms used to traverse a graph. These
same algorithms can also be used to find some node in the graph, find if a node is reachable etc.
Traversal is the process of exploring a graph by examining all its edges and vertices.
A list of some of the problems that are solved using graph traversal are:
1. Determining a path from vertex u to vertex v, or report an error if there is no such path.
2. Given a starting vertex s, finding the minimum number of edges from vertex s to all the other vertices of the
graph.
3. Testing of a graph G is connected.
4. Finding a spanning tree of a Graph.
5. Finding if there is some cycle in the graph.
Depth First Traversal
The DFS algorithm we start from starting point and go into depth of graph until we reach a dead end and then move up
to parent node (Backtrack). In DFS, we use stack to get the next vertex to start a search. Alternatively, we can use
recursion (system stack) to do the same.
Algorithm steps for DFS
1. Push the starting node in the stack.
2. Loop until the stack is empty.
3. Pop the node from the stack inside loop call this node current.
4. Process the current node. //Print, etc.
5. Traverse all the child nodes of the current node and push them into stack.
6. Repeat steps 3 to 5 until the stack is empty.
Stack based implementation of DFS
Example 13.2:
func (g *Graph) DFSStack() {
count := g.count
visited := make([]int, count)
var curr int
stk := new(stack.Stack)
for i := 0; i < count; i++ {
visited[i] = 0
}
visited[1] = 1
stk.Push(1)
for stk.Len() != 0 {
curr = stk.Pop().(int)
head := g.VertexList[curr].head
for head != nil {
if visited[head.destination] == 0 {
visited[head.destination] = 1
stk.Push(head.destination)
}
head = head.next
}
}
}
Recursion based implementation of DFS
Example 13.3:
func (g *Graph) DFS() {
count := g.count
visited := make([]int, count)
for i := 0; i < count; i++ {
visited[i] = 0
}
for i := 0; i < count; i++ {
if visited[i] == 0 {
visited[i] = 1
g.DFSRec(i, visited)
}
}
}
func (g *Graph) DFSRec(index int, visited []int) {
head := g.VertexList[index].head
fmt.Println(index)
for head != nil {
if visited[head.destination] == 0 {
//fmt.Println(index)
visited[head.destination] = 1
g.DFSRec(head.destination, visited)
}
head = head.next
}
}
Breadth First Traversal
In BFS algorithm, a graph is traversed in layer-by-layer fashion. The graph is traversed closer to the starting point. The
queue is used to implement BFS.
Algorithm steps for BFS
1. Push the starting node into the Queue.
2. Loop until the Queue is empty.
3. Remove a node from the Queue inside loop, and call this node current.
4. Process the current node.//print etc.
5. Traverse all the child nodes of the current node and push them into Queue.
6. Repeat steps 3 to 5 until Queue is empty.
Example 13.4:
func (g *Graph) BFS() {
count := g.count
visited := make([]int, count)
for i := 0; i < count; i++ {
visited[i] = 0
}
for i := 0; i < count; i++ {
if visited[i] == 0 {
g.BFSQueue(i, visited)
}
}
}
func (g *Graph) BFSQueue(index int, visited []int) {
var curr int
que := new(queue.Queue)
visited[index] = 1
que.Enqueue(index)
for que.Len() != 0 {
curr = que.Dequeue().(int)
head := g.VertexList[curr].head
for head != nil {
if visited[head.destination] == 0 {
visited[head.destination] = 1
que.Enqueue(head.destination)
}
head = head.next
}
}
}
A runtime analysis of DFS and BFS traversal is O(n+m) time, where n is the number of edges reachable from source
node and m is the number of edges incident on s.
The following problems have O(m+n) time performance:
1. Determining a path from vertex u to vertex v, or report an error if there is no such path.
2. Given a starting vertex s, finding the minimum number of edges from vertex s to all the other vertices of the graph.
3. Testing of a graph G is connected.
4. Finding a spanning tree of a Graph.
5. Finding if there is some cycle in the graph.
Problems in Graph
Determining a path from vertex u to vertex v
IF there is a path from u to v and we are doing DFS from u then v must be visited. Moreover, if there is no path then
report an error.
Example 13.5:
func (g *Graph) PathExist(source int, destination int) bool {
count := g.count
visited := make([]int, count)
for i := 0; i < count; i++ {
visited[i] = 0
}
visited[source] = 1
g.DFSRec(source, visited)
return visited[destination] != 0
}
Given a starting vertex s, finding the minimum number of edges from vertex s to all the
other vertices of the graph
Look for single source shortest path algorithm for each edge cost as 1 unit.
Testing of a graph G is connected.
IF there is a path from u to v and we are doing DFS from u then v must be visited. Moreover, if there is no path then
report an error.
Example 13.6:
func (g *Graph) IsConnected() bool {
count := g.count
visited := make([]int, count)
for i := 0; i < count; i++ {
visited[i] = 0
}
visited[0] = 1
g.DFSRec(0, visited)
for i := 0; i < count; i++ {
if visited[i] == 0 {
return false
}
}
return true
}
Finding if there is some cycle in the graph.
Modify DFS problem and get this done.
Directed Acyclic Graph
A Directed Acyclic Graph (DAG) is a directed graph with no cycle. A DAG represent relationship, which is more
general than a tree. Below is an example of DAG, this is how someone becomes ready for work. There are N other real
life examples of DAG such as coerces selection to being graduated from college
Topological Sort
A topological sort is a method of ordering the nodes of a directed graph in which nodes represent activities and the
edges represent dependency among those tasks. For topological sorting to work it is required that the graph should be a
DAG which means it should not have any cycle. Just use DFS to get topological sorting.
Example 13.7:
func (g *Graph) TopologicalSort() {
fmt.Print("Topological order of given graph is : ")
var count = g.count
stk := new(stack.Stack)
visited := make([]int, count)
for i := 0; i < count; i++ {
visited[i] = 0
}
for i := 0; i < count; i++ {
if visited[i] == 0 {
visited[i] = 1
g.TopologicalSortDFS(i, visited, stk)
}
}
for stk.Len() != 0 {
fmt.Print(stk.Pop().(int), " ")
}
fmt.Println()
}
func (g *Graph) TopologicalSortDFS(index int, visited []int, stk *stack.Stack) {
head := g.VertexList[index].head
for head != nil {
if visited[head.destination] == 0 {
visited[head.destination] = 1
g.TopologicalSortDFS(head.destination, visited, stk)
}
head = head.next
}
stk.Push(index)
}
Topology sort is DFS traversal of topology graph. First, the children of node are added to the stack then only the
current node is added. So the sorting order is maintained. Reader is requested to run some examples to understand this
algo.
Minimum Spanning Trees (MST)
A Spanning Tree of a graph G is a tree that contains all the edges of the Graph G.
A Minimum Spanning Tree is a spanning-tree whose sum of length/weight of edges is minimum as possible.
For example, if you want to setup communication between a set of cities, then you may want to use the least amount of
wire as possible. MST can be used to find the network path and wire cost estimate.
Prim’s Algorithm for MST
Prim’s algorithm grows a single tree T, one edge at a time, until it becomes a spanning tree.
We initialize T with zero edges and U with single node. Where T is spanning tree edges set and U is spanning tree
vertex set.
At each step, Prim’s algorithm adds the smallest value edge with one endpoint in U and other not in us. Since each edge
adds one new vertex to U, after n − 1 additions, U contain all the vertices of the spanning tree and T becomes a
spanning tree.
Example 13.8:
// Returns the MST by Prim’s Algorithm
// Input: A weighted connected graph G = (V, E)
// Output: Set of edges comprising a MST
Algorithm Prim(G)
T = {}
Let r be any vertex in G
U = {r}
for i = 1 to |V| - 1 do
e = minimum-weight edge (u, v)
With u in U and v in V-U
U = U + {v}
T = T + {e}
return T
Prim’s Algorithm using a priority queue (min heap) to get the closest fringe vertex
Time Complexity will be O(m log n) where n vertices and m edges of the MST.
Example 13.9:
func (g *Graph) Prims() {
count := g.count
previous := make([]int, count)
dist := make([]int, count)
que := new(PQueue)
que.Init()
source := 1
for i := 0; i < count; i++ {
previous[i] = -1
dist[i] = Infinite
}
dist[source] = 0
edge := &Edge{source, source, 0, nil}
que.Push(edge, edge.cost)
for que.Len() != 0 {
edge = que.Pop().(*Edge) // Pop from PQ
if dist[edge.destination] < edge.cost {
continue
}
dist[edge.destination] = edge.cost
previous[edge.destination] = edge.source
adl := g.VertexList[edge.destination]
adn := adl.head
for adn != nil {
if previous[adn.destination] == -1 {
edge = &Edge{adn.source, adn.destination, adn.cost, nil}
que.Push(edge, edge.cost)
}
adn = adn.next
}
}
// Printing result.
for i := 0; i < count; i++ {
if dist[i] == Infinite {
fmt.Println(" edge id ", i, " prev ", previous[i], " distance : Unreachable")
} else {
fmt.Println(" edge id ", i, " prev ", previous[i], " distance : ", dist[i])
}
}
}
Kruskal’s Algorithm
Kruskal’s Algorithm repeatedly chooses the smallest-weight edge that does not form a cycle.
Sort the edges in non-decreasing order of cost: c (e1) ≤ c (e2) ≤ · · · ≤ c (em).
Set T to be the empty tree. Add edges to tree one by one, if it does not create a cycle.
Example 13.10:
// Returns the MST by Kruskal’s Algorithm
// Input: A weighted connected graph G = (V, E)
// Output: Set of edges comprising a MST
Algorithm Kruskal(G)
Sort the edges E by their weights
T = { }
while |T | + 1 < |V | do
e = next edge in E
if T + {e} does not have a cycle then
T = T + {e}
return T
Kruskal’s Algorithm is O(E log V) using efficient cycle detection.
Shortest Path Algorithms in Graph
Single Source Shortest Path
For a graph G= (V, E), the single source shortest path problem is to find the shortest path from a given source vertex s
to all the vertices of V.
Single Source Shortest Path for unweighted Graph.
Find single source shortest path for unweighted graph or a graph with all the vertices of same weight.
Example 13.11:
func (g *Graph) ShortestPath(source int) {
var curr int
count := g.count
distance := make([]int, count)
path := make([]int, count)
que := new(queue.Queue)
for i := 0; i < count; i++ {
distance[i] = -1
}
que.Enqueue(source)
distance[source] = 0
path[source] = source
for que.Len() != 0 {
curr = que.Dequeue().(int)
head := g.VertexList[curr].head
for head != nil {
if distance[head.destination] == -1 {
distance[head.destination] = distance[curr] + 1
path[head.destination] = curr
que.Enqueue(head.destination)
}
head = head.next
}
}
for i := 0; i < count; i++ {
fmt.Println(path[i], " to ", i, " weight ", distance[i])
}
}
Dijkstra’s algorithm
Dijkstra’s algorithm for single-source shortest path problem for weighted edges with no negative weight. Given a
weighted connected graph G, find shortest paths from the source vertex s to each of the other vertices. Dijkstra’s
algorithm is similar to prims algorithm. It maintains a set of nodes for which shortest path is known.
The algorithm starts by keeping track of the distance of each node and its parents. All the distance is set to infinite in
the beginning as we do not know the actual path to the nodes and parents of all the vertices are set to null. All the
vertices are added to a priority queue (min heap implementation)
At each step algorithm takes one vertex from the priority queue (which will be the source vertex in the beginning).
Then update the distance list corresponding to all the adjacent vertices. When the queue is empty, then we will have the
distance and parent list fully populated.
Example 13.12:
// Solves SSSP by Dijkstra’s Algorithm
// Input: A weighted connected graph G = (V, E)
// with no negative weights, and source vertex v
// Output: The length and path from s to every v
Algorithm Dijkstra(G, s)
for each v in V do
D[v] = infinite // Unknown distance
P[v] = null //unknown previous node
add v to PQ //adding all nodes to priority queue
D[source] = 0 // Distance from source to source
while (PQ is not empty)
u = vertex from PQ with smallest D[u]
remove u from PQ
for each v adjacent from u do
alt = D[u] + length ( u , v)
if alt < D[v] then
D[v] = alt
P[v] = u
Return D[] , P[]
Time Complexity will be O(|E|log|V|)
Note: Dijkstra’s algorithm does not work for graphs with negative edges weight.
Note: Dijkstra’s algorithm is applicable to both undirected and directed graphs.
Example 13.13:
func (g *Graph) Dijkstra(source int) {
count := g.count
previous := make([]int, count)
dist := make([]int, count)
que := new(PQueue)
que.Init()
for i := 0; i < count; i++ {
previous[i] = -1
dist[i] = Infinite
}
dist[source] = 0
edge := &Edge{source, source, 0, nil}
que.Push(edge, edge.cost)
for que.Len() != 0 {
edge = que.Pop().(*Edge) // Pop from PQ
if dist[edge.destination] < edge.cost {
continue
}
dist[edge.destination] = edge.cost
previous[edge.destination] = edge.source
adl := g.VertexList[edge.destination]
adn := adl.head
for adn != nil {
if previous[adn.destination] == -1 {
edge = &Edge{adn.source, adn.destination, adn.cost + dist[adn.source], nil}
que.Push(edge, edge.cost)
}
adn = adn.next
}
}
// Printing result.
for i := 0; i < count; i++ {
if dist[i] == Infinite {
fmt.Println(" edge id ", i, " prev ", previous[i], " distance : Unreachable")
} else {
fmt.Println(" edge id ", i, " prev ", previous[i], " distance : ", dist[i])
}
}
}
Bellman Ford Shortest Path
The bellman ford algorithm works even when there are negative weight edges in the graph. It does not work if there is
some cycle in the graph whose total weight is negative.
Example 13.14:
func (g *Graph) BellmanFordShortestPath(source int) {
count := g.count
distance := make([]int, count)
path := make([]int, count)
for i := 0; i < count; i++ {
distance[i] = Infinite
}
distance[source] = 0
path[source] = source
for i := 0; i < count-1; i++ {
for j := 0; j < count; j++ {
head := g.VertexList[j].head
for head != nil {
newDistance := distance[j] + head.cost
if distance[head.destination] > newDistance {
distance[head.destination] = newDistance
path[head.destination] = j
}
head = head.next
}
}
}
for i := 0; i < count; i++ {
fmt.Println(path[i], " to ", i, " weight ", distance[i])
}
}
All Pairs Shortest Paths
Given a weighted graph G(V, E), the all pair shortest path problem is to find the shortest path between all pairs of
vertices u, v є V. Execute n instances of single source shortest path algorithm for each vertex of the graph.
The complexity of this algorithm will be O(n3)
Exercise
1. In the various path-finding algorithms, we have created a path array that just store immediate parent of a node, print
the complete path for it.
2. All the functions are implemented considering as if the graph is represented by adjacency list. To write all those
functions for graph representation as adjacency matrix.
3. Given a start string, end string and a set of strings, find if there exists a path between the start string and end string
via the set of strings.
A path exists if we can get from start string to end the string by changing (no addition/removal) only one character
at a time. The restriction is that the new string generated after changing one character has to be in the set.
Start: "cog"
End: "bad"
Set: ["bag", "cag", "cat", "fag", "con", "rat", "sat", "fog"]
One of the paths: "cog" -> "fog" -> "fag" -> "bag" -> "bad"