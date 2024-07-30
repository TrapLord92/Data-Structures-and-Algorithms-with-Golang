Complexity analysis of algorithms
1. Worst Case complexity: It is the complexity of solving the problem for the worst input of size n. It provides the
upper bound for the algorithm. This is the most common analysis done.
2. Average Case complexity: It is the complexity of solving the problem on an average. We calculate the time for all
the possible inputs and then take an average of it.
3. Best Case complexity: It is the complexity of solving the problem for the best input of size n.
Time Complexity Order
A list of commonly occurring algorithm Time Complexity in increasing order:
Name Notation
Constant O(1)
Logarithmic O(log n)
Linear O(n)
N-LogN O(n log n)
Quadratic O(n2)
Polynomial O(nc) c is a constant & c>1
Exponential O(cm) c is a constant & c>1
Factorial or N-power-N O(n!) or O(nn)
Constant Time: O(1)
An algorithm is said to run in constant time if the output is produced in constant time regardless of the input size.
Examples:
1. Accessing nth element of a list
2. Push and pop of a stack.
3. Enqueue and remove of a queue.
4. Accessing an element of Hash-Table.
Linear Time: O(n)
An algorithm is said to run in linear time if the execution time of the algorithm is directly proportional to the input size.
Examples:
1. List operations like search element, find min, find max etc.
2. Linked list operations like traversal, find min, find max etc.
Note: when we need to see/ traverse all the nodes of a data-structure for some task then complexity is no less than O(n)
Logarithmic Time: O(logn)
An algorithm is said to run in logarithmic time if the execution time of the algorithm is proportional to the logarithm of
the input size. Each step of an algorithm, a significant portion of the input is pruned out without traversing it.
Example: Binary search, we will read about these algorithms in this book.
N-LogN Time: O(nlog(n))
An algorithm is said to run in logarithmic time if the execution time of an algorithm is proportional to the product of
input size and logarithm of the input size.
Example:
1. Merge-Sort
2. Quick-Sort (Average case)
3. Heap-Sort
Note: Quicksort is a special kind of algorithm to sort a list of numbers. Its worst-case complexity is O(n2) and average
case complexity is O(n log n).
Quadratic Time: O(n2)
An algorithm is said to run in quadratic time if the execution time of an algorithm is proportional to the square of the
input size.
Examples:
1. Bubble-Sort
2. Selection-Sort
3. Insertion-Sort
Deriving the Runtime Function of an Algorithm
Constants
Each statement takes a constant time to run. Time Complexity is O(1)
Loops
The running time of a loop is a product of running time of the statement inside a loop and number of iterations in the
loop. Time Complexity is O(n)
Nested Loop
The running time of a nested loop is a product of running time of the statements inside loop multiplied by a product of
the size of all the loops. Time Complexity is O(nc). Where c is a number of loops. For two loops, it will be O(n2)
Consecutive Statements
Just add the running times of all the consecutive statements
If-Else Statement
Consider the running time of the larger of if block or else block. Moreover, ignore the other one.
Logarithmic statement
If each iteration the input size is decreases by a constant factors. Time Complexity