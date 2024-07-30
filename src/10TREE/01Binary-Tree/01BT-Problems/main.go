package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}
type Tree struct {
	root *Node
}

// Problems in Binary Tree

/*Create a Complete binary tree
Create a binary tree given a list of values.
Solution: Since there is no order defined in a binary tree, so nodes can be inserted in any order so it can be a skewed
binary tree. But it is inefficient to do anything in a skewed binary tree so we will create a Complete binary tree. At each
node, the middle value stored in the array is assigned to node. The left of array is passed to the left child of the node to
create left sub-tree and the right portion of array is passed to right child of the node to crate right sub-tree.*/

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}
func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

func (t *Tree) PreOrder() {
	preOrder(t.root)
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		preOrder(node.left)
		preOrder(node.right)
	}
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t2 := LevelOrderBinaryTree(arr)
	t2.PreOrder()
}

/*Complexity Analysis: This is an efficient algorithm for creating a complete binary tree.
Time Complexity: O(n), Space Complexity: O(n)*/

/*Pre-Order Traversal
Traversal is a process of visiting each node of a tree. In Pre-Order Traversal parent is visited / traversed first, then left
child and then right child. Pre-Order traversal is a type of depth-first traversal.*/
/*Solution: Preorder traversal is done using recursion. At each node, first the value stored in it is printed and then
followed by the value of left child and right child. At each node its value is printed followed by calling printTree()
function to its left and right child to print left and right sub-tree.*/

/*Example 10.2:*/
func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
	fmt.Println()
}
func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	printPreOrder(n.left)
	printPreOrder(n.right)
}

/*Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Note: When there is an algorithm in which all nodes are traversed then complexity cannot be less than O(n). When
there is a large portion of the tree, which is not traversed, then complexity reduces.*/

//Post-Order Traversal

/*In Post-Order Traversal left child is visited / traversed first, then right child and then parent
Post-Order traversal is a type of depth-first traversal.*/

/*Solution: At each node, first the left child is traversed then right child and in the end, current node value is printed to
the screen.*/

/*Example 10.3:*/
func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
	fmt.Println()
}
func printPostOrder(n *Node) {
	if n == nil {
		return
	}
	printPostOrder(n.left)
	printPostOrder(n.right)
	fmt.Print(n.value)
}

//Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)

///In-Order Traversal

/*In-Order Traversal
In In-Order Traversal, left child is visited / traversed first, then the parent value is printed and last right child is
traversed.
In-Order traversal is a type of depth-first traversal. The output of In-Order traversal of BST is a sorted list.
Solution: In In-Order traversal first, the value of left child is traversed, then the value of node is printed to the screen
and then the value of right child is traversed.*/

func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
	fmt.Println()
}
func printInOrder(n *Node) {
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Print(n.value)
	printInOrder(n.right)
}

/*
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Note: Pre-Order, Post-Order, and In-Order traversal are for all binary trees. They can be used to traverse any kind of a
binary tree.*/

/*Level order traversal / Breadth First traversal
Write code to implement level order traversal of a tree. Such that nodes at depth k is printed before nodes at depth k+1.*/

// Solution: Level order traversal or Breadth First traversal of a tree is done using a queue. At the start, the root node
// reference is added to queue. The traversal of tree happens until its queue is empty. When we traverse the tree, we first
// remove an element from the queue, print the value stored in that node and then its left child and right child will be
// added to the queue.
// Example 10.5:

// func (t *Tree) PrintBredthFirst() {
// 	que := new(queue.Queue)
// 	var temp *Node
// 	if t.root != nil {
// 		que.Put(t.root)
// 	}
// 	for que.Empty() == false {
// 		temp2, _ := que.Get(1)
// 		temp = temp2[0].(*Node)
// 		fmt.Print(temp.value, " ")
// 		if temp.left != nil {
// 			que.Put(temp.left)
// 		}
// 		if temp.right != nil {
// 			que.Put(temp.right)
// 		}
// 	}
// 	fmt.Println()
// }

//Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)

/*Tree Depth
Solution: Depth is tree is calculated recursively by traversing left and right child of the root. At each level of traversal
depth of left child is calculated and depth of right child is calculated. The greater depth among the left and right child is
added by one (which is the depth of the current node) and this value is returned.*/
func (t *Tree) TreeDepth() int {
	return treeDepth(t.root)
}
func treeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lDepth := treeDepth(root.left)
	rDepth := treeDepth(root.right)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

/*Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)*/

//Others Operations
/*

Nth Pre-Order
Solution: We want to print the node, which will be at the nth index when we print the tree in PreOrder traversal.
Therefore, we keep a counter to keep track of the index. When the counter is equal to index, then we print the value and
return the Nth preorder index node.
Example 10.7:
func (t *Tree) NthPreOrder(index int) {
var counter int
nthPreOrder(t.root, index, &counter)
}
func nthPreOrder(node *Node, index int, counter *int) {
if node != nil {
(*counter)++
if *counter == index {
fmt.Println(node.value)
}
nthPreOrder(node.left, index, counter)
nthPreOrder(node.right, index, counter)
}
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Nth Post Order
Solution: We want to print the node that will be at the nth index when we print the tree in post order traversal.
Therefore, we keep a counter to keep track of the index, but at this time, we will increment the counter after left child
and right child traversal. When the counter is equal to index, then we print the value and return the nth post-order index
node.
Example 10.8
func (t *Tree) NthPostOrder(index int) {
var counter int
nthPostOrder(t.root, index, &counter)
}
func nthPostOrder(node *Node, index int, counter *int) {
if node != nil {
nthPostOrder(node.left, index, counter)
nthPostOrder(node.right, index, counter)
(*counter)++
if *counter == index {
fmt.Println(node.value)
}
}
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Nth In Order
Solution: We want to print the node that will be at the nth index when we print the tree in in-order traversal. Therefore,
we keep a counter to keep track of the index, but at this time, we will increment the counter after left child traversal but
before the right child traversal. When the counter is equal to index, then we print the value and return the nth in-order
index node.
Example 10.9:
func (t *Tree) NthInOrder(index int) {
var counter int
nthInOrder(t.root, index, &counter)
}
func nthInOrder(node *Node, index int, counter *int) {
if node != nil {
nthInOrder(node.left, index, counter)
*counter++
if *counter == index {
fmt.Println(node.value)
}
nthInOrder(node.right, index, counter)
}
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(1)
Copy Tree
Solution: Copy tree is done by copy nodes of the input tree at each level of the traversal of the tree. At each level of the
traversal of nodes of the tree, a new node is created and the value of the input tree node is copied to it. The left child
tree is copied recursively and then reference to new subtree is returned which will be assigned to the left child reference
of the current new node. Similarly for the right child node too. Finally, the tree is copied.
Example 10.10:
func (t *Tree) CopyTree() *Tree {
Tree2 := new(Tree)
Tree2.root = copyTree(t.root)
return Tree2
}
func copyTree(curr *Node) *Node {
var temp *Node
if curr != nil {
temp = new(Node)
temp.value = curr.value
temp.left = copyTree(curr.left)
temp.right = copyTree(curr.right)
return temp
}
return nil
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Copy Mirror Tree
Solution: Copy mirror image of the tree is done same as copy tree, but in place of left child pointing to the tree formed
by left child traversal of input tree. This time left child points to the tree formed by right child traversal of the input
tree. Similarly right child point to the tree formed by the traversal of the left child of the input tree.
Example 10.11:
func (t *Tree) CopyMirrorTree() *Tree {
tree := new(Tree)
tree.root = copyMirrorTree(t.root)
return tree
}
func copyMirrorTree(curr *Node) *Node {
var temp *Node
if curr != nil {
temp = new(Node)
temp.value = curr.value
temp.right = copyMirrorTree(curr.left)
temp.left = copyMirrorTree(curr.right)
return temp
}
return nil
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Number of Element
Solution: Number of nodes at the right child and the number of nodes at the left child is added by one and we get the
total number of nodes in any tree / sub-tree.
Example 10.12:
func (t *Tree) NumNodes() int {
return numNodes(t.root)
}
func numNodes(curr *Node) int {
if curr == nil {
return 0
}
return (1 + numNodes(curr.right) + numNodes(curr.left))
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Number of Leaf nodes
Solution: If we add the number of leaf node in the right child with the number of leaf nodes in the left child, we will
get the total number of leaf node in any tree or subtree.
Example 10.13:
func (t *Tree) NumLeafNodes() int {
return numLeafNodes(t.root)
}
func numLeafNodes(curr *Node) int {
if curr == nil {
return 0
}
if curr.left == nil && curr.right == nil {
return 1
}
return (numLeafNodes(curr.right) + numLeafNodes(curr.left))
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Identical
Solution: Two trees have identical values if at each level the value is equal.
Example 10.14:
func (t *Tree) IsEqual(t2 *Tree) bool {
return isEqual(t.root, t2.root)
}
func isEqual(node1 *Node, node2 *Node) bool {
if node1 == nil && node2 == nil {
return true
} else if node1 == nil || node2 == nil {
return false
} else {
return ((node1.value == node2.value) &&
isEqual(node1.left, node2.left) &&
isEqual(node1.right, node2.right))
}
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Free Tree
Solution: You just need to make the root of the tree point to nil. The system will do garbage collection and recover the
memory assigned to the tree. Effectively you had done a single act and because of this action, the time complexity is
constant.
Example 10.15:
func (t *Tree) Free() {
t.root = nil
}
Complexity Analysis: Time Complexity: O(1), Space Complexity: O(1)
Print all the paths
Print all the paths from the roots to the leaf
Solution: Whenever we traverse a node, we add that node to the list. When we reach a leaf, we print the whole list.
When we return from a function, then we remove the element that was added to the list when we entered this function.
Example 10.16:
func (t *Tree) PrintAllPath() {
stk := new(Stack)
printAllPath(t.root, stk)
}
func printAllPath(curr *Node, stk *Stack) {
if curr == nil {
return
}
stk.Push(curr.value)
if curr.left == nil && curr.right == nil {
stk.Print()
stk.Pop()
return
}
printAllPath(curr.right, stk)
printAllPath(curr.left, stk)
stk.Pop()
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Least Common Ancestor
Solution: We recursively traverse the nodes of a binary tree. We find any one of the input node for which we are
searching a common ancestor then we return that node. When we get both the left and right as some valid reference
location other than null, we will return that node as the common ancestor.
Example 10.17:
func (t *Tree) LCA(first int, second int) (int, bool) {
ans := LCAUtil(t.root, first, second)
if ans != nil {
return ans.value, true
}
fmt.Println("NotFoundError")
return 0, false
}
func LCAUtil(curr *Node, first int, second int) *Node {
var left, right *Node
if curr == nil {
return nil
}
if curr.value == first || curr.value == second {
return curr
}
left = LCAUtil(curr.left, first, second)
right = LCAUtil(curr.right, first, second)
if left != nil && right != nil {
return curr
} else if left != nil {
return left
} else {
return right
}
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Find Max in Binary Tree
Solution: We recursively traverse the nodes of a binary tree. We will find the maximum value in the left and right
subtree of any node then will compare the value with the value of the current node and finally return the largest of the
three values.
Example 10.18:
func (t *Tree) FindMaxBT() int {
return findMaxBT(t.root)
}
func findMaxBT(curr *Node) int {
if curr == nil {
return math.MinInt32
}
max := curr.value
left := findMaxBT(curr.left)
right := findMaxBT(curr.right)
if left > max {
max = left
}
if right > max {
max = right
}
return max
}
Search value in a Binary Tree
Solution: To find if some value is there in a binary tree or not is done using exhaustive search of the binary tree. First,
the value of current node is compared with the value, which we are looking for. Then it is compared recursively inside
the left child and right child.
Example 10.19:
func SearchBT(root *Node, value int) bool {
var left, right bool
if root == nil || root.value == value {
return false
}
left = SearchBT(root.left, value)
if left {
return true
}
right = SearchBT(root.right, value)
if right {
return true
}
return false
}
Maximum Depth in a Binary Tree
Solution: To find the maximum depth of a binary tree we need to find the depth of the left tree and depth of right tree
then we need to store the value and increment it by one so that we get depth of the given node.
Example 10.20:
func (t *Tree) TreeDepth() int {
return treeDepth(t.root)
}
func treeDepth(root *Node) int {
if root == nil {
return 0
}
lDepth := treeDepth(root.left)
rDepth := treeDepth(root.right)
if lDepth > rDepth {
return lDepth + 1
}
return rDepth + 1
}
Number of Full Nodes in a BT
Solution: A full node is a node that have both left and right child. We will recursively traverse the whole tree and will
increase the count of full node as we find them.
Example 10.21:
func (t *Tree) NumFullNodesBT() int {
return numFullNodesBT(t.root)
}
func numFullNodesBT(curr *Node) int {
var count int
if curr == nil {
return 0
}
count = numFullNodesBT(curr.right) + numFullNodesBT(curr.left)
if curr.right != nil && curr.left != nil {
count++
}
return count
}
Maximum Length Path in a BT/ Diameter of BT
Solution: To find the diameter of BT we need to find the depth of left child and right child then will add these two
values and increment it by one so that we will get the maximum length path (diameter candidate) which contains the
current node. Then we will find max length path in the left child sub-tree. Will also find the max length path in the right
child sub-tree. Finally, we will compare the three values and return the maximum value out of these this will be the
diameter of the Binary tree.
Example 10.22:
func (t *Tree) MaxLengthPathBT() int {
return maxLengthPathBT(t.root)
}
func maxLengthPathBT(curr *Node) int {
var max, leftPath, rightPath, leftMax, rightMax int
if curr == nil {
return 0
}
leftPath = treeDepth(curr.left)
rightPath = treeDepth(curr.right)
max = leftPath + rightPath + 1
leftMax = maxLengthPathBT(curr.left)
rightMax = maxLengthPathBT(curr.right)
if leftMax > max {
max = leftMax
}
if rightMax > max {
max = rightMax
}
return max
}
Sum of All nodes in a BT
Solution: We will find the sum of all the nodes recursively. sumAllBT() will return the sum of all the node of left and
right subtree then will add the value of current node and will return the final sum.
Example 10.23:
func (t *Tree) SumAllBT() int {
return sumAllBT(t.root)
}
func sumAllBT(curr *Node) int {
var sum, leftSum, rightSum int
if curr == nil {
return 0
}
rightSum = sumAllBT(curr.right)
leftSum = sumAllBT(curr.left)
sum = rightSum + leftSum + curr.value
return sum
}




*/
