Symbol tables
A symbol table is present in memory during the program translation process. It can be
present in program binaries. A symbol table contains the symbol's name, location, and
address. In Go, the gosym package implements access to the Go symbol and line number
tables. Go binaries generated by the GC compilers have the symbol and line number tables.
A line table is a data structure that maps program counters to line numbers.
Containers
The containers package provides access to the heap, list, and ring functionalities in Go.
Containers are used in social networks, knowledge graphs, and other areas. Containers are
lists, maps, slices, channels, heaps, queues, and treaps. Lists were introduced in Chapter 1,
Data Structures and Algorithms. Maps and slices are built-in containers in Go. Channels in
Go are called queues. A heap is a tree data structure. This data structure satisfies the heap
property. A queue is modeled as a heap in Chapter 3, Linear Data Structures. A treap is a
mix of a tree and a heap. It is a binary tree with keys and values and a heap that maintains
priorities.
Non-Linear Data Structures Chapter 4
[ 137 ]
A ring is called a circular linked list and is presented in the next section.