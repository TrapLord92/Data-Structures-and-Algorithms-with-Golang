Introduction
Sorting is the process of placing elements from a list into ascending or descending order. For example, when we play
cards, sort cards, according to their value so that we can find the required card easily.

When we go to some library, the books are arranged according to streams (Algorithm, Operating systems, Networking
etc.). Sorting arranges data elements in order so that searching become easier. When books are arranged in proper
indexing order, then it is easy to find a book we are looking for.
This chapter discusses algorithms for sorting a list of N items. Understanding sorting algorithms are the first step
towards understanding algorithm analysis. Many sorting algorithms are developed and analyzed.

# A sorting algorithm like Bubble-Sort, Insertion-Sort and Selection-Sort are easy to implement and are suitable for the
small input set. However, for large dataset they are slow.

# A sorting algorithm like Merge-Sort, Quick-Sort and Heap-Sort are some of the algorithms that are suitable for sorting
large dataset. However, they are overkill if we want to sort the small dataset.
Some algorithm, which is suitable when we have some range information on input data.
Some other algorithm is there to sort a huge data set that cannot be stored in memory completely, for which external
sorting technique is developed.
Before we start a discussion of the various algorithms one by one. First, we should look at comparison function that is
used to compare two values.

# Less function will return 1 if value1 is less than value2 otherwise it will return 0.
func less(value1 int, value2 int) bool {
return value1 < value2
}
# More function will return 1 if value1 is more than value2 otherwise it will return 0.
func more(value1 int, value2 int) bool {
return value1 > value2
}
The value in various sorting algorithms is compared using one of the above functions and it will be swapped depending
upon the return value of these functions. If more() comparison function is used, then sorted output will be increasing in
order and if less() is used then resulting output will be in descending order.

# Type of Sorting
Internal Sorting: All the elements can be read into memory at the same time and sorting is performed in memory.
1. Selection-Sort
2. Insertion-Sort
3. Bubble-Sort
4. Quick-Sort
5. Merge-Sort
# External Sorting: In this, the dataset is so big that it is impossible to load the whole dataset into memory so sorting is
done in chunks.
1. Merge-Sort
Three things to consider in choosing, sorting algorithms for application:
1. Number of elements in list
2. A number of different orders of list required
3. The amount of time required to move the data or not move the data