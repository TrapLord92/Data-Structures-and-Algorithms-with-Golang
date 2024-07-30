package main

/*Introduction
Brute Force is a straightforward approach of solving a problem based on the problem statement. It is one of the easiest
approaches to solve a particular problem. It is useful for solving small size dataset problem.
Many times, other algorithm techniques can be used to get a better solution of the same problem.
Some examples of brute force algorithms are:
· Bubble-Sort
· Selection-Sort
· Sequential search in a list
· Computing pow (a, n) by multiplying a, n times.
· Convex hull problem
· String matching
· Exhaustive search
· Traveling salesman
· Knapsack
· Assignment problems
Problems in Brute Force Algorithm
Bubble-Sort
In Bubble-Sort, adjacent elements of the list are compared and are exchanged if they are out of order.
// Sorts a given list by Bubble Sort
// Input: A list A of orderable elements
// Output: List A[0..n - 1] sorted in ascending order
Algorithm BubbleSort(A[0..n - 1])
sorted = false
while !sorted do
sorted = true
for j = 0 to n - 2 do
if A[j] > A[j + 1] then
swap A[j] and A[j + 1]
sorted = false
The Time Complexity of the algorithm is Θ(n2)
Selection-Sort
The entire given list of N elements is traversed to find its smallest element and exchange it with the first element. Then,
the list is traversed again to find the second element and exchanged it with the second element. After N-1 passes, the
list will be fully sorted.
//Sorts a given list by selection sort
//Input: A list A[0..n-1] of orderable elements
//Output: List A[0..n-1] sorted in ascending order
Algorithm SelectionSort (A[0..n-1])
for i = 0 to n - 2 do
min = i
for j = i + 1 to n - 1 do
if A[j] < A[min]
min = j
swap A[i] and A[min]
The Time Complexity of the algorithm is Θ(n2)
Sequential Search
The algorithm compares consecutive elements of a given list with a given search keyword until either a match is found
or the list is exhausted.
Algorithm SequentialSearch (A[0..n], K)
i = 0
While A [i] ≠ K do
i = i + 1
if i < n
return i
else
return -1
Worst case Time Complexity is Θ (n).
Computing pow (a, n)
Computing an (a > 0, and n is a nonnegative integer) based on the definition of exponentiation.
N-1 multiplications are required in brute force method.
// Input: A real number a and an integer n = 0
// Output: a power n
Algorithm Power(a, n)
result = 1
for i = 1 to n do
result = result * a
return result
The algorithm requires Θ (n)
String matching
A brute force string matching algorithm takes two inputs, first text consists of n characters and a pattern consist of m
character (m<=n). The algorithm starts by comparing the pattern with the beginning of the text. Each character of the
patters is compared to the corresponding character of the text. Comparison starts from left to right until all the
characters are matched or a mismatch is found. The same process is repeated until a match is found. Each time the
comparison starts one position to the right.
// Input: A list T[0..n - 1] of n characters representing a text
// a list P[0..m - 1] of m characters representing a pattern
// Output: The position of the first character in the text that starts the first
// matching substring if the search is successful and -1 otherwise.
Algorithm BruteForceStringMatch (T[0..n - 1], P[0..m - 1])
for i = 0 to n - m do
j = 0
while j < m and P[j] = T[i + j] do
j = j + 1
if j = m then
return i
return -1
In the worst case, the algorithm is O(mn).
Closest-Pair Brute-Force Algorithm
The closest-pair problem is to find the two closest points in a set of n points in a 2-dimensional space.
A brute force implementation of this problem computes the distance between each pair of distinct points and find the
smallest distance pair.
// Finds two closest points by brute force
// Input: A list P of n >= 2 points
// Output: The closest pair
Algorithm BruteForceClosestPair(P)
dmin = infinite
for i = 1 to n - 1 do
for j = i + 1 to n do
d = (xi - xj)2 + (yi - yj)2
if d < dmin then
dmin = d
imin = i
jmin = j
return imin, jmin
In the Time Complexity of the algorithm is Θ(n2)
Convex-Hull Problem
Convex-hull of a set of points is the smallest convex polygon containing all the points. All the points of the set will lie
on the convex hull or inside the convex hull. Illustrate the rubber-band interpretation of the convex hull. The convexhull
of a set of points is a subset of points in the given sets.
How to find this subset?
Answer: The rest of the points of the set are all on one side.
Two points (x1, y1), (x2, y2) make the line ax + by = c
Where a = y2-y1, b = x1-x2, and c = x1y2 - y1x2
And divides the plane by ax + by - c < 0 and ax + by - c > 0
So we need to only check ax + by - c for the rest of the points
If we find all the points in the set lies one side of the line with either all have ax + by - c < 0 or all the points have ax +
by - c > 0 then we will add these points to the desired convex hull point set.
For each of n (n -1) /2 pairs of distinct points, one needs to find the sign of ax + by - c in each of the other n - 2 points.
What is the worst-case cost of the algorithm: O(n3)
Algorithm ConvexHull
for i=0 to n-1
for j=0 to n-1
if (xi,yi) !=(xj,yj)
draw a line from (xi,yi) to (xj,yj)
for k=0 to n-1
if(i!=k and j!=k)
if ( all other points lie on the same side of the
line (xi,yi) and (xj,yj))
then add (xi,yi) to (xj,yj) to the convex hull set
Exhaustive Search
Exhaustive search is a brute force approach applies to combinatorial problems.
In exhaustive search, we generate all the possible combinations. See if the combinations satisfy the problem constraints
and then finding the desired solution.
Examples of exhaustive search are:
· Traveling salesman problem
· Knapsack problem
· Assignment problem
Traveling Salesman Problem (TSP)
In the traveling salesman problem we need to find the shortest tour through a given set of N cities that salesperson visits
each city exactly once before returning to the city where he started.
Alternatively, finding the shortest Hamiltonian circuit in a weighted connected graph. A cycle that passes through all
the vertices of the graph exactly once.
Tours where A is starting city:
Tour Cost
A→B→C→D→A 1+3+6+5 = 15
A→B→D→C→A 1+4+6+8 = 19
A→C→B→D→A 8+3+4+5 = 20
A→C→D→B→A 8+6+4+1 = 19
A→D→B→C→A 5+4+3+8 = 20
A→D→C→B→A 5+6+3+1 = 15
Algorithm TSP
Select a city
MinTourCost = infinite
For ( All permutations of cities ) do
If( LengthOfPathSinglePermutation < MinTourCost )
MinTourCost = LengthOfPath
Total number of possible combinations = (n-1)!
Cost for calculating the path: Θ(n)
So the total cost for finding the shortest path: Θ(n!)
Knapsack Problem
Given an item with cost C1, C2,..., Cn, and volume V1, V2,..., Vn and knapsack of capacity Vmax, find the most
valuable (max ΣCj) that fit in the knapsack (ΣVj ≤ Vmax).
The solution is one of all the subset of the set of object taking 1 to n objects at a time, so the Time Complexity will be
O(2n)
Algorithm KnapsackBruteForce
MaxProfit = 0
For ( All permutations of objects ) do
CurrProfit = sum of objects selected
If( MaxProfit < CurrProfit )
MaxProfit = CurrProfit
Store the current set of objects selected
Conclusion
Brute force is the first algorithm that comes into mind when we see some problem. They are the simplest algorithms
that are very easy to understand. However, these algorithms rarely provide an optimum solution. Many cases we will
find other effective algorithm that is more efficient than the brute force method. This is the most simple to understand
the kind of problem solving technique.*/
