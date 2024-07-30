package main

# Search Element in a Linked-List
Search element in linked list. Given a head reference and value. Returns true if value found in list else returns false.
Note: Search in a single linked list can be only done in one direction. Since all elements in the list have reference to the
next item in the list. Therefore, traversal of linked list is linear in nature.
Example 7.9:

func (list *List) IsPresent(data int) bool {
temp := list.head
for temp != nil {
if temp.value == data {
return true
}
temp = temp.next
}
return false
}
Analysis:
· We create a temp variable, which will point to head of the list.
· Using a “for” loop we will iterate through the list.
· Value of each element of list is compared with the given value. If value is found, then the function will return true.
· If the value is not found, then false will be returned from the function in the end.

# Delete element from the linked list
## Delete First element in a linked list.
Example 7.10:
func (list *List) RemoveHead() (int, bool) {
if list.IsEmpty() {
fmt.Println("EmptyListError")
return 0, false
}
value := list.head.value
list.head = list.head.next
list.count--
return value, true
}
Analysis:
· First, we need to check if the list is empty. If list is already empty then we print error statement and return error.
· If list is not empty then store the value of head node in a temporary variable value.
· We need to find the second element of the list and assign it as head of the linked list.
· Since the first node is no longer referenced so it will be automatically deleted.
· Decrease the size of list. Then return the value stored in temporary variable value.

## Delete node from the linked list given its value.
Example 7.11:
func (list *List) DeleteNode(delValue int) bool {
temp := list.head
if list.IsEmpty() {
fmt.Println("EmptyListError")
return false
}
if delValue == list.head.value {
list.head = list.head.next
list.count--
return true
}
for temp.next != nil {
if temp.next.value == delValue {
temp.next = temp.next.next
list.count--
return true
}
temp = temp.next
}
return false
}

Analysis:
· If the list is empty then we will return false from the function which indicate that the deleteNode() method executed
with error.
· If the node that need to be deleted is head node then head reference need is modified and point to the next node.
· In a for loop we will traverse the link list and try to find the node that need to be deleted. If the node is found then,
we will point its reference to the node next to it and return true.
· If the node is not found then we will return false.

## Delete all the occurrence of particular value in linked list.
Example 7.12:
func (list *List) DeleteNodes(delValue int) {
currNode := list.head
for currNode != nil && currNode.value == delValue {
list.head = currNode.next
currNode = list.head
}
for currNode != nil {
nextNode := currNode.next
if nextNode != nil && nextNode.value == delValue {
currNode.next = nextNode.next
} else {
currNode = nextNode
}
}
}
Analysis:
· In the first for loop will delete all the nodes that are at the front of the list, which have valued equal to delValue. In
this, we need to update head of the list.
· In the second for loop, we will be deleting all the nodes that are having value equal to the delValue. Remember that
we are not returning even though we have the node that we are looking for.

## Delete a single linked list
### Delete all the elements of a linked list, given a reference to head of linked list.
Example 7.13:
func (list *List) FreeList() {
list.head = nil
list.count = 0
}
Analysis: We just need to point head to null. The reference to the list is lost so it will automatically deleted.
Reverse a linked list.
Reverse a singly linked List iteratively using three Pointers
Example 7.14:
func (list *List) Reverse() {
curr := list.head
var prev, next *Node
for curr != nil {
next = curr.next
curr.next = prev
prev = curr
curr = next
}
list.head = prev
}
Analysis: The list is iterated. Make next equal to the next node of the curr node. Make curr node’s next point to prev
node. Then iterate the list by making prev point to curr and curr point to next.
Recursively Reverse a singly linked List
Reverse a singly linked list using Recursion.
Example 7.15:
func (list *List) ReverseRecurse() {
list.head = list.reverseRecurseUtil(list.head, nil)
}
func (list *List) reverseRecurseUtil(currentNode *Node, nextNode *Node) *Node {
var ret *Node
if currentNode == nil {
return nil
}
if currentNode.next == nil {
currentNode.next = nextNode
return currentNode
}
ret = list.reverseRecurseUtil(currentNode.next, currentNode)
currentNode.next = nextNode
return ret
}
Analysis:
· ReverseRecurse function will call a reverseRecurseUtil function to reverse the list and the reference returned by the
reverseRecurseUtil will be the head of the reversed list.
· The current node will point to the nextNode that is previous node of the old list.
Note: A linked list can be reversed using two approaches the first approach is by using three references. The Second
approach is using recursion both are linear solution, but three-reference solution is more efficient.

## Remove duplicates from the linked list
### Remove duplicate values from the linked list. The linked list is sorted and it contains some duplicate values, you need
to remove those duplicate values. (You can create the required linked list using SortedInsert() function)
Example 7.16:
func (list *List) RemoveDuplicate() {
curr := list.head
for curr != nil {
if curr.next != nil && curr.value == curr.next.value {
curr.next = curr.next.next
} else {
curr = curr.next
}
}
}
Analysis: for loop is used to traverse the list. Whenever there is a node whose value is equal to the next node’s value,
that current node next will point to the next of next node. Which will remove the next node from the list.

# Copy List Reversed
## Copy the content of linked list in another linked list in reverse order. If the original linked list contains elements in
order 1,2,3,4, the new list should contain the elements in order 4,3,2,1.
Example 7.17:
func (list *List) CopyListReversed() *List {
var tempNode, tempNode2 *Node
curr := list.head
for curr != nil {
tempNode2 = &Node{curr.value, tempNode}
curr = curr.next
tempNode = tempNode2
}
ll2 := new(List)
ll2.head = tempNode
return ll2
}
Analysis: Traverse the list and add the node’s value to the new list. Since the list is traversed in the forward direction
and each node’s value is added to another list so the formed list is reverse of the given list.

## Copy the content of given linked list into another linked list
### Copy the content of given linked list into another linked list. If the original linked list contains elements in order
1,2,3,4, the new list should contain the elements in order 1,2,3,4.
Example 7.18:
func (list *List) CopyList() *List {
var headNode, tailNode, tempNode *Node
curr := list.head
if curr == nil {
ll2 := new(List)
ll2.head = nil
return ll2
}
headNode = &Node{curr.value, nil}
tailNode = headNode
curr = curr.next
for curr != nil {
tempNode = &Node{curr.value, nil}
tailNode.next = tempNode
tailNode = tempNode
curr = curr.next
}
ll2 := new(List)
ll2.head = headNode
return ll2
}
Analysis: Traverse the list and add the node’s value to new list, but this time always at the end of the list. Since the list
is traversed in the forward direction and each node’s value is added to the end of another list. Therefore, the formed list
is same as the given list.
# Compare List
## Compare the values of two linked lists given their head pointers.
Example 7.19: Compare two list given
func (list *List) CompareList(ll *List) bool {
return list.compareListUtil(list.head, ll.head)
}
func (list *List) compareListUtil(head1 *Node, head2 *Node) bool {
if head1 == nil && head2 == nil {
return true
} else if (head1 == nil) || (head2 == nil) || (head1.value != head2.value) {
return false
} else {
return list.compareListUtil(head1.next, head2.next)
}
}
Analysis:
· List is compared recursively. Moreover, if we reach the end of the list and both the lists are null. Then both the lists
are equal and so return true.
· List is compared recursively. If either one of the list is empty or the value of corresponding nodes is unequal, then
this function will return false.
· Recursively calls compare list function for the next node of the current nodes.

# Find Length
Example 7.20: Find the length of given linked list.
func (list *List) FindLength() int {
curr := list.head
count := 0
for curr != nil {
count++
curr = curr.next
}
return count
}
Analysis: Length of linked list is found by traversing the list until we reach the end of list.
Nth Node from Beginning
Example 7.21: : Find Nth node from beginning
func (list *List) NthNodeFromBegining(index int) (int, bool) {
if index > list.Size() || index < 1 {
fmt.Println("TooFewNodes")
return 0, false
}
count := 0
curr := list.head
for curr != nil && count < index-1 {
count++
curr = curr.next
}
return curr.value, true
}
Analysis: Nth node can be found by traversing the list N-1 number of time and then return the node. If list does not
have N elements then the method return null.

# Nth Node from End
Example 7.22: Find Nth node from end
func (list *List) NthNodeFromEnd(index int) (int, bool) {
size := list.findLength()
if size != 0 && size < index {
fmt.Println("TooFewNodes")
return 0, false
}
startIndex := size - index + 1
return list.NthNodeFromBegining(startIndex)
}
Analysis: First, find the length of list, then nth node from end will be (length – nth +1) node from the beginning.
Example 7.23:*/
func (list *List) NthNodeFromEnd2(index int) (int, bool) {
count := 1
forward := list.head
curr := list.head
for forward != nil && count <= index {
count++
forward = forward.next
}
if forward == nil {
fmt.Println("TooFewNodes")
return 0, false
}
for forward != nil {
forward = forward.next
curr = curr.next
}
return curr.value, true
}

/*Analysis: Second approach is to use two references one is N steps / nodes ahead of the other when forward reference
reach the end of the list then the backward reference will point to the desired node.
Loop Detect
Find if there is a loop in a linked list. If there is a loop, then return 1 if not, then return 0.
There are many ways to find if there is a loop in a linked list:
Approach 1: User some map or hash-table
a) Traverse through the list.
b) If the current node is, not there in the Hash-Table then insert it into the Hash-Table.
c) If the current node is already in the Hashtable then we have a loop.
Approach 2: Slow reference and fast reference approach (SPFP)
Approach 3: Reverse list approach”
Slow reference and fast reference approach (SPFP)
We will use two references, one will move 2 steps at a time and another will move 1 step at time. If there is, a loop then
both will meet at a point.

Example 7.24:*/
func (list *List) LoopDetect() bool {
slowPtr := list.head
fastPtr := list.head
for fastPtr.next != nil && fastPtr.next.next != nil {
slowPtr = slowPtr.next
fastPtr = fastPtr.next.next
if slowPtr == fastPtr {
fmt.Println("loop found")
return true
}
}
fmt.Println("loop not found")
return false
}
/*Analysis:
· The list is traversed with two references, one is slow reference and another is fast reference. Slow reference always
moves one-step. Fast reference always moves two steps. If there is no loop, then control will come out of for loop.
So return false.
· If there is a loop, then there came a point in a loop where the fast reference will come and try to pass slow reference
and they will meet at a point. When this point arrives, we come to know that there is a loop in the list. So return
true.
# Reverse List Loop Detect
If there is a loop in a linked list, then reverse list function will give head of the original list as the head of the new list.
Example 7.25: Find if there is a loop in a linked list. Use reverse list approach.*/
func (list *List) ReverseListLoopDetect() bool {
tempHead := list.head
list.Reverse()
if tempHead == list.head {
list.Reverse()
fmt.Println("loop found")
return true
}
list.Reverse()
fmt.Println("loop not found")
return false
}
/*Analysis:
· Store reference of the head of list in a temp variable.
· Reverse the list
· Compare the reversed list head reference to the current list head reference.
· If the head of reversed list and the original list are same then reverse the list back and return true.
· If the head of the reversed list and the original list are not same, then reverse the list back and return false. Which
means there is no loop.
Note: Both SPFP and Reverse List approaches are linear in nature, but still in SPFP approach, we do not require to
modify the linked list so it is preferred.
Loop Type Detect
Find if there is a loop in a linked list. If there is no loop, then return 0, if there is loop return 1, if the list is circular then
2. # Use slow reference fast reference approach.
Example 7.26:*/
func (list *List) LoopTypeDetect() int {
slowPtr := list.head
fastPtr := list.head
for fastPtr.next != nil && fastPtr.next.next != nil {
if list.head == fastPtr.next || list.head == fastPtr.next.next {
fmt.Println("circular list loop found")
return 2
}
slowPtr = slowPtr.next
fastPtr = fastPtr.next.next
if slowPtr == fastPtr {
fmt.Println("loop found")
return 1
}
}
fmt.Println("loop not found")
return 0
}
/*Analysis: This program is same as the loop detect program only if it is a circular list than the fast reference reach the
slow reference at the head of the list this means that there is a loop at the beginning of the list.*/

# *Remove Loop
Example 7.27: Given there is a loop in linked list remove the loop.*/
func (list *List) RemoveLoop() {
loopPoint := list.LoopPointDetect()
if loopPoint == nil {
return
}
firstPtr := list.head
if loopPoint == list.head {
for firstPtr.next != list.head {
firstPtr = firstPtr.next
}
firstPtr.next = nil
return
}
secondPtr := loopPoint
for firstPtr.next != secondPtr.next {
firstPtr = firstPtr.next
secondPtr = secondPtr.next
}
secondPtr.next = nil
}
func (list *List) LoopPointDetect() *Node {
slowPtr := list.head
fastPtr := list.head
for fastPtr.next != nil && fastPtr.next.next != nil {
slowPtr = slowPtr.next
fastPtr = fastPtr.next.next
if slowPtr == fastPtr {
return slowPtr
}
}
return nil
}
/*Analysis:
· Loop through the list by two reference, one fast reference and one slow reference. Fast reference jumps two nodes
at a time and slow reference jump one node at a time. The point where these two reference intersect is a point in the
loop.
· If that intersection point is head of the list, this is a circular list case and you need to again traverse through the list
and make the node before head point to null.
· In the other case, you need to use two-reference variables one start from head and another start form the
intersection-point. They both will meet at the point of loop. (You can mathematically prove it ;) )*/

# *Find Intersection
Example 7.28: Given two linked list which meet at some point find that intersection point.*/
func (list *List) FindIntersection(head *Node, head2 *Node) *Node {
l1 := 0
l2 := 0
tempHead := head
tempHead2 := head2
for tempHead != nil {
l1++
tempHead = tempHead.next
}
for tempHead2 != nil {
l2++
tempHead2 = tempHead2.next
}
var diff int
if l1 < 12 {
temp := head
head = head2
head2 = temp
diff = l2 - l1
} else {
diff = l1 - l2
}
for ; diff > 0; diff-- {
head = head.next
}
for head != head2 {
head = head.next
head2 = head2.next
}
return head
}
// Analysis: Find length of both the lists. Find the difference of length of both the lists. Increment the longer list by diff
// steps, and then increment both the lists and get the intersection point.