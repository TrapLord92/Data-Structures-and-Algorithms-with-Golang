package main

/*Introduction
Greedy algorithms are generally used to solve optimization problems. To find the solution that minimizes or maximizes
some value (cost/profit/count etc.).
In greedy algorithm, solution is constructed through a sequence of steps. At each step, choice is made which is locally
optimal. We always take the next data to be processed depending upon the dataset which we have already processed
and then choose the next optimum data to be processed.
Greedy algorithms does not always give optimum solution. For some problems, greedy algorithm gives an optimal
solution. For most, they do not, but can be useful for fast approximations.
Greedy is a strategy that works well on optimization problems with the following characteristics:
1. Greedy choice: A global optimum can be arrived at by selecting a local optimum.
2. Optimal substructure: An optimal solution to the problem is made from optimal solutions of sub problems.
Some examples of brute force algorithms are:
Optimal solutions:
· Minimal spanning tree:
o Prim’s algorithm,
o Kruskal’s algorithm
· Dijkstra’s algorithm for single-source shortest path
· Huffman trees for optimal encoding
· Scheduling problems
Approximate solutions:
· Greedy algorithm for the Knapsack problem
· Coin exchange problem
Problems on Greedy Algorithm
Coin exchange problem
How can a given amount of money N be made with the least number of coins of given denominations D= {d1… dn}?
The Indian coin system {5, 10, 20, 25, 50,100}
Suppose we want to give change of a certain amount of 40 paisa.
We can make a solution by repeatedly choosing a coin ≤ to the current amount, resulting in a new amount. In the
greedy algorithm, we always choose the largest coin value possible without exceeding the total amount.
For 40 paisa: {25, 10, and 5}
The optimal solution will be {20, 20}
The greedy algorithm did not give us optimal solution, but it gave a fair approximation.
Algorithm MAKE-CHANGE (N)
C = {5, 20, 25, 50, 100} // constant denominations.
S = {} // set that will hold the solution set.
Value = N
WHILE Value != 0
x = largest item in set C such that x < Value
IF no such item THEN
RETURN "No Solution"
S = S + x
Value = Value - x
RETURN S
Minimum Spanning Tree
A spanning tree of a connected graph is a tree containing all the vertices.
A minimum spanning tree of a weighted graph is a spanning tree with the smallest sum of the edge weights.
Prim’s Algorithm
Prim’s algorithm grows a single tree T, one edge at a time, until it becomes a spanning tree.
We initialize T with zero edges and U with single node. Where T is spanning tree edges set and U is spanning tree
vertex set.
At each step, Prim’s algorithm adds the smallest value edge with one endpoint in U and other not in us.
Since each edge adds one new vertex to U, after n − 1 additions, U contain all the vertices of the spanning tree and T
becomes a spanning tree.
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
Kruskal’s Algorithm
Kruskal’s Algorithm is used to create minimum spanning tree. Spanning tree is created by choosing smallest weight
edge that does not form a cycle. And repeating this process until all the edges from the original set is exhausted.
Sort the edges in non-decreasing order of cost: c (e1) ≤ c (e2) ≤ · · · ≤ c (em).
Set T to be the empty tree. Add edges to tree one by one if it does not create a cycle. (If the new edge form cycle then
ignore that edge.)
// Returns the MST by Kruskal’s Algorithm
// Input: A weighted connected graph G = (V, E)
// Output: Set of edges comprising a MST
Algorithm Kruskal(G)
Sort the edges E by their weights
T = {}
while |T | + 1 < |V | do
e = next edge in E
if T + {e} does not have a cycle then
T = T + {e}
return T
Kruskal’s Algorithm is O(E log V) using efficient cycle detection.
Dijkstra’s algorithm for single-source shortest path problem
Dijkstra’s algorithm for single-source shortest path problem for weighted edges with no negative weight. It determine
the length of the shortest path from the source to each of the other nodes of the graph. Given a weighted graph G, we
need to find shortest paths from the source vertex s to each of the other vertices.
The algorithm starts by keeping track of the distance of each node and its parents. All the distance is set to infinite in
the beginning as we do not know the actual path to the nodes and parents of all the vertices are set to null. All the
vertices are added to a priority queue (min heap implementation)
At each step algorithm takes one vertex from the priority queue (which will be the source vertex in the beginning).
Then update the distance list corresponding to all the adjacent vertices. When the queue is empty, then we will have the
distance and parent list fully populated.
// Solve SSSP by Dijkstra’s Algorithm
// Input: A weighted connected graph G = (V, E)
// with no negative weights, and source vertex v
// Output: The length and path from s to every v
Algorithm Dijkstra(G, s)
for each v in V do
D[v] = infinite // Unknown distance
P[v] = null // Unknown previous node
add v to PQ // Adding all nodes to priority queue
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
Time Complexity will be O(|E|log|V|).
Note: Dijkstra’s algorithm does not work for graphs with negative edges weight.
Note: Dijkstra’s algorithm is applicable to both undirected and directed graphs.
Huffman trees for optimal encoding
Encoding is an assignment of bit strings of alphabet characters.
There are two types of encoding:
· Fixed-length encoding (eg., ASCII)
· Variable-length encoding (eg., Huffman code)
Variable length encoding can only work on prefix free encoding. Which means that no code word is a prefix of another
code word.
Huffman codes are the best prefix free code. Any binary tree with edges labeled as 0 and 1 will produce a prefix free
code of characters assigned to its leaf nodes.
Huffman’s algorithm is used to construct a binary tree whose leaf value is assigned a code, which is optimal for the
compression of the whole text need to be processed. For example, the most frequently occurring words will get the
smallest code so that the final encoded text is compressed.
Initialize n one-node trees with words and the tree weights with their frequencies. Join the two binary tree with smallest
weight into one and the weight of the new formed tree as the sum of weight of the two small trees. Repeat the above
process N-1 times and when there is just one big tree left you are done.
Mark edges leading to left and right subtrees with 0’s and 1’s, respectively.
Word Frequency
Apple 30
Banana 25
Mango 21
Orange 14
Pineapple 10
Word Value Code
Apple 30 11
Banana 25 10
Mango 21 01
Orange 14 001
Pineapple 10 000
It is clear that more frequency words gets smaller Huffman’s code.
// Computes optimal prefix code.
// Input: List W of character probabilities
// Output: The Huffman tree.
Algorithm Huffman(C[0..n - 1], W[0..n - 1])
PQ = {} // priority queue
for i = 0 to n - 1 do
T.char = C[i]
T.weight = W[i]
add T to priority queue PQ
for i = 0 to n - 2 do
L = remove min from PQ
R = remove min from PQ
T = node with children L and R
T.weight = L.weight + R.weight
add T to priority queue PQ
return T
The Time Complexity is O(nlogn).
Activity Selection Problem
Suppose that activities require exclusive use of common resources, and you want to schedule as many activities as
possible.
Let S = {a1,..., an} be a set of n activities.
Each activity ai needs the resource during a time period starting at si and finishing before fi, i.e., during [si, fi).
The optimization problem is to select the non-overlapping largest set of activities from S.
We assume that activities S = {a1,..., an} are sorted in finish time f1 ≤ f2 ≤ ... fn-1 ≤ fn (this can be done in Θ(n lg n)).
Example: Consider these activities:
I 1 2 3 4 5 6 7 8 9 10 11
S[i] 1 3 0 5 3 5 6 8 8 2 11
F[i] 4 5 6 7 8 9 10 11 12 13 14
Here is a graphic representation:
We chose an activity that start first, and then look for the next activity that starts after it is finished. This could result in
{a4, a7, a8}, but this solution is not optimal.
An optimal solution is {a1, a3, a6, a8}. (It maximizes the objective function of a number of activities scheduled.)
Another one is {a2, a5, a7, a9}. (Optimal solutions are not necessarily unique.)
How do we find (one of) these optimal solutions? Let us consider it as a dynamic programming problem.
We are trying to optimize the number of activities. Let us be greedy!
· The more time left after running an activity, the subsequent activities we can fit in.
· If we choose the first activity to finish, the more time will be left.
· Since activities are sorted by finish time, we will always start with a1.
· Then we can solve the single sub problem of activity scheduling in this remaining time.
Algorithm ActivitySelection(S[], F[], N)
Sort S[] and F [] in increasing order of finishing time
A = {a1}
K = 1
For m = 2 to N do
If S[m] >= F[k]
A = A + {am}
K = m
Return A
Knapsack Problem
A thief enters a store and sees a number of items with their cost and weight mentioned. His Knapsack can hold a max
weight. What should he steal to maximize profit?
Fractional Knapsack problem
A thief can take a fraction of an item (they are divisible substances, like gold powder).
The fractional knapsack problem has a greedy solution one should first sort the items in term of cost density against
weight. Then fill up as much of the most valuable substance by weight as one can hold, then as much of the next most
valuable substance, etc. Until W is reached.
Item A B C
Cost 300 190 180
Weight 3 2 2
Cost/weight 100 95 90
For a knapsack of capacity of 4 kg.
The optimum solution of the above will take 3kg of A and 1 kg of B.
Algorithm FractionalKnapsack(W[], C[], Wk)
For i = 1 to n do
X[i] = 0
Weight = 0
//Use Max heap
H = BuildMaxHeap(C/W)
While Weight < Wk do
i = H.GetMax()
If(Weight + W[i] <= Wk) do
X[i] = 1
Weight = Weight + W[i]
Else
X[i] = (Wk – Weight)/W[i]
Weight = Wk
Return X
0/1 Knapsack Problem
A thief can only take or leave the item. He cannot take a fraction.
A greedy strategy same as above could result in empty space, reducing the overall cost density of the knapsack.
In the above example, after choosing object A there is no place for B or C so there leaves empty space of 1kg.
Moreover, the result of the greedy solution is not optimal.
The optimal solution will be when we take object B and C. This problem can be solved by dynamic programming that
we will see in the*/
