Memory Management
Memory management is a way to control and organize memory. Memory divisions are
called blocks, and they are used for running different processes. The basic goal of memory
management algorithms is to dynamically designate segments of memory to programs on
demand. The algorithms free up memory for reuse when the objects in the memory are
never required again. Garbage collection, cache management, and space allocation
algorithms are good examples of memory management techniques. In software
engineering, garbage collection is used to free up memory that's been allocated to those
objects that won't be used again, thus helping in memory management. The cache provides
in-memory storage for data. You can sort the data in the cache into locale-specific groups.
The data can be stored using key and value sets.
This chapter covers the garbage collection, cache management, and space allocation
algorithms. The memory management algorithms are presented with code samples and
efficiency analyses. The following topics will be covered in this chapter:
Garbage collection
Cache management
Space allocation
Concepts—Go memory management