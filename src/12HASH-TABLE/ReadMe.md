Introduction
In the searching chapter, we have looked into various searching techniques. Consider a problem of searching a value in
an array. If the array is not sorted then we have no other option, but to look into each element one by one so the
searching Time Complexity will be O(n). If the array is sorted then we can search the value in O(logn) logarithmic
time using binary search.
What if the possible location / index of the value that we are looking in the array is returned by a magic function in
constant time? We can directly go into that location and tell whether the value we are searching for is present or not in
just O(1) constant time. Such a function is called a Hash function.
The process of storing objects using a hash function is as follows:
1. Create a list of size M to store objects; this list is called Hash-Table.
2. Find a hash code of an object by passing it through the hash function.
3. Take module of hash code by the size of Hashtable to get the index of the table where objects will be stored.
4. Finally store these objects in the designated index.
The process of searching objects in Hash-Table using a hash function is as follows:
1. Find a hash code of the object we are searching for by passing it through the hash function.
2. Take module of hash code by the size of Hashtable to get the index of the table where objects are stored.
3. Finally, retrieve the object from the designated index.