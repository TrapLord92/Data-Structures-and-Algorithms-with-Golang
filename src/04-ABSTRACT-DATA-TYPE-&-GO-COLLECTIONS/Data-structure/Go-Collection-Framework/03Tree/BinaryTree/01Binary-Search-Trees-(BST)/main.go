package main

/*Binary Search Trees (BST)
A binary search tree (BST) is a binary tree on which nodes are ordered in the following way:
1. The key in the left subtree is less than the key in its parent node.
2. The key in the right subtree is greater or equal the key in its parent node.
Binary Search Tree ADT Operations
· Insert(k): Insert an element k into the tree.
· Delete(k): Delete an element k from the tree.
· Search(k): Search a particular value k into the tree if it is present or not.
· FindMax(): Find the maximum value stored in the tree.
· FindMin(): Find the minimum value stored in the tree.
The average Time Complexity of all the above operations on a binary search tree is O(log n), the case when the tree is
balanced. The worst-case Time Complexity will be O(n) when the tree is skewed. A binary tree is skewed when tree is
not balanced.
There are two types of skewed tree.
1. Right Skewed binary tree: A binary tree in which each node is having either a right child or no child.
2. Left Skewed binary tree: A binary tree in which each node is having either a left child or no child.
Balanced Binary search tree
There are few binary search tree, which keeps themselves balanced. Most important among them are Red-Black Tree
and AVL tree. Ordered dictionary in collections is implemented using Red Black-Tree.*/
