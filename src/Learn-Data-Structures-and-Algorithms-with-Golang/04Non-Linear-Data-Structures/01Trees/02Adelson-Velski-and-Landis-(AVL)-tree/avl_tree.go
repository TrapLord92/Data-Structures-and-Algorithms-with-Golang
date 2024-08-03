// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"encoding/json"
	"fmt"
)

/*Adelson, Velski, and Landis pioneered the AVL tree data structure and hence it is named
after them. It consists of height adjusting binary search trees. The balance factor is obtained
by finding the difference between the heights of the left and right sub-trees. Balancing is
done using rotation techniques. If the balance factor is greater than one, rotation shifts the
nodes to the opposite of the left or right sub-trees. The search, addition, and deletion
operations are processed in the order of O(log n).*/

// KeyValue type
/*The KeyValue interface
The KeyValue interface has the LessThan and EqualTo methods. The LessThan and
EqualTo methods take KeyValue as a parameter and return a Boolean value after checking
the less than or equal to condition. This is shown in the following code:*/
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}

// TreeNode class
/*The TreeNode class
The TreeNode class has KeyValue, BalanceValue, and LinkedNodes as properties. The
AVL tree is created as a tree of nodes of the TreeNode type, as shown here:*/
type TreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  [2]*TreeNode
}

// opposite method
/*The opposite method
The opposite method takes a node value and returns the opposite node's value. In the
following code snippet, the opposite method takes the nodeValue integer as a parameter
and returns the opposite node's value:*/
func opposite(nodeValue int) int {
	return 1 - nodeValue
}

// single rotation method
/*The singleRotation method
The singleRotation method rotates the node opposite to the specified sub-tree. As
shown in the following snippet, the singleRotation function rotates the node opposite
the left or right sub-tree. The method takes the pointer to rootNode and a nodeValue
integer as parameters and returns a TreeNode pointer:*/
func singleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {

	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// double rotation method
/*The doubleRotation method
Here, the doubleRotation method rotates the node twice. The method returns a
TreeNode pointer, taking parameters such as rootNode, which is a treeNode pointer, and
nodeValue, which is an integer. This is shown in the following code:*/
func doubleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {

	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue]

	rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue] = saveNode.LinkedNodes[opposite(nodeValue)]
	saveNode.LinkedNodes[opposite(nodeValue)] = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode

	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// adjust balance method
/*The adjustBalance method
The adjustBalance method adjusts the balance of the tree. In the following code snippet,
the adjustBalance method does a double rotation given the balance factor, rootNode,
and nodeValue. The adjustBalance method takes rootNode, which is an instance of the
TreeNode type, nodeValue, and balanceValue (which are both integers) as parameters:*/
func adjustBalance(rootNode *TreeNode, nodeValue int, balanceValue int) {

	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var oppNode *TreeNode
	oppNode = node.LinkedNodes[opposite(nodeValue)]
	switch oppNode.BalanceValue {
	case 0:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

// balanceTree method
/*The BalanceTree method
The BalanceTree method changes the balance factor by a single or double rotation. The
method takes rootNode (a TreeNode pointer) and nodeValue (an integer) as parameters.
The BalanceTree method returns a TreeNode pointer, as shown here:*/

func BalanceTree(rootNode *TreeNode, nodeValue int) *TreeNode {
	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var balance int
	balance = 2*nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, opposite(nodeValue))
	}
	adjustBalance(rootNode, nodeValue, balance)
	return doubleRotation(rootNode, opposite(nodeValue))
}

// insertRNode method
/*The insertRNode method
The insertRNode method inserts the node and balances the tree. This
method inserts rootNode with the KeyValue key, as presented in the following code
snippet. The method takes rootNode, which is a TreeNode pointer, and the key as an
integer as parameters. The method returns a TreeNode pointer and a Boolean value if the
rootNode is inserted:*/
func insertRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return &TreeNode{KeyValue: key}, false
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = insertRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir - 1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode, true
	case 1, -1:
		return rootNode, false
	}
	return BalanceTree(rootNode, dir), true
}

// InsertNode method
/*The InsertNode method
The InsertNode method inserts a node into the AVL tree. This method takes treeNode,
which is a double TreeNode pointer, and the key value as parameters:*/
func InsertNode(treeNode **TreeNode, key KeyValue) {
	*treeNode, _ = insertRNode(*treeNode, key)
}

// RemoveNode method
/*The RemoveNode method
In the following code, the RemoveNode method removes the element from the AVL tree by
invoking the removeRNode method. The method takes treeNode, which is a double
TreeNode pointer, and KeyValue as parameters:*/
func RemoveNode(treeNode **TreeNode, key KeyValue) {
	*treeNode, _ = removeRNode(*treeNode, key)
}

// removeBalance method
/*The removeBalance method
The removeBalance method removes the balance factor in a tree. This method adjusts the
balance factor after removing the node and returns a treeNode pointer and a Boolean if the
balance is removed. The method takes rootNode (an instance of TreeNode) and
nodeValue (an integer) as parameters. This is shown in the following code:*/
func removeBalance(rootNode *TreeNode, nodeValue int) (*TreeNode, bool) {
	var node *TreeNode
	node = rootNode.LinkedNodes[opposite(nodeValue)]
	var balance int
	balance = 2*nodeValue - 1
	switch node.BalanceValue {
	case -balance:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, nodeValue), false
	case balance:
		adjustBalance(rootNode, opposite(nodeValue), -balance)
		return doubleRotation(rootNode, nodeValue), false
	}
	rootNode.BalanceValue = -balance
	node.BalanceValue = balance
	return singleRotation(rootNode, nodeValue), true
}

// removeRNode method
/*The removeRNode method
The removeRNode method removes the node from the tree and balances the tree.
This method takes rootNode, which is a TreeNode pointer, and the key value. This
method returns a TreeNode pointer and Boolean value if RNode is removed, as shown in
the following code snippet:*/

func removeRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return nil, false
	}
	if rootNode.KeyValue.EqualTo(key) {
		switch {
		case rootNode.LinkedNodes[0] == nil:
			return rootNode.LinkedNodes[1], false
		case rootNode.LinkedNodes[1] == nil:
			return rootNode.LinkedNodes[0], false
		}
		var heirNode *TreeNode
		heirNode = rootNode.LinkedNodes[0]
		for heirNode.LinkedNodes[1] != nil {
			heirNode = heirNode.LinkedNodes[1]
		}
		rootNode.KeyValue = heirNode.KeyValue
		key = heirNode.KeyValue
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = removeRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (1 - 2*dir)
	switch rootNode.BalanceValue {
	case 1, -1:
		return rootNode, true
	case 0:
		return rootNode, false
	}
	return removeBalance(rootNode, dir)
}

type integerKey int

func (k integerKey) LessThan(k1 KeyValue) bool { return k < k1.(integerKey) }
func (k integerKey) EqualTo(k1 KeyValue) bool  { return k == k1.(integerKey) }

// main method
/*The main method
In the following code snippet, the main method creates an AVL tree by inserting nodes
with the 5, 3, 8, 7, 6, and 10 keys. Nodes with the 3 and 7 keys are removed. The tree data
structure is converted in to JSON in bytes. The JSON bytes are printed after being changed
to a string:*/
func main() {

	var treeNode *TreeNode
	fmt.Println("Tree is empty")
	var avlTree []byte
	avlTree, _ = json.MarshalIndent(treeNode, "", "   ")
	fmt.Println(string(avlTree))

	fmt.Println("\n Add Tree")
	InsertNode(&treeNode, integerKey(5))
	InsertNode(&treeNode, integerKey(3))
	InsertNode(&treeNode, integerKey(8))
	InsertNode(&treeNode, integerKey(7))
	InsertNode(&treeNode, integerKey(6))
	InsertNode(&treeNode, integerKey(10))
	avlTree, _ = json.MarshalIndent(treeNode, "", "   ")
	fmt.Println(string(avlTree))

	fmt.Println("\n Delete Tree")
	RemoveNode(&treeNode, integerKey(3))
	RemoveNode(&treeNode, integerKey(7))
	avlTree, _ = json.MarshalIndent(treeNode, "", "   ")
	fmt.Println(string(avlTree))
}
