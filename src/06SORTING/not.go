package sorting

/*Generalized Bucket Sort
There are cases when the element falling into a bucket are not unique but are in the same range. When we want to sort
an index of a name, we can use the reference bucket to store names.
The buckets are already sorted and the elements inside each bucket can be kept sorted by using an Insertion-Sort
algorithm. We are leaving this generalized bucket sort implementation to the reader of this book. The similar data
structure will be defined in the coming chapter of Hash-Table using separate chaining.
Heap-Sort
Heap-Sort we will study in the Heap chapter.
Complexity Analysis:
Data structure List
Worst case performance O(nlogn)
Average case performance O(nlogn)
Worst case Space Complexity O(1)
Tree Sorting
In-order traversal of the binary search tree can also be seen as a sorting algorithm. We will see this in binary search tree
section of tree chapter.
Complexity Analysis:
Worst Case Time Complexity O(n2)
Best Case Time Complexity O(nlogn)
Average Time Complexity O(nlogn)
Space Complexity O(n)
Stable Sorting Yes
External Sort (External Merge-Sort)
When data need to be sorted is huge and it is not possible to load it completely in memory (RAM), for such a dataset
we use external sorting. Such data is sorted using external Merge-Sort algorithm. First data is picked in chunks and it is
sorted in memory. Then this sorted data is written back to disk. Whole data are sorted in chunks using Merge-Sort.
Now we need to combine these sorted chunks into final sorted data.
Then we create queues for the data, which will read from the sorted chunks. Each chunk will have its own queue. We
will pop from this queue and these queues are responsible for reading from the sorted chunks. Let us suppose we have
K different chunks of sorted data each of length M.
The third step is using a Min-Heap, which will take input data from each of this queue. It will take one element from
each queue. The minimum value is taken from the Heap and added to the final sorted element output. Then queue from
which this min element is inserted in the heap will again popped and one more element from that queue is added to the
Heap. Finally, when the data is exhausted from some queue that queue is removed from the input list. Finally, we will
get a sorted data came out from the heap.
We can optimize this process further by adding an output buffer, which will store data coming out of Heap and will do
a limited number of the write operation in the final Disk space.
Note: No one will be asking to implement external sorting in an interview, but it is good to know about it.
Comparisons of the various sorting algorithms.
Below is comparison of various sorting algorithms:
Sort Average Time Best Time Worst Time Space Stable
Bubble Sort O(n2) O(n2) O(n2) O(1) Yes
Modified Bubble
Sort O(n2) O(n) O(n2) O(1) Yes
Selection Sort O(n2) O(n2) O(n2) O(1) No
Insertion Sort O(n2) O(n) O(n2) O(1) Yes
Heap Sort O(n * log(n)) O(n * log(n)) O(n * log(n)) O(1) No
Merge Sort O(n * log(n)) O(n * log(n)) O(n * log(n)) O(n) Yes
Quick Sort O(n * log(n)) O(n * log(n)) O(n2)
O(n) worst case
O( log(n) ) average case No
Bucket Sort O(n k) O(n k) O(n k) O(n k) Yes
Selection of Best Sorting Algorithm
No sorting algorithm is perfect. Each of them has their own pros and cons. Let us read one by one:
Quick-Sort: When you do not need a stable sort and average case performance matters more than worst-case
performance. When data is random, we prefer the Quick-Sort. Average case Time Complexity of Quick-Sort is
O(nlogn) and worst-case Time Complexity is O(n2). Space Complexity of Quick-Sort is O(logn) auxiliary storage,
which is stack space used in recursion.
Merge-Sort: When you need a stable sort and Time Complexity of O(nlogn), Merge-Sort is used. In general, Merge-
Sort is slower than Quick-Sort because of lot of copy happening in the merge phase. There are two uses of Merge-Sort
when we want to merge two sorted linked lists and Merge-Sort is used in external sorting.
Heap-Sort: When you do not need a stable sort and you care more about worst-case performance than average case
performance. It has guaranteed to be O(nlogn), and uses O(1) auxiliary space, meaning that you will not unpredictably
run out of memory on very large inputs.
Insertion-Sort: When we need a stable sort, When N is guaranteed to be small, including as the base case of a Quick-
Sort or Merge-Sort. Worst-case Time Complexity is O(n2). It has a very small constant factor multiplied to calculate
actual time taken. Therefore, for smaller input size it performs better than Merge-Sort or Quick-Sort. It is also useful
when the data is already pre-sorted. In this case, its running time is O(N).
Bubble-Sort: Where we know the data is nearly sorted. Say only two elements are out of place. Then in one pass,
Bubble Sort will make the data sorted and in the second pass, it will see everything is sorted and then exit. Only takes 2
passes of the list.
Selection-Sort: Best Worst Average Case running time all O(n2). It is only useful when you want to do something
quick. They can be used when you are just doing some prototyping.
Counting-Sort: When you are sorting data within a limited range.
Radix-Sort: When log(N) is significantly larger than K, where K is the number of radix digits.
Bucket-Sort: When your input is more or less uniformly distributed.
Note: A stable sort is one that has guaranteed not to reorder elements with identical keys.
Exercise
1. Given a text file, print the words with their frequency. Now print the kth word in term of frequency.
Hint:-
a) First approach may be you can use the sorting and return the kth element.
b) Second approach: You can use the kth element quick select algorithm.
c) Third approach: You can use Hashtable or Trie to keep track of the frequency. Use Heap to get the Kth element.
2. Given K input streams of number in sorted order. You need to make a single output stream, which contains all the
elements of the K streams in sorted order. The input streams support ReadNumber() operation and output stream
support WriteNumber() operation.
Hint:-
a) Read the first number from all the K input streams and add them to a Priority Queue. (Nodes should keep track
of the input stream)
b) Dequeue one element at a time from PQ, Put this element value to the output stream, Read the input stream
number and from the same input stream add another element to PQ.
c) If the stream is empty, just continue
d) Repeat until PQ is empty.
3. Given K sorted Lists of fixed length M. Also, given a final output list of length M*K. Give an efficient algorithm to
merge all the Lists into the final list, without using any extra space.
Hint: you can use the end of the final list to make PQ.
4. How will you sort 1 PB numbers? 1 PB = 1000 TB.
5. What will be the complexity of the above solution?
6. Any other improvement on question 3 solution if the number of cores is eight.
7. Given an integer list that support three function findMin, findMax, findMedian. Sort the list.
8. Given a pile of patient files of High, mid and low priority. Sort these files such that higher priority comes first, then
mid and last low priority.
Hint: Bucket sort.
9. Write pros and cons of Heap-Sort, Merge-Sort and Quick-Sort.
10. Given a rotated - sorted list of N integers. (The list was sorted then it was rotated some arbitrary number of times.)
If all the elements in the list were unique the find the index of some value.
Hint: Modified binary search
11. In the problem 9, what if there are repetitions allowed and you need to find the index of the first occurrence of the
element in the rotated-sorted list.
12. Merge two sorted Lists into a single sorted list.
Hint: Use merge method of Merge-Sort.
13. Given a list contain 0’s and 1’s, sort the list such that all the 0’s come before 1’s.
14. Given a list of English characters, sort the list in linear time.
15. Write a method to sort a list of strings so that all the anagrams are next to each other.
Hint:-
a) Loop through the list.
b) For each word, sort the characters and add it to the hash map with keys as sorted word and value as the original
word. At the end of the loop, you will get all anagrams as the value to a key (which is sorted by its constituent
chars).
c) Iterate over the hashmap, print all values of a key together and then move to the next key.
Space Complexity: O(n), Time Complexity: O(n)*/
