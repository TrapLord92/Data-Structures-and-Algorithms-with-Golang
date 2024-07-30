package main

/*Priority Queue (Heap)
Priority queue is implemented using a binary heap data structure. In a heap, the records are stored in an array. Each
node in the heap follows the same rule that the parent value is greater than its children are.
There are two types of the heap data structure:
1. Max heap: each node should be greater than or equal to each of its children.
2. Min heap: each node should be smaller than or equal to each of its children.
A heap is a useful data structure when you want to get max/min one by one from data. Heap-Sort uses max heap to sort
data in increasing/decreasing order.
Heap ADT Operations
· Insert() - Adding a new element to the heap. The Time Complexity of this operation is O(log(n))
· remove() - Extracting max for max heap case (or min for min heap case). The Time Complexity of this
operation is O(log(n))
· Heapify() – To convert a list of numbers in a list into a heap. This operation has a Time Complexity O(n)
PriorityQueue implementation
For implementation of priority queue, please refer Priority queue chapter.*/
