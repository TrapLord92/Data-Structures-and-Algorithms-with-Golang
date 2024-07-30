Introduction
Let us suppose we have a list that contains following five elements 1, 2, 4, 5, 6. We want to insert a new element with
value “3” in between “2” and “4”. In the list, we cannot do so easily. We need to create another list that is long enough
to store the current values and one more space for “3”. Then we need to copy these elements in the new space. This
copy operation is inefficient. To remove this copy operation linked list is used.

# Linked List
The linked list is a list of items, called nodes. Nodes have two parts, value part and link part. Value part is used to
stores the data. Either the value part of the node can be a basic data-type like an integer or it can be some other datatype
like an object of some class.
The link part is a reference, which is used to store addresses of the next element in the list.

# Types of Linked list
There are different types of linked lists. The main difference among them is how their nodes refer to each other.
## Singly Linked List
Each node (Except the last node) has a reference to the next node in the linked list. The link portion of node contains
the address of the next node. The link portion of the last node contains the value null
## Doubly Linked list
The node in this type of linked list has reference to both previous and the next node in the list.
## Circular Linked List
This type is similar to the singly linked list except that the last element have reference to the first node of the list. The
link portion of the last node contains the address of the first node.

## The various parts of linked list
1. Head: Head is a reference that holds the address of the first node in the linked list.
2. Nodes: Items in the linked list are called nodes.
3. Value: The data that is stored in each node of the linked list.
4. Link: Link part of the node is used to store the reference of other node.
a. We will use “next” and “prev” to store address of next or previous node.