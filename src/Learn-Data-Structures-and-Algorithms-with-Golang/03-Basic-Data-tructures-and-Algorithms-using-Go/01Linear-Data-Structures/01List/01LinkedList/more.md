"Contiguously in memory" means that the elements of a data structure are stored next to each other in a continuous block of memory. Here's a deeper dive into what this means:

### Memory Contiguity:
1. **Memory Layout**: When a data structure is stored contiguously, all its elements are located in adjacent memory addresses. For example, if the first element of an array is stored at memory address `0x1000`, and each element takes up 4 bytes, then the second element will be at `0x1004`, the third at `0x1008`, and so on.

2. **Performance**: This layout allows for fast access to any element in the data structure using simple arithmetic on the index. For example, accessing the `i`-th element in an array can be done in constant time O(1) because the address can be directly calculated as `base_address + i * element_size`.

3. **Example in Arrays**:
   - Suppose you have an array of integers:
     ```go
     arr := [5]int{10, 20, 30, 40, 50}
     ```
     If the base address of `arr` is `0x1000`, the memory layout would be:
     ```
     Address     Value
     0x1000      10
     0x1004      20
     0x1008      30
     0x100C      40
     0x1010      50
     ```
   - Each integer typically takes 4 bytes, so each subsequent element is exactly 4 bytes away from the previous one.

### Non-Contiguous Memory:
In contrast, non-contiguous memory means the elements of a data structure are scattered throughout memory, linked together using pointers. This is typical of data structures like linked lists.

1. **Memory Layout**: Elements (nodes) can be stored at any random memory address. Each node contains the data and a pointer/reference to the next node.
   
2. **Performance**: Accessing elements is slower compared to contiguous structures because you must traverse the list from the beginning to reach a particular element.

3. **Example in Linked Lists**:
   - Suppose you have a linked list with nodes containing integer values:
     ```go
     type Node struct {
         value int
         next  *Node
     }
     ```
     Let's create and link some nodes:
     ```go
     node1 := &Node{value: 10}
     node2 := &Node{value: 20}
     node3 := &Node{value: 30}
     node4 := &Node{value: 40}
     node5 := &Node{value: 50}
     
     node1.next = node2
     node2.next = node3
     node3.next = node4
     node4.next = node5
     ```
     The memory addresses of these nodes could be completely arbitrary, like:
     ```
     Node       Address
     node1      0x2000
     node2      0x3000
     node3      0x5000
     node4      0x7000
     node5      0x9000
     ```
     Each node knows the address of the next node, allowing traversal from one node to the next.

### Summary:
- **Contiguous Memory**: Elements are stored next to each other. Common in arrays. Allows for fast, direct access.
- **Non-Contiguous Memory**: Elements are scattered, with pointers linking them. Common in linked lists. Access involves traversal and is generally slower.

This understanding helps in selecting the right data structure based on the performance needs and memory usage of your application.