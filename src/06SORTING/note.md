Generalized Bucket Sort
There are cases when the element falling into a bucket are not unique but are in the same range. When we want to sort
an index of a name, we can use the reference bucket to store names.
The buckets are already sorted and the elements inside each bucket can be kept sorted by using an Insertion-Sort
algorithm. We are leaving this generalized bucket sort implementation to the reader of this book. The similar data
structure will be defined in the coming chapter of Hash-Table using separate chaining.
Heap-Sort
Heap-Sort we will study in the Heap chapter.
Complexity Analysis:
Data structure List
Worst case performance O(nlogn)
Average case performance O(nlogn)
Worst case Space Complexity O(1)
Tree Sorting
In-order traversal of the binary search tree can also be seen as a sorting algorithm. We will see this in binary search tree
section of tree chapter.
Complexity Analysis:
Worst Case Time Complexity O(n2)
Best Case Time Complexity O(nlogn)
Average Time Complexity O(nlogn)
Space Complexity O(n)
Stable Sorting Yes