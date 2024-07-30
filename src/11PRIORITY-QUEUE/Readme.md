Introduction
A Priority-Queue also knows as Binary-Heap, is a variant of queue. Items are removed from the start of the queue.
However, in a Priority-Queue the logical ordering of objects is determined by their priority. The highest priority item
are at the front of the Priority-Queue. When you add an item to the Priority-Queue, the new item can more to the front
of the queue. A Priority-Queue is a very important data structure. Priority-Queue is used in various Graph algorithms
like Prim’s Algorithm and Dijkstra’s algorithm. Priority-Queue is also used in the timer implementation etc.
A Priority-Queue is implemented using a Heap (Binary Heap). A Heap data structure is an array of elements that can be
observed as a complete binary tree. The tree is completely filled on all levels except possibly the lowest. Heap satisfies
the heap ordering property. In max-heap, the parent’s value is greater than or equal to its children value. In min-heap,
the parent’s value is less than or equal to its children value. A heap is a complete binary tree so the height of tree with
N nodes is always O(logn).
A heap is not a sorted data structure and can be regarded as partially ordered. As you see from the picture, there is no
relationship among nodes at any given level, even among the siblings.
Heap is implemented using an array. Moreover, because heap is a complete binary tree, the left child of a parent (at
position x) is the node that is found in position 2x in the array. Similarly, the right child of the parent is at position 2x+1
in the array. To find the parent of any node in the heap, we can simply division. Given the index y of a node, the parent
index will by y/2.