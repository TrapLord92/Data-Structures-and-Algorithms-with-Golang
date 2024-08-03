// package main

// // TreeSet class
// type TreeSet struct {
// 	bst *BinarySearchTree
// }

// // InsertTreeNode method
// /*InsertTreeNode method
// The InsertTreeNode method of the TreeSet class takes treeNodes variable arguments
// of the TreeNode type. In the following code, the elements with the key and value are
// inserted in the binary search tree of TreeSet:*/
// func (treeset *TreeSet) InsertTreeNode(treeNodes ...TreeNode) {
// 	var treeNode TreeNode
// 	for _, treeNode = range treeNodes {
// 		treeset.bst.InsertElement(treeNode.key, treeNode.value)
// 	}
// }

// // Delete method
// /*Delete method
// The Delete method of the TreeSet class is shown in the following code snippet. In this
// method, treeNodes with the provided key are removed:*/
// func (treeset *TreeSet) Delete(treeNodes ...TreeNode) {
// 	var treeNode TreeNode
// 	for _, treeNode = range treeNodes {
// 		treeset.bst.RemoveNode(treeNode.key)
// 	}
// }

// // Search method
// /*Search method
// The Search method of the TreeSet class takes a variable argument named treeNodes of
// the TreeNode type and returns true if one of those treeNodes exists; otherwise, it returns
// false. The code following snippet outlines the Search method:*/
// func (treeset *TreeSet) Search(treeNodes ...TreeNode) bool {
// 	var treeNode TreeNode
// 	var exists bool
// 	for _, treeNode = range treeNodes {
// 		if exists = treeset.bst.SearchNode(treeNode.key); !exists {
// 			return false
// 		}
// 	}
// 	return true
// }

// // String method
// /*The String method
// In the following code snippet, the String method of the TreeSet class returns the string
// version of bst:*/
// func (treeset *TreeSet) String() {

// 	treeset.bst.String()

// }

// // main method
// func main() {
// 	var treeset *TreeSet = &TreeSet{}

// 	treeset.bst = &BinarySearchTree{}

// 	var node1 TreeNode = TreeNode{8, 8, nil, nil}
// 	var node2 TreeNode = TreeNode{3, 3, nil, nil}
// 	var node3 TreeNode = TreeNode{10, 10, nil, nil}
// 	var node4 TreeNode = TreeNode{1, 1, nil, nil}
// 	var node5 TreeNode = TreeNode{6, 6, nil, nil}

// 	treeset.InsertTreeNode(node1, node2, node3, node4, node5)

// 	treeset.String()

// }
package main
