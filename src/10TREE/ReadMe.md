Introduction
We have already read about various linear data structures like a list, linked list, stack, queue etc.
Both list and linked list have a drawback of linear time required for searching an element.
A tree is a nonlinear data structure, which is used to represent hierarchical relationships (parent-child relationship).
Each node is connected by another node by directed edges.
Example 1: Tree in organization

Example 2: Tree in a file system

# Terminology in tree
## Root: 
The root of the tree is the only node without any incoming edges. It is the top node of a tree.
## Node: 
It is a fundamental element of a tree. Each node has data and two references that may point to null or its children
## Edge: 
It is also a fundamental part of a tree, which is used to connect two nodes.
## Path: 
A path is an ordered list of nodes that are connected by edges.
## Leaf: 
A leaf node is a node that has no children.
## Height of the tree: 
The height of a tree is the number of edges on the longest path between the root and a leaf.
## The level of node:
 The level of a node is the number of edges on the path from the root node to that node.
## Children: 
Nodes that have incoming edges from the same node to be said to be the children of that node.
## Parent: 
Node is a parent of all the child nodes that are linked by outgoing edges.
## Sibling:
 Nodes in the tree that are children of the same parent are called siblings.
## Ancestor:
 A node reachable by repeated moving from child to parent.