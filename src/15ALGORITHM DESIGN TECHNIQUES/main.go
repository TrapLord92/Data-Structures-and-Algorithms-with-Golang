package main

/*Introduction
In real life when we are asked to do some work, we try to correlate it with our experience and then try to solve it.
Similarly, when we get a new problem to solve. We first try to find the similarity of the current problem with some
problems for which we already know the solution. Then solve the current problem and get our desired result.
This method provides following benefits:
1) It provides a template for solving a wide range of problems.
2) It provides us the idea of the suitable data structure for the problem.
3) It helps us in analyzing, space and Time Complexity of algorithms.
In the previous chapters, we have used various algorithms to solve different kind of problems. In this chapter, we will
read about various techniques of solving algorithmic problems.
Various Algorithm design techniques are:
1) Brute Force
2) Greedy Algorithms
3) Divide-and-Conquer, Decrease-and-Conquer
4) Dynamic Programming
5) Reduction / Transform-and-Conquer
6) Backtracking and Branch-and-Bound
Brute Force Algorithm
Brute Force is a straightforward approach of solving a problem based on the problem statement. It is one of the easiest
approaches to solve a particular problem. It is useful for solving small size dataset problem.
Some examples of brute force algorithms are:
· Bubble-Sort
· Selection-Sort
· Sequential search in a list
· Computing pow(a, n) by multiplying a, n times.
· Convex hull problem
· String matching
· Exhaustive search: Traveling salesman, Knapsack, and Assignment problems
Greedy Algorithm
In greedy algorithm, solution is constructed through a sequence of steps. At each step, choice is made which is locally
optimal. Greedy algorithms are generally used to solve optimization problems. We always take the next data to be
processed depending upon the dataset which we have already processed and then choose the next optimum data to be
processed. Greedy algorithms does not always give optimum solution.
Some examples of brute force algorithms are:
· Minimal spanning tree: Prim’s algorithm, Kruskal’s algorithm
· Dijkstra’s algorithm for single-source shortest path problem
· Greedy algorithm for the Knapsack problem
· The coin exchange problem
· Huffman trees for optimal encoding
Divide-and-Conquer, Decrease-and-Conquer
Divide-and-Conquer algorithms involve basic three steps, first split the problem into several smaller sub-problems,
second solve each sub problem and then finally combine the sub problems results to produce the result.
In divide-and-conquer the size of the problem is reduced by a factor (half, one-third, etc.), While in decrease-andconquer
the size of the problem is reduced by a constant.
Examples of divide-and-conquer algorithms:
· Merge-Sort algorithm (using recursion)
· Quicksort algorithm (using recursion)
· Computing the length of the longest path in a binary tree (using recursion)
· Computing Fibonacci numbers (using recursion)
· Quick-hull
Examples of decrease-and-conquer algorithms:
· Computing pow(a, n) by calculating pow(a, n/2) using recursion.
· Binary search in a sorted list (using recursion)
· Searching in BST
· Insertion-Sort
· Graph traversal algorithms (DFS and BFS)
· Topological sort
· Warshall’s algorithm (using recursion)
· Permutations (Minimal change approach, Johnson-Trotter algorithm)
· Computing a median, Topological sorting, Fake-coin problem (Ternary search)
Consider the problem of exponentiation Compute xn
Brute Force: n-1 multiplications
Divide and conquer: T(n) = 2*T(n/2) + 1 = n-1
Decrease by one: T (n) = T (n-1) + 1 = n-1
Decrease by constant factor: T (n) = T (n/a) + a-1 = (a-1) n = n when a = 2
Dynamic Programming
While solving problems using Divide-and-Conquer method, there may be a case when recursively sub-problems can
result in the same computation being performed multiple times. This problem arises when there are identical subproblems
arise repeatedly in a recursion.
Dynamic programming is used to avoid the requirement of repeated calculation of same sub-problem. In this method,
we usually store the result of sub - problems in a table and refer that table to find if we have already calculated the
solution of sub - problems before calculating it again.
Dynamic programming is a bottom up technique in which the smaller sub-problems are solved first and the result of
these are sued to find the solution of the larger sub-problems.
Examples:
· Fibonacci numbers computed by iteration.
· Warshall’s algorithm for transitive closure implemented by iterations
· Floyd’s algorithms for all-pairs shortest paths
func fibonacci(n int) int {
if n <= 1 {
return n
}
return fibonacci(n-1) + fibonacci(n-2)
}
Using divide and conquer the same sub problem is solved again and again, which reduce the performance of the
algorithm. This algorithm has an exponential Time Complexity and linear Space Complexity.
func fibonacci(n int) int {
first := 0
second := 1
temp := 0
if n == 0 {
return first
} else if n == 1 {
return second
}
for i := 2; i <= n; i++ {
temp = first + second
first = second
second = temp
}
return temp
}
Using this algorithm, we will get Fibonacci in linear Time Complexity and constant Space Complexity.
Reduction / Transform-and-Conquer
These methods works as two-stage procedure. First, the problem is transformed into a known problem for which we
know optimal solution. In the second stage, the problem is solved.
The most common types of transformation are sort of a list. For example, Given a list of numbers finds the two closest
number.
The brute force solution for this problem will take distance between each element in the list and will try to keep the
minimum distance pair; this approach will have a Time Complexity of O(n2)
Transform and conquer solution, will be first sort the list in O(nlogn) time and then find the closest number by
scanning the list in another O(n). Which will give the total Time Complexity of O(nlogn).
Examples:
· Gaussian elimination
· Heaps and Heapsort
Backtracking
In real life, let us suppose someone gave you a lock with a number (three digit lock, number range from 1 to 9).
Moreover, you do not have the exact password key for the lock. You need to test every combination until you got the
right one. Obviously, you need to test starting from something like “111”, then “112” and so on. You will get your key
before you reach “999”. Therefore, what you are doing is backtracking.
Suppose the lock produce some sound “click” correct digit is selected for any level. If we can listen to this sound such
intelligence/ heuristics will help you to reach your goal much faster. These functions are called Pruning function or
bounding functions.
Backtracking is a method by which solution is found by exhaustively searching through large but finite number of
states, with some pruning or bounding function that will narrow down our search.
For all the problems like (NP hard problems) for which there does not exist any other method we use backtracking.
Backtracking problems have the following components:
1. Initial state
2. Target / Goal state
3. Intermediate states
4. Path from the initial state to the target / goal state
5. Operators to get from one state to another
6. Pruning function (optional)
The solving process of backtracking algorithm starts with the construction of state’s tree, whose nodes represents the
states. The root node is the initial state and one or more leaf node will be our target state. Each edge of the tree
represents some operation. The solution is obtained by searching the tree until a Target state is found.
Backtracking uses depth-first search:
1) Store the initial state in a stack
2) While the stack is not empty, repeat:
3) Read a node from the stack.
4) While there are available operators, do:
a. Apply an operator to generate a child
b. If the child is a goal state – return solution
c. If it is a new state, and pruning function does not discard it push the child into the stack.
There are three monks and three demons at one side of a river. We want to move all of them to the other side using a
small boat. The boat can carry only two persons at a time. Given if on any shore the number of demons will be more
than monks then they will eat the monks. How can we move all of these people to the other side of the river safely?
Same as the above problem there is a farmer who has a goat, a cabbage and a wolf. If the farmer leaves, goat with
cabbage, goat will eat the cabbage. If the farmer leaves wolf alone with goat, wolf will kill the goat. How can the
farmer move all his belongings to the other side of the river?
You are given two jugs, a 4-gallon one and a 3-gallon one. There are no measuring markers on jugs. A tap can be used
to fill the jugs with water. How can you get 2 gallons of water in the 4-gallon jug?
Branch-and-bound
Branch and bound method is used when we can evaluate cost of visiting each node by a utility functions. At each step,
we choose the node with lowest cost to proceed further. Branch-and bound algorithms are implemented using a priority
queue. In branch and bound, we traverse the nodes in breadth-first manner.
A* Algorithm
A* is sort of an elaboration on branch-and-bound. In branch-and-bound, at each iteration we expand the shortest path
that we have found so far. In A*, instead of just picking the path with the shortest length so far, we pick the path with
the shortest estimated total length from start to goal, where the total length is estimated as length traversed so far plus a
heuristic estimate of the remaining distance from the goal.
Branch-and-bound will always find an optimal solution, which is shortest path. A* will always find an optimal solution
if the heuristic is correct. Choosing a good heuristic is the most important part of A* algorithm.
Conclusion
Usually a given problem can be solved using a number of methods; however, it is not wise to settle for the first method
that comes to our mind. Some methods result in a much more efficient solution than others do.
For example, the Fibonacci numbers calculated recursively (decrease-and-conquer approach), and computed by
iterations (dynamic programming). In the first case, the complexity is O( 2n), and in the other case, the complexity is
O(n).
Another example, consider sorting based on the Insertion-Sort and basic bubble sort. For almost sorted files, Insertion-
Sort will give almost linear complexity, while bubble sort sorting algorithms have quadratic complexity.
So the most important question is how to choose the best method?
First, you should understand the problem statement.
Second by knowing various problems and their solutions.*/
