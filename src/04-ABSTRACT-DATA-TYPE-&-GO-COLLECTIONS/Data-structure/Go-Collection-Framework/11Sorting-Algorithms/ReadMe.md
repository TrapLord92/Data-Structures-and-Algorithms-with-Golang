Sorting Algorithms
Sorting is the process of placing elements from a collection into ascending or descending order.
Sorting arranges data elements in order so that searching become easier.
There are good sorting functions available which does sorting in O(nlogn) time, so in this book when we need sorting
we will use sort() function and will assume that the sorting is done in O(nlogn) time.
Counting Sort
Counting sort is the simplest and most efficient type of sorting. Counting sort has a strict requirement of a predefined
range of data.
Like, sort how many people are in which age group. We know that the age of people can vary between 1 and 130.
If we know the range of input, then sorting can be done using counting in O(n+k). Where n is the number of people and
k is the max age possible let us