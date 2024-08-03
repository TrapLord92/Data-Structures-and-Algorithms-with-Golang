package main

// // Synchronized TreeSets
// // Operations that are performed on synchronized TreeSets are synchronized across multiple
// // calls that access the elements of TreeSets. Synchronization in TreeSets is achieved using a
// // sync.RWMutex lock. The lock method on the tree's lock instance is invoked, and the
// // unlock method is deferred before the tree nodes are inserted, deleted, or updated:
// // InsertElement method
// func (tree *BinarySearchTree) InsertElement(key int, value int) {
// 	tree.lock.Lock()
// 	defer tree.lock.Unlock()
// 	var treeNode *TreeNode
// 	treeNode = &TreeNode{key, value, nil, nil}
// 	if tree.rootNode == nil {
// 		tree.rootNode = treeNode
// 	} else {
// 		insertTreeNode(tree.rootNode, treeNode)
// 	}
// }

// // Dynamic Data Structures Chapter 7
// // [ 199 ]
// // Mutable TreeSets
// // Mutable TreeSets can use add, update, and delete operations on the tree and its nodes.
// // insertTreeNode updates the tree by taking the rootNode and treeNode parameters to
// // be updated. The following code snippet shows how to insert a TreeNode with a given
// // rootNode and TreeNode:
// // insertTreeNode method
// func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
// 	if newTreeNode.key < rootNode.key {
// 		if rootNode.leftNode == nil {
// 			rootNode.leftNode = newTreeNode
// 		} else {
// 			insertTreeNode(rootNode.leftNode, newTreeNode)
// 		}
// 	} else {
// 		if rootNode.rightNode == nil {
// 			rootNode.rightNode = newTreeNode
// 		} else {
// 			insertTreeNode(rootNode.rightNode, newTreeNode)
// 		}
// 	}
// }

// /*Let's discuss the different mutable TreeSets in the following sections.
// RemoveNode method
// The RemoveNode method of a BinarySearchTree is as follows:
// // RemoveNode method*/
// // func (tree *BinarySearchTree) RemoveNode(key int) {
// // 	tree.lock.Lock()
// // 	defer tree.lock.Unlock()
// // 	removeNode(tree.rootNode, key)
// // }

// /*Dynamic Data Structures Chapter 7
// [ 200 ]
//  Treeset
// The TreeNode's can be updated by accessing treeset.bst and traversing the binary
//  search tree from the rootNode and the left and right nodes of rootNode, as shown here:*/
// func main1() {

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
