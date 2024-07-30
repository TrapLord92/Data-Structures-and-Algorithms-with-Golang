Problems in Heap
Kth Smallest in a Min Heap
Just call DeleteMin() operation K-1 times and then again call DeleteMin() this last operation will give Kth smallest
value. Time Complexity O(KlogN)
Kth Largest in a Max Heap
Just call DeleteMax() operation K-1 times and then again call DeleteMax () this last operation will give Kth smallest
value. Time Complexity O(KlogN)
100 Largest in a Stream
There are billions of integers coming out of a stream some getInt() function is providing integers one by one. How
would you determine the largest 100 numbers?
Solution: Large hundred (or smallest hundred etc.) such problems are solved very easily using a Heap. In our case, we
will create a min heap.
1. First from 100 first integers builds a min heap.
2. Then for each coming integer compare if it is greater than the top of the min heap.
3. If not, then look for next integer. If yes, then remove the top min value from the min heap, insert the new value at
the top of the heap, use procolateDown, and move it to its proper position down the heap.
4. Every time you have largest 100 values stored in your head
Merge two Heap
How can we merge two heaps?
Solution: There is no single solution for this. Let us suppose the size of the bigger heap is N and the size of the smaller
heap is M.
1. If both heaps are comparable size, then put both heap lists in same bigger Lists. Alternatively, in one of the Lists if
they are big enough. Then apply CreateHeap() function which will take theta(N+M) time.
2. If M is much smaller than N then add() each element of M list one by one to N heap. This will take O(MlogN) the
worst case or O(M) best case.
Get Median function
Give a data structure that will provide median of given values in constant time.
Solution: We will be using two heaps, one min heap and other max heap. Max heap will contain the first half of data
and min heap will contain the second half of the data. Max heap will contain the smaller half of the data and its max
value that is at the top of the heap will be the median contender. Similarly, the Min heap will contain the larger values
of the data and its min value that is at its top will contain the median contender. We will keep track of the size of heaps.
Whenever we insert a value to heap, we will make sure that the size of two heaps differs by max one element,
otherwise we will pop one element from one and insert into another to keep them balanced.
Example 11.6:
type MedianHeap struct {
minHeap *Heap
maxHeap *Heap
}
func NewMedianHeap() *MedianHeap {
min := NewHeap(nil, true)
max := NewHeap(nil, false)
return &MedianHeap{
minHeap: min,
maxHeap: max,
}
}
func (h *MedianHeap) insert(value int) {
empty := h.maxHeap.Empty()
if empty {
h.maxHeap.Add(value)
} else {
top := h.maxHeap.Peek()
if top >= value {
h.maxHeap.Add(value)
} else {
h.minHeap.Add(value)
}
}
// size balancing
if h.maxHeap.Size() > h.minHeap.Size()+1 {
value := h.maxHeap.Remove()
h.minHeap.Add(value)
}
if h.minHeap.Size() > h.maxHeap.Size()+1 {
value := h.minHeap.Remove()
h.maxHeap.Add(value)
}
}
func (h *MedianHeap) getMedian() (int, bool) {
if h.maxHeap.Size() == 0 && h.minHeap.Size() == 0 {
fmt.Println("HeapEmptyError")
return 0, false
}
if h.maxHeap.Size() == h.minHeap.Size() {
val1 := h.maxHeap.Peek()
val2 := h.minHeap.Peek()
return (val1 + val2) / 2, true
} else if h.maxHeap.Size() > h.minHeap.Size() {
val1 := h.maxHeap.Peek()
return val1, true
} else {
val2 := h.minHeap.Peek()
return val2, true
}
}
Is Min Heap
Example 11.7: Given a list, find if it is a binary Heap is Min Heap
func IsMinHeap(arr []int) bool {
size := len(arr)
for i := 0; i <= (size-2)/2; i++ {
if 2*i+1 < size {
if arr[i] > arr[2*i+1] {
return false
}
}
if 2*i+2 < size {
if arr[i] > arr[2*i+2] {
return false
}
}
}
return true
}
Is Max Heap
Example 11.8: Given a list find if it is a binary Heap Max heap
func IsMaxHeap(arr []int) bool {
size := len(arr)
for i := 0; i <= (size-2)/2; i++ {
if 2*i+1 < size {
if arr[i] < arr[2*i+1] {
return false
}
}
if 2*i+2 < size {
if arr[i] < arr[2*i+2] {
return false
}
}
}
return true
Analysis: If each parent value is greater than its children value then heap property is true. We will traverse from start to
half of the array and compare the value of index node with its left child and right child node.
Traversal in Heap
Heaps are not designed to traverse to find some element they are made to get min or max element fast. Still if you want
to traverse a heap just traverse the list sequentially. This traversal will be level order traversal. This traversal will have
linear Time Complexity.
Deleting Arbiter element from Min Heap
Again, heap is not designed to delete an arbitrary element, but still if you want to do so. Find the element by linear
search in the heap list. Replace it with the value stored at the end of the Heap value. Reduce the size of the heap by
one. Compare the new inserted value with its parent. If its value is smaller than the parent value, then percolate up.
Else if its value is greater than its left and right child then percolate down. Time Complexity is O(logn)
Deleting Kth element from Min Heap
Again, heap is not designed to delete an arbitrary element, but still if you want to do so. Replace the kth value with the
value stored at the end of the Heap value. Reduce the size of the heap by one. Compare the new inserted value with its
parent. If its value is smaller than the parent value, then percolate up. Else if its value is greater than its left and right
child then percolate down. Time Complexity is O(logn)
Print value in Range in Min Heap
Linearly traverse through the heap and print the value that are in the given range.
Priority Queue generic implementation.
Let us look into priority queue implementation, which internally implements heap. We will Add() items and their
priority to the queue.
While removing items form priority queue using Remove() function, highest priority item is removed first and so on.
Example 11.9:
type PQueue struct {
items []*item
Count int
isMin bool
}
type item struct {
value interface{}
priority int
}
func newItem(value interface{}, priority int) *item {
return &item{
value: value,
priority: priority,
}
}
func NewPQueue(isMin bool) *PQueue {
items := make([]*item, 1)
items[0] = nil // Heap queue first element should always be nil
return &PQueue{
items: items,
Count: 0,
isMin: isMin,
}
}
func (pq *PQueue) comp(i, j int) bool { // always i < j in use
if pq.isMin == true {
return pq.items[i].priority > pq.items[j].priority // swaps for min heap
}
return pq.items[i].priority < pq.items[j].priority // swap for max heap.
}
func (pq *PQueue) proclateDown(position int) {
lChild := 2 * position
rChild := lChild + 1
small := -1
if lChild <= pq.Count {
small = lChild
}
if rChild <= pq.Count && pq.comp(lChild, rChild) {
small = rChild
}
if small != -1 && pq.comp(position, small) {
pq.items[position], pq.items[small] = pq.items[small], pq.items[position]
pq.proclateDown(small)
}
}
func (pq *PQueue) proclateUp(position int) {
parent := position / 2
if parent == 0 {
return
}
if pq.comp(parent, position) {
pq.items[position], pq.items[parent] = pq.items[parent], pq.items[position]
pq.proclateUp(parent)
}
}
func (pq *PQueue) Add(value interface{}, priority int) {
item := newItem(value, priority)
pq.items = append(pq.items, item)
pq.Count++
pq.proclateUp(pq.Count)
}
func (pq *PQueue) Print() {
n := pq.Count
for i := 1; i <= n; i++ {
fmt.Print(*(pq.items[i]), " ")
}
fmt.Println()
}
func (pq *PQueue) Remove() (interface{}, bool) {
if pq.IsEmpty() {
fmt.Println("Heap Empty Error.")
return 0, false
}
value := pq.items[1].value
pq.items[1] = pq.items[pq.Count]
pq.items = pq.items[0:pq.Count]
pq.Count--
pq.proclateDown(1)
return value, true
}
func (pq *PQueue) IsEmpty() bool {
return (pq.Count == 0)
}
func (pq *PQueue) Len() int {
return pq.Count
}
func (pq *PQueue) Peek() (interface{}, bool) {
if pq.IsEmpty() {
fmt.Println("Heap empty Error.")
return 0, false
}
return pq.items[1].value, true
}
func main() {
pq := NewPQueue(true)
pq.Add("banana", 2)
pq.Add("apple", 1)
pq.Add("orange", 4)
pq.Add("mango", 3)
fmt.Println(pq.Len())
for pq.IsEmpty() == false {
fmt.Println(pq.Remove())
}
}
Output:
4 apple true
banana true
mango true
orange true
Analysis:
In this implementation, the elements of heap have two fields, first field of type interface{} which will store data and
second field is of type integer which is priority. Elements will be removed according to their priority order.
Heap operations like Add(), Remove(), proclateUp(), procolateDown() are already discussed in Heap section.
Priority Queue using heap from container.
Example 11.10:
import "container/heap"
type Item struct {
value interface{}
priority int
}
type ItemList []*Item
func (lst ItemList) Len() int {
return len(lst)
}
func (lst ItemList) Less(i, j int) bool {
return lst[i].priority < lst[j].priority
}
func (lst ItemList) Swap(i, j int) {
lst[i], lst[j] = lst[j], lst[i]
}
func (lst *ItemList) Push(val interface{}) {
item := val.(*Item)
*lst = append(*lst, item)
}
func (lst *ItemList) Pop() interface{} {
old := *lst
n := len(old)
item := old[n-1]
*lst = old[0 : n-1]
return item
}
type PQueue struct {
pq ItemList
}
func NewPQueue() *PQueue {
queue := new(PQueue)
queue.pq = make(ItemList, 0)
heap.Init(&queue.pq)
return queue
}
func (queue *PQueue) Init() {
queue.pq = make(ItemList, 0)
heap.Init(&queue.pq)
}
func (queue *PQueue) Add(value interface{}, priority int) {
heap.Push(&queue.pq, &Item{value: value, priority: priority})
}
func (queue *PQueue) Remove() interface{} {
return heap.Pop(&queue.pq).(*Item).value
}
func (queue *PQueue) Len() int {
return queue.pq.Len()
}
func (queue *PQueue) IsEmpty() bool {
return queue.pq.Len() == 0
}
func main() {
pq := NewPQueue()
// pq := new(PQueue)
// pq.Init()
pq.Add("banana", 2)
pq.Add("apple", 1)
pq.Add("orange", 4)
pq.Add("mango", 3)
fmt.Println(pq.Len())
for pq.IsEmpty() == false {
fmt.Println(pq.Remove())
}
}
Exercise
1. What is the worst-case runtime Complexity of finding the smallest item in a min-heap?
2. Find max in a min heap.
Hint: normal search in the complete list. There is one more optimization you can search from the mid of the list at
index N/2
3. What is the worst-case time Complexity of finding the largest item in a min-heap?
4. What is the worst-case time Complexity of deleteMin in a min-heap?
5. What is the worst-case time Complexity of building a heap by insertion?
6. Is a heap full or complete binary tree?
7. What is the worst time runtime Complexity of sorting a list of N elements using heapsort?
8. Given a sequence of numbers: 1, 2, 3, 4, 5, 6, 7, 8, 9
a. Draw a binary Min-heap by inserting the above numbers one by one
b. Also draw the tree that will be formed after calling Dequeue() on this heap
9. Given a sequence of numbers: 1, 2, 3, 4, 5, 6, 7, 8, 9
a. Draw a binary Max-heap by inserting the above numbers one by one
b. Also draw the tree that will be formed after calling Dequeue() on this heap
10. Given a sequence of numbers: 3, 9, 5, 4, 8, 1, 5, 2, 7, 6. Construct a Min-heap by calling CreateHeap function.
11. Show a list that would be the result after the call to deleteMin() on this heap
12. Given a list: [3, 9, 5, 4, 8, 1, 5, 2, 7, 6]. Apply heapify over this to make a min heap and sort the elements in
decreasing order?
13. In Heap-Sort once a root element has been put in its final position, how much time, does it take to re-heapify the
structure so that the next removal can take place? In other words, what is the Time Complexity of a single element removal from the heap of size N?
14. What do you think the overall Time Complexity for heapsort is? Why do you feel this way?