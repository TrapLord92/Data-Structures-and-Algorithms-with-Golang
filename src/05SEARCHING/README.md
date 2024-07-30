# Introduction
Searching is the process of finding a particular item in a collection of items. The item may be a keyword in a file, a
record in a database, a node in a tree or a value in a list etc.

/************************************************************************************************************************/
# Why Searching?

Imagine you are in a library with millions of books. You want to get a specific book with specific title. How will you
find? You will search the book in the section of library, which contains the books whose name starts with the initial
letter of the desired book. Then you continue matching with a whole book title until you find your book. (By doing this
small heuristic you have reduced the search space by a factor of 26, consider we have an equal number of books whose
title begin with particular char.)

Similarly, computer stores lots of information and to retrieve this information efficiently, we need very efficient
searching algorithms. To make searching efficient, we keep the data in some proper order. There are certain ways of
organizing the data. If you keep the data in proper order, it is easy to search required element. For example, 
# Sorting is one of the process for making data organized.

# Different Searching Algorithms
· Linear Search – Unsorted Input
· Linear Search – Sorted Input
· Binary Search (Sorted Input)
· String Search: Tries, Suffix Trees, Ternary Search.
· Hashing and Symbol Tables

# How sorting is useful in Selection Algorithm?
Selection problems can be converted to sorting problems. Once the list is sorted, it is easy to find the minimum /
maximum (or desired element) from the sorted list. The method ‘Sorting and then Selecting’ is inefficient for selecting
a single element, but it is efficient when many selections need to be made from the list. It is because only one initial
expensive sort is needed, followed by many cheap selection operations.
For example, if we want to get the maximum element from a list. After sorting the list, we can simply return the last
element from the list. What if we want to get second maximum. Now, we do not have to sort the list again and we can
return the second last element from the sorted list. Similarly, we can return the kth maximum element by just one scan
of the sorted list.
So with the above discussion, sorting is used to improve the performance. In general this method requires O(nlogn)
(for sorting) time. With the initial sorting, we can answer any query in one scan, O(n).