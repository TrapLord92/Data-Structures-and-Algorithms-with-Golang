package main

import "fmt"

// Types of Heap
// There are two types of heap and the type depends on the ordering of the elements. The ordering can be done in two
// ways: Min-Heap and Max-Heap
// Max Heap
// Max-Heap: the value of each node is less than or equal to the value of its parent, with the largest-value element at the
// root.
// Max Heap Operations
// Insert O(logn)
// DeleteMax O(logn)
// Remove O(logn)
// FindMax O(1)
// Min Heap
// Min-Heap: the value of each node is greater than or equal to the value of its parent, with the minimum-value element at
// the root.
// Use it whenever you need quick access to the smallest item, because that item will always be at the root of the tree or
// the first element in the list. However, the remainder of the list is kept partially sorted. Thus, instant access is only
// possible for the smallest item.
// Min Heap Operations
// Insert O(logn)
// DeleteMin O(logn)
// Remove O(logn)
// FindMin O(1)
// Throughout this chapter, the word "heap" will always refer to a max-heap. The implementation of min-heap is left for
// the user to do it as an exercise.
// Heap ADT Operations
// The basic operations of binary heap are as follows:
// Binary Heap Create a new empty binary heap O(1)
// Insert Adding a new element to the heap O(logn)
// DeleteMax Delete the maximum element form the heap. O(logn)
// FindMax Find the maximum element in the heap. O(1)
// isEmpty return true if the heap is empty else return false O(1)
// Size Return the number of elements in the heap. O(1)
// BuildHeap Build a new heap from the list of elements O(logn)
// Operation on Heap
// Create Heap from a list
// Heapify is the process of converting a list into Heap. The various steps are:
// 1. Values are present in list.
// 2. Starting from the middle of the list move downward towards the start of the list. At each step, compare parent value
// with its left child and right child. In addition, restore the heap property by shifting the parent value with its largestvalue
// child. Such that the parent value will always be greater than or equal to left child and right child.
// 3. For all elements from middle of the list to the start of the list. We are doing comparisons and shift, until we reach
// the leaf nodes of the heap. The Time Complexity of build heap is O(N).
// Given a list as input to create heap function. Value of index i is
// compared with value of its children nodes that is at index ( i*2 + 1 )
// and ( i*2 + 2 ). Middle of list N/2 that is index 3 is comapred with
// index 7. If the children node value is greater than parent node then
// the value will be swapped.
// Similarly, value of index 2 is compared with index 5 and 6. The
// largest of the value is 7 which will be swapped with the value at the
// index 2.
// Similarly, value of index 1 is compared with index 3 and 4 The
// largest of the value is 8 which will be swapped with the value at the
// index 1.
// Percolate down function is used to subsequently adjest the value
// replased in the previous step by comparing it with its children nodes.
// Now value at index 0 is comared with index 1 and 2. 8 is the largest
// value so it swapped with the value at index 0.
// Percolate down function is used to further compare the value at
// index 1 with its children nodes at index 3 and 4.
// In the end max heap is created.
// Example 11.1:
type Heap struct {
	size  int
	arr   []int
	isMin bool
}

func NewHeap(arrInput []int, isMin bool) *Heap {
	size := len(arrInput)
	arr := []int{1}
	arr = append(arr, arrInput...)
	h := &Heap{size: size, arr: arr, isMin: isMin}
	for i := (h.size / 2); i > 0; i-- {
		h.proclateDown(i)
	}
	return h
}
func (h *Heap) proclateDown(parent int) {
	lChild := 2 * parent
	rChild := lChild + 1
	small := -1
	if lChild <= h.size {
		small = lChild
	}
	if rChild <= h.size && h.comp(lChild, rChild) {
		small = rChild
	}
	if small != -1 && h.comp(parent, small) {
		h.swap(parent, small)
		h.proclateDown(small)
	}
}
func NewHeap2(isMin bool) *Heap {
	arr := []int{1}
	h := &Heap{size: 0, arr: arr, isMin: isMin}
	return h
}
func (h *Heap) comp(i, j int) bool { // always i < j in use
	if h.isMin == true {
		return h.arr[i] > h.arr[j] // swaps for min heap
	}
	return h.arr[i] < h.arr[j] // swap for max heap.
}
func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
func (h *Heap) Empty() bool {
	return (h.size == 0)
}
func (h *Heap) Size() int {
	return h.size
}
func (h *Heap) Peek() (
	int, bool) {
	if h.Empty() {
		fmt.Println("Heap empty Error.")
		return 0, false
	}
	return h.arr[1], true
}

// Initializing an empty Heap
// Example 11.2:
func (h *Heap) proclateUp(child int) {
	parent := child / 2
	if parent == 0 {
		return
	}
	if h.comp(parent, child) {
		h.swap(child, parent)
		h.proclateUp(parent)
	}
}

//Operations

/*Enqueue / Insert
1. Add the new element at the end of the list. This keeps the structure as a complete binary tree, but it might no longer
be a heap since the new element might have a value greater than its parent’s value.
2. Swap the new element with its parent until it has value greater than its parent’s value.
3. Step 2 will terminate when the new element reaches the root or when the new element's parent have a value greater
than or equal to the new element's value.
Let us take an example of the Max heap created in the above example.
Let us take an example by inserting element with value 9 to the heap. The element is added to the end of the heap list.
Now the value will be percolate up by comparing it with the parent. The value is added to index 8 and its parent will be
(N-1)/2 = index 3.
Since the value 9 is greater thn 4 it will be swapped with it.
Percolate up is used and the value is moved up until heap property is
satisfied.
Now the value at index 1 is compared with index 0 and to satisfy heap
property it is further swapped.
Now finally max heap is created by inserting new node.
Example 11.3:*/
func (h *Heap) Add(value int) {
	h.size++
	h.arr = append(h.arr, value)
	h.proclateUp(h.size)
}

/*Dequeue / Delete
1. Copy the value at the root of the heap to the variable used to return a value.
2. Copy the last element of the heap to the root, and then reduce the size of heap by 1. This element is called the "outof-
place" element.
3. Restore heap property by swapping the out-of-place element with its largest-value child. Repeat this process until
the out-of-place element reaches a leaf or its value is greater or equal to all its children.
4. Return the answer that was saved in Step 1.
To remove an element from heap its top & end value is swapped and size of heap is reduced by 1.
Since value at end of the heap is copied to head of heap. Heap property
is disturbed so we need to percolate down by comparing node with its
children nodes and restore heap property.
Percolate down continued by comparing with its children nodes.
Percolate down
Percolate down Complete
Example 11.4:*/
func (h *Heap) Remove() (int, bool) {
	if h.Empty() {
		fmt.Println("HeapEmptyError.")
		return 0, false
	}
	value := h.arr[1]
	h.arr[1] = h.arr[h.size]
	h.size--
	h.proclateDown(1)
	h.arr = h.arr[0 : h.size+1]
	return value, true
}
func main1() {
	a := []int{1, 9, 6, 7, 8, -1, 2, 4, 5, 3}
	hp := NewHeap(nil, true)
	// hp := NewHeap(a, true)
	n := len(a)
	for i := 0; i < n; i++ {
		hp.Add(a[i])
	}
	for i := 0; i < n; i++ {
		//fmt.Println("pop value :: ", //hp.Remove())
	}
}

/*Heap-Sort
1. Use create heap function to build a max heap from the given list of elements. This operation will take O(N) time.
2. Dequeue the max value from the heap and store this value to the end of the list at location arr[size-1]
a) Copy the value at the root of the heap to end of the list.
b) Copy the last element of the heap to the root, and then reduce the size of heap by 1. This element is called
the "out-of-place" element.
c) Restore heap property by swapping the out-of-place element with its greatest-value child. Repeat this
process until the out-of-place element reaches a leaf or it has a value that is greater or equal to all its
children
3. Repeat this operation until there is just one element in the heap.
Let us take example of the heap that we had created at the start of the chapter. Heap sort is algorithm starts by creating
a heap of the given list, which is done in linear time. Then at each step head of the heap is swapped with the end of the
heap and the heap size is reduced by 1. Then percolate down is used to restore the heap property. Moreover, this same
is done multiple times until the heap contain just one element.
We had started with max heap. The maximum value as the first
element of the Heap list is swapped with the last element of the list.
Now the largest value is at the end of the list. Then we will reduce the
size of the heap by one.
Since 1 is at the top of the heap. Moreover, heap property is lost we
will use Percolate down method to regain the heap property.
Percolate down cont.
Since heap property is regained. Then we will copy the first element of
the heap array to the second last position.
Heap size is further reduced and percolate down cont.
Percolate down cont.
Again swap.
Size of heap reduced by 1 and percolate down.
Again swap.
Size of heap reduced and percolate down.
Again swap.
Size of heap reduced and percolate down.
Again swap.
Again swap.
End.
Final list, which is sorted in increasing order.
Example 11.5:*/
// func HeapSort(arrInput []int) {
// 	hp := NewHeap(arrInput, true)
// 	n := len(arrInput)
// 	for i := 0; i < n; i++ {
// 		//arrInput[i] = hp.Remove()
// 	}
//}//
// func main() {
// 	a := []int{1, 9, 6, 7, 8, -1, 2, 4, 5, 3}
// 	HeapSort(a)
// 	fmt.Println(a)
// }

/*Data structure List
Worst Case Time Complexity O(nlogn)
Best Case Time Complexity O(nlogn)
Average Time Complexity O(nlogn)
Space Complexity O(1)
Note: Heap-Sort is not a Stable sort and do not require any extra space for sorting a list.
Uses of Heap
1. Heapsort: One of the best sorting methods being in-place and log(N) time complexity in all scenarios.
2. Selection algorithms: Finding the min, max, both the min and max, median, or even the kth largest element can be
done in linear time (often constant time) using heaps.
3. Priority Queues: Heap Implemented priority queues are used in Graph algorithms like Prim’s
Algorithm and Dijkstra’s algorithm. A heap is a useful data structure when you need to remove the object with the
highest (or lowest) priority. Schedulers, timers
4. Graph algorithms: By using heaps as internal traversal data structures, run time will be reduced by polynomial
order. Examples of such problems are Prim's minimal
5. Because of the lack of references, the operations are faster than a binary tree. In addition, some more complicated
heaps (such as binomial) can be merged efficiently,which is not easy to do for a binary tree.*/
