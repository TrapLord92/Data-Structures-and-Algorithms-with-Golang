package main

/*Binary Search Tree (BST)
A binary search tree (BST) is a binary tree on which nodes are ordered in the following way:
· The key in the left subtree is less than the key in its parent node.
· The key in the right subtree is greater the key in its parent node.
· No duplicate key allowed.
Note: there can be two separate key and value fields in the tree node. But for simplicity, we are considering value as the
key. All problems in the binary search tree are solved using this supposition that the value in the node is key for the
tree.
Note: Since binary search tree is a binary tree to all the above algorithm of a binary tree are applicable to a binary
search tree.*/

type Node struct {
	value       int
	left, right *Node
}
type Tree struct {
	root *Node
}

/*Problems in Binary Search Tree (BST)
All binary tree algorithms are valid for binary search tree too.
Create a binary search tree from sorted list
Create a binary tree given list of values in a list in sorted order. Since the elements in the list are in sorted order and we
want to create a binary search tree in which left subtree nodes are having values less than the current node and right
subtree nodes have value greater than the value of the current node.
Solution: We have to find the middle node to create a current node and send the rest of the list to construct left and
right subtree.*/

// func CreateBinaryTree(arr []int) *Tree {
// 	t := new(Tree)
// 	size := len(arr)
// 	t.root = createBinaryTreeUtil(arr, 0, size-1)
// 	return t
// 	}
// 	func createBinaryTreeUtil(arr []int, start int, end int) *Node {
// 	if start > end {
// 	return nil
// 	}
// 	mid := (start + end) / 2
// 	curr := new(Node)
// 	curr.value = arr[mid]
// 	curr.left = createBinaryTreeUtil(arr, start, mid-1)
// 	curr.right = createBinaryTreeUtil(arr, mid+1, end)
// 	return curr
// 	}
// 	func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	t := CreateBinaryTree(arr)
// 	}

/*Insertion
Nodes with key 6,4,2,5,1,3,8,7,9,10 are inserted in a tree. Below is step by step tree after inserting nodes in the order.

Solution: Smaller values will be added to the left child sub-tree of a node and greater value will be added to the right
child sub-tree of the current node.
Example 10.25:
func (t *Tree) Add(value int) {
t.root = addUtil(t.root, value)
}
func addUtil(n *Node, value int) *Node {
if n == nil {
n = new(Node)
n.value = value
return n
}
if value < n.value {
n.left = addUtil(n.left, value)
} else {
n.right = addUtil(n.right, value)
}
return n
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
Find Node
Solution: The value greater than the current node value will be in the right child sub-tree and the value smaller than the
current node is in the left child sub-tree. We can find a value by traversing the left or right subtree iteratively.
Example 10.26: Find the node with the value given.
func (t *Tree) Find(value int) bool {
var curr *Node = t.root
for curr != nil {
if curr.value == value {
return true
} else if curr.value > value {
curr = curr.left
} else {
curr = curr.right
}
}
return false
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(1)
Find Min
Find the node with the minimum value.
Solution: left most child of the tree will be the node with the minimum value.
Example 10.27:
func (t *Tree) FindMin() (int, bool) {
var node *Node = t.root
if node == nil {
fmt.Println("EmptyTreeError")
return 0, false
}
for node.left != nil {
node = node.left
}
return node.value, true
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(1)
Find Max
Find the node in the tree with the maximum value.
Solution: Right most node of the tree will be the node with the maximum value.
Example 10.28:
func (t *Tree) FindMax() (int, bool) {
var node *Node = t.root
if node == nil {
fmt.Println("EmptyTreeError")
return 0, false
}
for node.right != nil {
node = node.right
}
return node.value, true
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(1)
Is tree a BST
Approach 1: At each node we check, max value of left subtree is smaller than the value of current node and min value
of right subtree is greater than the current node.
Example 10.29:
func IsBST3(root *Node) bool {
if root == nil {
return true
}
if root.left != nil && FindMax(root.left).value > root.value {
return false
}
if root.right != nil && FindMin(root.right).value <= root.value {
return false
}
return (IsBST3(root.left) && IsBST3(root.right))
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n)
The above solution is correct but it is not efficient as same tree nodes are traversed many times.
Approach 2: A better solution will be the one in which we will look into each node only once. This is done by
narrowing the range. We will be using an isBSTUtil() function which take the max and min range of the values of the
nodes. The initial value of min and max will be INT_MIN and INT_MAX.
Example 10.30:
func (t *Tree) IsBST() bool {
return IsBST(t.root, math.MinInt32, math.MaxInt32)
}
func IsBST(curr *Node, min int, max int) bool {
if curr == nil {
return true
}
if curr.value < min || curr.value > max {
return false
}
return IsBST(curr.left, min, curr.value) && IsBST(curr.right, curr.value, max)
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n) for stack
Approach 3: Above method is correct and efficient but there is an easy method to do the same. We can do in-order
traversal of nodes and see if we are getting a strictly increasing sequence
Example 10.31:
func (t *Tree) IsBST2() bool {
var c int
return IsBST2(t.root, &c)
}
func IsBST2(root *Node, count *int) bool {
var ret bool
if root != nil {
ret = IsBST2(root.left, count)
if !ret {
return false
}
if *count > root.value {
return false
}
*count = root.value
ret = IsBST2(root.right, count)
if !ret {
return false
}
}
return true
}
Complexity Analysis: Time Complexity: O(n), Space Complexity: O(n) for stack
Delete Node
Description: Remove the node x from the binary search tree, making the necessary, reorganize nodes of binary search
tree to maintain its properties.
There are three cases in delete node, let us call the node that need to be deleted as x.
Case 1: node x has no children. Just delete it (i.e. Change parent node so that it does not point to x)
Case 2: node x has one child. Splice out x by linking x’s parent to x’s child
Case 3: node x has two children. Splice out the x’s successor and replace x with x’s successor
When the node to be deleted have no children
This is a trivial case, in which we directly delete the node and return null.
When the node to be deleted have only one child.
In this case, we save the child in a temp variable, then delete current node, and finally return the child.
We want to remove node with value 9. The node has only
one child.
Right child of the parent of node with value 9 that is node
with value 8 will point to the node with value 10.
Finally, node with value 9 is removed from the tree.
When the node to be deleted has two children.
We want to delete node with value 6. Which have two
children.
We had found minimum value node of the right child of
node with value 6.
Minimum value node value is copied to the node with
value 6.
Delete node with minimum value 7 is called over the
right child tree of the node.
Finally the tree with both the children is created.
Example 10.32:
func (t *Tree) DeleteNode(value int) {
t.root = DeleteNode(t.root, value)
}
func DeleteNode(node *Node, value int) *Node {
var temp *Node = nil
if node != nil {
if node.value == value {
if node.left == nil && node.right == nil {
return nil
}
if node.left == nil {
temp = node.right
return temp
}
if node.right == nil {
temp = node.left
return temp
}
maxNode := FindMax(node.left)
maxValue := maxNode.value
node.value = maxValue
node.left = DeleteNode(node.left, maxValue)
} else {
if node.value > value {
node.left = DeleteNode(node.left, value)
} else {
node.right = DeleteNode(node.right, value)
}
}
}
return node
}
Analysis: Time Complexity: O(n), Space Complexity: O(n)
Least Common Ancestor
In a tree T. The least common ancestor between two nodes n1 and n2 is defined as the lowest node in T that has both n1
and n2 as descendants.
Example 10.33:
func (t *Tree) LcaBST(first int, second int) (int, bool) {
return LcaBST(t.root, first, second)
}
func LcaBST(curr *Node, first int, second int) (int, bool) {
if curr == nil {
fmt.Println("NotFound")
return 0, false
}
if curr.value > first && curr.value > second {
return LcaBST(curr.left, first, second)
}
if curr.value < first && curr.value < second {
return LcaBST(curr.right, first, second)
}
return curr.value, true
}
Trim the Tree nodes which are Outside Range
Given a range as min, max. We need to delete all the nodes of the tree that are out of this range.
Solution: Traverse the tree and each node that is having value outside the range will delete itself. All the deletion will
happen from inside out so we do not have to care about the children of a node as if they are out of range then they
already had deleted themselves.
Example 10.34:
func (t *Tree) TrimOutsidedataRange(min int, max int) {
t.root = trimOutsidedataRange(t.root, min, max)
}
func trimOutsidedataRange(curr *Node, min int, max int) *Node {
if curr == nil {
return nil
}
curr.left = trimOutsidedataRange(curr.left, min, max)
curr.right = trimOutsidedataRange(curr.right, min, max)
if curr.value < min {
return curr.right
}
if curr.value > max {
return curr.left
}
return curr
}
Print Tree nodes which are in Range
Print only those nodes of the tree whose value is in the range given.
Solution: Just normal inorder traversal and at the time of printing we will check if the value is inside the range
provided.
Example 10.35:
func (t *Tree) PrintDataInRange(min int, max int) {
printDataInRange(t.root, min, max)
}
func printDataInRange(root *Node, min int, max int) {
if root == nil {
return
}
printDataInRange(root.left, min, max)
if root.value >= min && root.value <= max {
fmt.Print(root.value, " ")
}
printDataInRange(root.right, min, max)
}
Find Ceil and Floor value inside BST given key
Given a tree and a value we need to find the ceil value of node in tree which is smaller than the given value and need to
find the floor value of node in tree which is bigger. Our aim is to find ceil and floor value as close as possible then the
given value.
Example 10.36:
func (t *Tree) FloorBST(val int) int {
curr := t.root
floor := math.MaxInt32
for curr != nil {
if curr.value == val {
floor = curr.value
break
} else if curr.value > val {
curr = curr.left
} else {
floor = curr.value
curr = curr.right
}
}
return floor
}
func (t *Tree) CeilBST(val int) int {
curr := t.root
ceil := math.MinInt32
for curr != nil {
if curr.value == val {
ceil = curr.value
break
} else if curr.value > val {
ceil = curr.value
curr = curr.left
} else {
curr = curr.right
}
}
return ceil
}*/

/*More*/

/*Segment Tree
Segment tree is a binary tree that is used to make multiple range queries and range update in an array.
Examples of problems for which Segment Tree can be used are:
1. Finding the sum of all the elements of an array in a given range of index
2. Finding the maximum value of the array in a given range of index.
3. Finding the minimum value of the array in a given range of index (also known as Range Minimum Query problem)
Properties of Segment Tree:
1. Segment tree is a binary tree.
2. Each node in a segment tree represent an interval in the array.
3. The root of tree represent the whole array.
4. Each leaf node of represent a single element.
Note:- Segment tree solve problems, which can be solve in linear time by just scanning and updating the elements of
array. The only benefit we are getting from segment tree is that it does update and query operation in logarithmic time
that is more efficient than the linear approach.
Let us consider a simple problem:
Given an array of N numbers. You need to perform the following operations:
1. Update any element in the array
2. Find the maximum in any given range (i, j)
Solution 1:
Updating: Just update the element in the array, a[i] =x. Finding maximum in the range (i, j), by traversing through the
elements of the array in that range.
Time Complexity of Update is O(1) and of Finding is O(n)
Solution 2: The above solution is good. However, can we improve performance of Finding?
The answer is yes. In fact we can do both the operation in O(log n) where n is the size of the array. This we can do
using a segment tree.
Let us suppose we are given an input array A = {1, 8, 2, 7, 3, 6, 4, 5}. Moreover, the below diagram will represent the
segment tree formed corresponding to the input array A.

AVL Trees
An AVL tree is a binary search tree (BST) with an additional property that the subtrees of every node differ in height
by at most one. An AVL tree is a height balanced BST.
AVL tree is a balanced binary search tree. Adding or removing a node form AVL tree may make the AVL tree
unbalanced. Such violations of AVL balance property is corrected by one or two simple steps called rotations. Let us
assume that insertion of a new node had converted a previously balanced AVL tree into an unbalanced tree. Since the
tree is previously balanced and a single new node is added to it, maximum the unbalance difference in height will be 2.
Therefore, in the bottom most unbalanced node there are only four cases:
Case 1: The new node is left child of the left child of the current node.
Case 2: The new node is right child of the left child of the current node.
Case 3: The new node is left child of the right child of the current node.
Case 4: The new node is right child of the right child of the current node.
Case 1 can be re-balanced using a single Right Rotation.
Case 4 is symmetrical to Case 1: can be re-balanced using a single Left Rotation
Case 2 can be re-balanced using a double rotation. First, rotate left than rotation right.
Case 3 is symmetrical to Case 2: can be re-balanced using a double rotation. First, rotate right than rotation left.
Time Complexity of Insertion: To search the location where a new node need to be added in done in O(log(n)). Then on
the way back, we look for the AVL balanced property and fixes them with rotation. Since the rotation at each node is
done in constant time, the total amount of word is proportional to the length of the path. Therefore, the final time
complexity of insertion is O(log(n)).

Red-Black Tree
The red-black tree contains its data, left and right children like any other binary tree. In addition to this its node also
contains an extra bit of information which represents colour which can either red or black. Red-Black tree also contains
a specialized class of nodes called NULL nodes. NULL nodes are pseudo nodes that exists at the leaf of the tree. All
internal nodes has data associated with them.
Red-Black tree have the following properties:
1. Root of tree is black.
2. Every leaf node (NULL node) is black.
3. If a node is red then both of its children are black.
4. Every path from a node to a descendant leaf contains the same number of black nodes.
The first three properties are self-explanatory. The forth property states that, from any node in the tree to any leaf
(NULL), the number of black nodes must be the same.
In the above figure, from the root node to the leaf node (NULL) the number of black node is always three nodes.
Like the AVL tree, red-black trees are also self-balancing binary search tree. Whereas the balance property of an AVL
tree was a direct relationship between the heights of left and right subtrees of each node. In red-black trees, the
balancing property is governed by the four rules mentioned above. Adding or removing a node form red-black tree may
violate the properties of a red-black tree. The red-black properties are restored through recolouring and rotation. Insert,
delete, and search operation time complexity is O(log(n))
Splay tree
A splay tree is a self-adjusting binary search tree with the additional property that recently accessed elements are quick
to access again. It performs basic operations such as insertion, look-up and removal in O(log n) amortized time.
Elements of the tree are rearranged so that the recently accessed element is placed at the top of the tree. When an
element is searched then we use standard BST search and then use rotation to bring the element to the top.
Splay tree Average Case Worst Case
Space complexity O(n) O(n)
Time complexity search O(log(n)) Amortized O(log(n))
Time complexity insert O(log(n)) Amortized O(log(n))
Time complexity delete O(log(n)) Amortized O(log(n))
Unlike the AVL tree, the splay tree is not guaranteed to be height balanced. What is guaranteed is that the total cost of
the entire series of accesses will be cheap.
B-Tree
As we had already seen various types of binary tree for searching, insertion and deletion of data in the main memory.
However, these data structures are not appropriate for huge data that cannot fit into main memory, the data that is
stored in the disk.
A B-tree is a self-balancing search tree that allows searches, insertions, and deletions in logarithmic time. The B-tree is
a tree in which a node can have multiple children. Unlike self-balancing binary search trees, the B-tree is optimized for
systems that read and write entire blocks (page) of data. The read - write operation from disk is very slow as compared
with the main memory. The main purpose of B-Tree is to reduce the number of disk access. The node in a B-Tree is
having a huge number of references to the children nodes. Thereby reducing the size of the tree. While accessing data
from disk, it make sense to read an entire block of data and store into a node of tree. B-Tree nodes are designed such
that entire block of data (page) fits into it. It is commonly used in databases and filesystems.
B-Tree of minimum degree d has the following properties:
1. All the leaf nodes must be at same level.
2. All nodes except root must have at least (d-1) keys and maximum of (2d-1) keys. Root may contain minimum 1 key.
3. If the root node is a non-leaf node, then it must have at least 2 children.
4. A non-leaf node with N keys must have (N+1) number of children.
5. All the key values within a node must be in Ascending Order.
6. All keys of a node are sorted in ascending order. The child between two keys, K1 and K2 contains all keys in range
from K1 and K2.
B-Tree Average Case Worst Case
Space complexity O(n) O(n)
Time complexity search O(log(n)) O(log(n))
Time complexity insert O(log(n)) O(log(n))
Time complexity delete O(log(n)) O(log(n))
Below is the steps of creation of B-Tree by adding value from 1 to 7.
1
Insert 1 to the tree.
Stable
2
Insert 2 to the tree.
Stable
3
Insert 3 to the tree.
Intermediate
4
New node is created and data is distributed.
Stable
5
Insert 4 to the tree.
Stable
6
Insert 5 to the tree.
Intermediate
7
New node is created and data is distributed.
Stable
8
Insert 6 to the tree.
Stable
9
Insert 7 to the tree. New node is created and
data is distributed.
Intermediate
10
After rearranging the intermediate node is, also
have more than maximum number of keys.
Intermediate
11
New node is created and data is distributed. The
height of the tree is increased.
Stable
Note:- 2-3 tree is a B-tree of degree three.
B+ Tree
B+ Tree is a variant of B-Tree. The B+ Tree store records only at the leaf nodes. The internal nodes store keys. These
keys are used in insertion, deletion and search. The rules of splitting and merging of nodes is same as B-Tree.
b-order B+ tree Average Case Worst Case
Space complexity O(n) O(n)
Time complexity search O(logb(n)) O(logb(n))
Time complexity insert O(logb(n)) O(logb(n))
Time complexity delete O(logb(n)) O(logb(n))
Below is the B+ Tree created by adding value from 1 to 5.
1.
Value 1 is inserted to leaf node.
2.
Value 2 is inserted to leaf node.
3.
Value 3 is inserted to leaf node. Content of the leaf node passed
the maximum number of elements. Therefore, node is split and
intermediate / key node is created.
4.
Value 4 is further inserted to the leaf node. Which further split the
leaf node.
5.
Value 5 is added to the leaf node the number of nodes in the leaf
passed the maximum number of nodes that it can contain so it is
divided into 2. One more key is added to the intermediate node,
which make it also, passed maximum number of nodes it can
contain, and finally divided and a new node is created.
B* Tree
The B* tree is identical to the B+ tree, except for the rules for splitting and merging of nodes. Instead of splitting a node
into two halves when it overflows, the B* tree node tries to gives some of its records to its neighbouring sibling. If the
sibling is also full, then a new node is created and records are distributed into three.
Exercise
1. Construct a tree given its in-order and pre-order traversal strings.
o inorder: 1 2 3 4 5 6 7 8 9 10
o pre-order: 6 4 2 1 3 5 8 7 9 10
2. Construct a tree given its in-order and post-order traversal strings.
o inorder: 1 2 3 4 5 6 7 8 9 10
o post-order: 1 3 2 5 4 7 10 9 8 6
3. Write a delete node function in Binary tree.
4. Write a function print depth first in a binary tree without using system stack
Hint: you may want to keep another element to tree node like visited flag.
5. The worst-case runtime Complexity of building a BST with n nodes
o O(n2)
o O(n * log n)
o O(n)
o O(logn)
6. The worst-case runtime Complexity of insertion into a BST with n nodes is
o O(n2)
o O(n * log n)
o O(n)
o O(logn)
7. The worst-case runtime Complexity of a search of a value in a BST with n nodes.
o O(n2)
o O(n * log n)
o O(n)
o O(logn)
8. Which of the following traversals always gives the sorted sequence of the elements in a BST?
o Preorder
o Ignored
o Postorder
o Undefined
9. The height of a Binary Search Tree with n nodes in the worst case?
o O(n * log n)
o O(n)
o O(logn)
o O(1)
o
10. Check whether a given Binary Tree is Complete or not
o In a complete binary tree, every level except the last one is completely filled. All nodes in the left are filled
first, then the right one.
11. Check whether a given Binary Tree is Full/ Strictly binary tree or not. The full binary tree is a binary tree in which
each node has zero or two children.
12. Check whether a given Binary Tree is a Perfect binary tree or not. The perfect binary tree- is a type of full binary
trees in which each non-leaf node has exactly two child nodes.
13. Check whether a given Binary Tree is Height-balanced Binary Tree or not. A height-balanced binary tree is a
binary tree such that the left & right subtrees for any given node differ in height by no more than one
14. Isomorphic: two trees are isomorphic if they have the same shape, it does not matter what the value is. Write a
program to find if two given tree are isomorphic or not.
15. Try to optimize the above solution to give a DFS traversal without using recursion use some stack or queue.
16. This is an open exercise for the readers. Every algorithm that is solved using recursion (system stack) can also be
solved using user defined or library defined stack. So try to figure out what all algorithms that are using recursion
and try to figure out how you will do this same issue using user layer stack.
17. In a binary tree, print the nodes in zigzag order. In the first level, nodes are printed in the left to right order. In the
second level, nodes are printed in right to left and in the third level again in the order left to right.
Hint: Use two stacks. Pop from first stack and push into another stack. Swap the stacks alternatively.
18. Find nth smallest element in a binary search tree.
Hint: Nth inorder in a binary tree.
19. Find the floor value of key that is inside a BST.
20. Find the Ceil value of key, which is inside a BST.*/
