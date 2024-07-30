package main

/*Binary Tree
A binary tree is a type tree in which each node has at most two children (0, 1 or 2), which are referred to as the left
child and the right child.
Below is a node of the binary tree with "a" stored as data and whose left child (lChild) and whose right child (rchild)
both pointing towards null.
Below is a class definition used to define node.*/
type Node struct {
	value       int
	left, right *Node
}
type Tree struct {
	root *Node
}

// Below is a binary tree whose nodes contains data from 1 to 10
// In the rest of the book, binary tree will be represented as below:
// Properties of Binary tree are:
// 1. The maximum number of nodes on level i of a binary tree is 2i , where i >= 1
// 2. The maximum number of nodes in a binary tree of depth k is 2k+1, where k >= 1
// 3. There is exactly one path from the root to any nodes in a tree.
// 4. A tree with N nodes have exactly N-1 edges connecting these nodes.
// 5. The height of a complete binary tree of N nodes is log2N.

// var tree = &Tree{
//     root: &Node{
//         value: 1,
//         left: &Node{
//             value: 2,
//             left:  &Node{value: 4},
//             right: &Node{value: 5},
//         },
//         right: &Node{
//             value: 3,
//             left:  &Node{value: 6},
//             right: &Node{value: 7},
//         },
//     },
// }
// Types of Binary trees

// Complete binary tree

// In a complete binary tree, every level except the last one is completely filled. All nodes in the left are filled first, then
// the right one. A binary heap is an example of a complete binary tree.

/*Full/ Strictly binary tree
The full binary tree is a binary tree in which each node has exactly zero or two children.*/

/*Perfect binary tree
The perfect binary tree is a type of full binary tree in which each non-leaf node has exactly two child nodes. All leaf
nodes have identical path length and all possible node slots are occupied*/

// Right skewed binary tree
// A binary tree in which either each node is have a right child or no child (leaf) is called as right skewed binary tree

/*Left skewed binary tree
A binary tree in which either each node is have a left child or no child (leaf) is called as Left skewed binary tree*/

/*Height-balanced Binary Tree
A height-balanced binary tree is a binary tree such that the left & right subtrees for any given node differ in height by
max one. AVL tree and RB tree are an example of height balanced tree we will discuss these trees later in this chapter.
Note: Each complete binary tree is a height-balanced binary tree*/
