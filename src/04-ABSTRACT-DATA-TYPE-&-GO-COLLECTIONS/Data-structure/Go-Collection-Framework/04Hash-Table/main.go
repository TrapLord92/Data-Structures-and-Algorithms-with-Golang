package main

/*Hash-Table
A Hash-Table is a data structure that maps keys to values. Each position of the Hash-Table is called a slot. The Hash-
Table uses a hash function to calculate an index of a list. We use the Hash-Table when the number of key actually
stored is small relatively to the number of possible keys.
The process of storing objects using a hash function is as follows:
1. Create a list of size M to store objects; this list is called Hash-Table.
2. Find a hash code of an object by passing it through the hash function.
3. Take module of hash code by the size of Hash-Table to get the index where objects will be stored.
4. Finally store these objects in the designated index.
The process of searching objects in Hash-Table using a hash function is as follows:
1. Find a hash code of the object we are searching for by passing it through the hash function.
2. Take module of hash code by the size of Hash-Table to get the index where object is stored.
3. Finally, retrieve the object from the designated index.
Hash-Table Abstract Data Type (ADT)
ADT of Hash-Table contains the following functions:
Insert(x): Add object x to the data set.
Delete(x): Delete object x from the data set.
Search(x): Search object x in data set.
The Hash-Table is a useful data structure for implementing dictionary. The average time to search for an element in a
Hash-Table is O(1). A Hash Table generalizes the notion of a list.*/
