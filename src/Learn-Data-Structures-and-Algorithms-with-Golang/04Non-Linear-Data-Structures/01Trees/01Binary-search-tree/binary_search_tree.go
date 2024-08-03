///main package has examples shown
// in Go Data Structures and algorithms book

package main

// importing fmt package
import (
	"fmt"
	"sync"
)

/*Binary search tree
A binary search tree is a data structure that allows for the quick lookup, addition, and
removal of elements. It stores the keys in a sorted order to enable a faster lookup. This data
structure was invented by P. F. Windley, A. D. Booth, A. J. T. Colin, and T. N. Hibbard. On
average, space usage for a binary search tree is of the order O(n), whereas the insert, search,
and delete operations are of the order O(log n). A binary search tree consists of nodes with
properties or attributes:
A key integer
A value integer
The leftNode and rightNode instances of TreeNode
They can be represented in the following code:*/

// TreeNode class
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

/*The BinarySearchTree class
In the following code snippet, the BinarySearchTree class consists of a rootNode that's
of the TreeNode type, and lock, which is of the sync.RWMutex type. The binary search tree
is traversed from rootNode by accessing the nodes to the left and right of rootNode:*/

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

/*The InsertElement method
The InsertElement method inserts the element with the given key and value in the binary
search tree. The tree's lock() instance is locked first and the unlock() method is deferred
before inserting the element. The InsertTreeNode method is invoked by passing
rootNode and the node to be created with the key and value, as shown here:*/
// InsertElement method
func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

/*The insertTreeNode method
The insertTreenode method inserts the new TreeNode in the binary search tree. In the
following code, the insertTreeNode method takes rootNode and newTreeNode, both of
the TreeNode type, as parameters. Note that newTreeNode is inserted at the right place in
the binary search tree by comparing the key values:*/
// insertTreeNode method
func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

// InOrderTraverseTree method
/*The inOrderTraverse method visits all nodes in order. The RLock() method on the
tree lock instance is called first. The RUnLock() method is deferred on the tree lock
instance before invoking the inOrderTraverseTree method, as presented in the
following code snippet:*/
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode, function)
}

// inOrderTraverseTree method
/*The inOrderTraverseTree method
The inOrderTraverseTree method traverses the left, the root, and the right tree. The
inOrderTraverseTree method takes treeNode of the TreeNode type and function as
parameters. The inOrderTraverseTree method is called on leftNode and rightNode
with function as a parameter. A function is passed with treeNode.value, as shown in
the following code snippet:*/

func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

// PreOrderTraverse method
/*The PreOrderTraverseTree method
The PreOrderTraverseTree method visits all the tree nodes with preorder traversing.
The tree lock instance is locked first and the Unlock method is deferred before
preOrderTraverseTree is called. In the following code snippet, the
preOrderTraverseTree method is passed with rootNode and function as parameters:
*/
/* Preorder traversing in a binary tree means visiting nodes in the following order:
1. Visit the root node.
2. Traverse the left subtree.
3. Traverse the right subtree.

### Example:
Given a binary tree:
```
    1
   / \
  2   3
 / \
4   5
```

Preorder traversal of this tree would result in the sequence: `1, 2, 4, 5, 3`.*/
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTraverseTree(tree.rootNode, function)
}

// preOrderTraverseTree method
/*The preOrderTraverseTree method is passed with treeNode of the TreeNode type and
function as parameters. The preOrderTraverseTree method is called by passing
leftNode and rightNode with function as parameters. The function is invoked with
treeNode.value, as shown here:*/
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

// PostOrderTraverseTree method
/*The PostOrderTraverseTree method
The PostOrderTraverseTree method traverses the nodes in a post order (left, right,
current node). In the following code snippet, the PostOrderTraverseTree method of the
BinarySearchTree class visits all nodes with post-order traversing. The
function method is passed as a parameter to the method. The tree.lock instance is
locked first and the Unlock method is deferred on the tree lock instance before calling the
postOrderTraverseTree method:*/

/*Postorder traversing in a binary tree means visiting nodes in the following order:
1. Traverse the left subtree.
2. Traverse the right subtree.
3. Visit the root node.

### Example:
Given a binary tree:
```
    1
   / \
  2   3
 / \
4   5
```

Postorder traversal of this tree would result in the sequence: `4, 5, 2, 3, 1`.*/

func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	postOrderTraverseTree(tree.rootNode, function)
}

// postOrderTraverseTree method
/*The postOrderTraverseTree method
The postOrderTraverseTree method is passed with treeNode of the TreeNode type
and function as parameters. The postOrderTraverseTree method is called by passing
leftNode and rightNode with function as parameters. In the following code snippet,
function is invoked with treeNode.value as a parameter:*/
func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, function)
		postOrderTraverseTree(treeNode.rightNode, function)
		function(treeNode.value)
	}
}

// MinNode method

/*MinNode finds the node with the minimum value in the binary search tree. In the following
code snippet, the RLock method of the tree lock instance is invoked first and the RUnlock
method on the tree lock instance is deferred. The MinNode method returns the element
with the lowest value by traversing from rootNode and checking whether the value of
leftNode is nil:*/

func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	var treeNode *TreeNode
	treeNode = tree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}

// MaxNode method
/*The MaxNode method
MaxNode finds the node with maximum property in the binary search tree. The RLock
method of the tree lock instance is called first and the RUnlock method on the tree lock
instance is deferred. The MaxNode method returns the element with the highest value after
traversing from rootNode and finding a rightNode with a nil value. This is shown in the
following code:*/
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	var treeNode *TreeNode
	treeNode = tree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

// SearchNode method
/*The SearchNode method
The SearchNode method searches the specified node in the binary search tree. First, the
RLock method of the tree lock instance is called. Then, the RUnlock method on the
tree lock instance is deferred. The SearchNode method of the BinarySearchTree class
invokes the searchNode method with the rootNode and the key integer value
as parameters, as shown here:*/
func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, key)
}

// searchNode method
/*The searchNode method
In the following code, the searchNode method takes treeNode, a pointer of the TreeNode
type, and a key integer value as parameters. The method returns true or false after
checking whether treeNode with the same value as key exists:*/
func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return true
}

// RemoveNode method
/*The RemoveNode method
The RemoveNode method of the BinarySearchTree class removes the element with key
that's passed in. The method takes the key integer value as the parameter. The Lock()
method is invoked on the tree's lock instance first. The Unlock() method of the tree lock
instance is deferred, and removeNode is called with rootNode and the key value as
parameters, as shown here:*/

/*Locks are used in binary tree interactions to ensure thread safety and prevent race conditions when multiple threads concurrently modify or traverse the tree. This ensures data integrity and consistency.*/
func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, key)
}

// removeNode method
/*The removeNode method
The removeNode method takes treeNode of the TreeNode type and a key integer value as
parameters. In the following code snippet, the method recursively searches the leftNode
instance of treeNode and the key value of rightNode if it matches the parameter value:*/
func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}
	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}
	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}

	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}
	var leftmostrightNode *TreeNode
	leftmostrightNode = treeNode.rightNode
	for {

		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)
	return treeNode
}

// String method
/*The String method turns the tree into a string format. At first, the Lock() method is
invoked on the tree lock instance. Then, the Unlock() method of the tree lock instance is
deferred. The String method prints a visual representation of tree:*/
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("************************************************")
	stringify(tree.rootNode, 0)
	fmt.Println("************************************************")
}

// stringify method
/*The stringify method
In the following code snippet, the stringify method takes a treeNode instance of the
TreeNode type and level (an integer) as parameters. The method recursively prints the
tree based on the level:*/

func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "***> "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}

// print method
func print(tree *BinarySearchTree) {
	if tree != nil {

		fmt.Println(" Value", tree.rootNode.value)
		fmt.Printf("Root Tree Node")
		printTreeNode(tree.rootNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

// printTreeNode method
func printTreeNode(treeNode *TreeNode) {
	if treeNode != nil {
		fmt.Println(" Value", treeNode.value)
		fmt.Printf("TreeNode Left")
		printTreeNode(treeNode.leftNode)
		fmt.Printf("TreeNode Right")
		printTreeNode(treeNode.rightNode)
	} else {
		fmt.Printf("Nil\n")
	}

}

// main method
/*In the following code, the main method creates the binary search tree and inserts the
elements 8, 3, 10, 1, and 6 into it. tree is printed by invoking the String method:*/
func main() {
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(8, 8)
	tree.InsertElement(3, 3)
	tree.InsertElement(10, 10)
	tree.InsertElement(1, 1)
	tree.InsertElement(6, 6)
	tree.String()

}
