package hashtable

/*Hash-Table
A Hash-Table is a data structure that maps keys to values. Each position of the Hash-Table is called a slot. The Hash-
Table uses a hash function to calculate an index of a list. We use the Hash-Table when the number of keys actually
stored is small relatively to the number of possible keys.
Hash-Table Abstract Data Type (ADT)
ADT of Hash-Table contains the following functions:
1. Insert(x), add object x to the data set.
2. Delete(x), delete object x from the data set.
3. Search(x), search object x in data set.
Hash Function
A hash function is a function that generates an index in a table for a given object.
An ideal hash function that generate a unique index for every object is called the perfect hash function.
Example 12.1: Most simple hash function
func (ht *HashTable) ComputeHash(key int) int {
return key % ht.tableSize
}
There are many hash functions. The above function is a very simple hash function. Various hash generation logics will
be added to this function to generate a better hash.
Collisions
When a hash function generates the same index for the two or more different objects, the problem known as the
collision. Ideally, hash function should return a unique address for each key, but practically it is not possible.
Properties of good hash function:
1. It should provide a uniform distribution of hash values. A non-uniform distribution increased the number of
collisions and the cost of resolving them.
2. Choose a hash function, which can be computed quickly and returns values within the range of the Hash-Table.
3. Chose a hash function with a good collision resolution algorithm which can be used to compute alternative index if
the collision occurs.
4. Choose a hash function, which uses the necessary information provided in the key.
5. It should have high load factor for a given set of keys.
Load Factor
Load factor = Number of elements in Hash-Table / Hash-Table size
Based on the above definition, Load factor tells whether the hash function is distributing the keys uniformly or not.
Therefore, it helps in determining the efficiency of the hashing function. It also works as decision parameter when we
want to expand or rehash the existing Hash-Table entries.
Collision Resolution Techniques
Hash collisions are practically unavoidable when hashing large number of objects. Techniques that are used to find the
alternate location in the Hash-Table is called collision resolution. There are a number of collision resolution techniques
to handle the collision in hashing.
Most common and widely used techniques are:
· Open addressing
· Separate chaining
Hashing with Open Addressing
When using linear open addressing, the Hash-Table is represented by a one-dimensional list with indices that range
from 0 to the desired table size-1.
One method of resolving collision is the look into a Hash-Table and find another free slot the hold the object that have
caused the collision. A simple way is to move from one slot to another in some sequential order until we find a free
space. This collision resolution process is called Open Addressing.
Linear Probing
In Linear Probing, we try to resolve the collision of an index of a Hash-Table by sequentially searching the Hash-Table
free location. Let us suppose, if k is the index retrieved from the hash function. If the kth index is already filled then we
will look for (k+1) %M, then (k+2) %M and so on. When we get a free slot, we will insert the object into that free slot.
Example 12.2: The resolver function of linear probing
func (ht *HashTable) ResolverFun(index int) int {
return index
}
Quadratic Probing
In Quadratic Probing, we try to resolve the collision of the index of a Hash-Table by quadratic ally increasing the
search index free location. Let us suppose, if k is the index retrieved from the hash function. If the kth index is already
filled then we will look for (k+1^2) %M, then (k+2^2) %M and so on. When we get a free slot, we will insert the object
into that free slot.
Example 12.3: The resolver function of quadratic probing
func (ht *HashTable) ResolverFun(index int) int {
return index * index
}
Table size should be a prime number to prevent early looping should not be too close to 2powN
Linear Probing implementation
Example 12.4: Below is a linear probing collision resolution Hash-Table implementation.
const (
EmptyNode byte = iota
LazyDeleted
FillledNode
)
type HashTable struct {
Arr []int
Flag []byte
tableSize int
}
func (ht *HashTable) Init(tSize int) {
ht.tableSize = tSize
ht.Arr = make([]int, (tSize + 1))
ht.Flag = make([]byte, (tSize + 1))
}
Table list size will be 50 and we have defined two constant values EMPTY_NODE and LAZY_DELETED.
func (ht *HashTable) ComputeHash(key int) int {
return key % ht.tableSize
}
This is the most simple hash generation function, which just take the modulus of the key.
func (ht *HashTable) ResolverFun(index int) int {
return index
}
When the hash index is already occupied by some element the value will be placed in some other location to find that
new location resolver function is used.
Hash-Table has two component one is table size and other is reference to list.
Example 12.5:
func (ht *HashTable) Add(value int) bool {
hashValue := ht.ComputeHash(value)
for i := 0; i < ht.tableSize; i++ {
if ht.Flag[hashValue] == EmptyNode || ht.Flag[hashValue] == LazyDeleted {
ht.Arr[hashValue] = value
ht.Flag[hashValue] = FillledNode
return true
}
hashValue += ht.ResolverFun(i)
hashValue %= ht.tableSize
}
return false
}
An insert node function is used to add values to the list. First hash is calculated. Then we try to place that value in the
Hash-Table. We look for empty node or lazy deleted node to insert value. In case insert did not success, we try new
location using a resolver function.
Example 12.6:
func (ht *HashTable) Find(value int) bool {
hashValue := ht.ComputeHash(value)
for i := 0; i < ht.tableSize; i++ {
if ht.Flag[hashValue] == EmptyNode {
return false
}
if ht.Flag[hashValue] == FillledNode && ht.Arr[hashValue] == value {
return true
}
hashValue += ht.ResolverFun(i)
hashValue %= ht.tableSize
}
return false
}
Find node function is used to search values in the array. First hash is calculated. Then we try to find that value in the
Hash-Table. We look for over desired value or empty node. In case we find the value that we are looking, then we
return that value or in case it is not found we return -1. We use a resolver function to find the next probable index to
search.
Example 12.7:
func (ht *HashTable) Remove(value int) bool {
hashValue := ht.ComputeHash(value)
for i := 0; i < ht.tableSize; i++ {
if ht.Flag[hashValue] == EmptyNode {
return false
}
if ht.Flag[hashValue] == FillledNode && ht.Arr[hashValue] == value {
ht.Flag[hashValue] = LazyDeleted
return true
}
hashValue += ht.ResolverFun(i)
hashValue %= ht.tableSize
}
return false
}
Delete node function is used to delete values from a Hashtable. We do not actually delete the value we just mark that
value as LAZY_DELETED. Same as the insert and search we use resolverFun to find the next probable location of the
key.
Example 12.8:
func (ht *HashTable) Print() {
fmt.Println("\nValues Stored in HashTable are::")
for i := 0; i < ht.tableSize; i++ {
if ht.Flag[i] == FillledNode {
fmt.Println("Node at index [", i, " ] :: ", ht.Arr[i])
}
}
}
func main() {
ht := new(HashTable)
ht.Init(1000)
ht.Add(89)
fmt.Println("89 found : ", ht.Find(89))
ht.Remove(89)
fmt.Println("89 found : ", ht.Find(89))
ht.Print()
}
Print method print the content of hash table. Main function demonstrating how to use hash table.
Quadratic Probing implementation.
Everything will be same as linear probing implementation only resolver function will be changed.
func (ht *HashTable) ResolverFun(index int) int {
return index * index
}
Hashing with separate chaining
Another method for collision resolution is based on an idea of putting the keys that collide in a linked list. This method
is called separate chaining. To speed up search we use Insertion-Sort or keeping the linked list sorted.
Separate Chaining implementation
‘Example 12.9: Below is separate chaining implementation of hash tables.
type Node struct {
value int
next *Node
}
type HashTableSC struct {
listArray [](*Node)
tableSize int
}
func (h *HashTableSC) Init() {
h.tableSize = 101
h.listArray = make([](*Node), h.tableSize)
for i := 0; i < h.tableSize; i++ {
h.listArray[i] = nil
}
}
func (h *HashTableSC) ComputeHash(key int) int {
hashValue := key
return hashValue % h.tableSize
}
func (h *HashTableSC) Add(value int) {
index := h.ComputeHash(value)
temp := new(Node)
temp.value = value
temp.next = h.listArray[index]
h.listArray[index] = temp
}
func (h *HashTableSC) Remove(value int) bool {
index := h.ComputeHash(value)
var nextNode, head *Node
head = h.listArray[index]
if head != nil && head.value == value {
h.listArray[index] = head.next
return true
}
for head != nil {
nextNode = head.next
if nextNode != nil && nextNode.value == value {
head.next = nextNode.next
return true
}
head = nextNode
}
return false
}
func (h *HashTableSC) Print() {
for i := 0; i < h.tableSize; i++ {
head := h.listArray[i]
if head != nil {
fmt.Print("\nValues at index :: ", i, " Are :: ")
}
for head != nil {
fmt.Print(head.value, " ")
head = head.next
}
}
fmt.Println()
}
func (h *HashTableSC) Find(value int) bool {
index := h.ComputeHash(value)
head := h.listArray[index]
for head != nil {
if head.value == value {
return true
}
head = head.next
}
return false
}
func main() {
ht := new(HashTableSC)
ht.Init()
ht.Add(89)
ht.Print()
fmt.Println("89 found : ", ht.Find(89))
}
Note: It is important to note that the size of the “skip” must be such that all the slots in the table will eventually be
occupied. Otherwise, part of the table will be unused. To ensure this, it is often suggested that the table size being a
prime number. This is the reason we have been using 11 in our examples.
Problems in Hashing
Anagram solver
An anagram is a word or phrase formed by reordering the letters of another word or phrase.
Example 12.10: Two words are anagram if they are of same size and their characters are same.
func IsAnagram(str1 string, str2 string) bool {
size1 := len(str1)
size2 := len(str2)
if size1 != size2 {
return false
}
cm := make(Counter)
for _, ch := range str1 {
cm.Add(ch)
}
for _, ch := range str2 {
cm.Remove(ch)
}
return (len(cm) == 0)
}
Remove Duplicate
Solution: We can use a second list or the same list, as the output list. In the below example Hash-Table is used to solve
this problem.
Example 12.11: Remove duplicates in a list of numbers
func RemoveDuplicate(str string) string {
input := []rune(str)
hs := make(Set)
var output []rune
for _, ch := range input {
if hs.Find(ch) == false {
output = append(output, ch)
hs.Add(ch)
}
}
return string(output)
}
Find Missing
Example 12.12: There is a list of integers we need to find the missing number in the list.
func FindMissing(arr []int, start int, end int) (int, bool) {
hs := make(Set)
for _, i := range arr {
hs.Add(i)
}
for curr := start; curr <= end; curr++ {
if hs.Find(curr) == false {
return curr, true
}
}
return 0, false
}
All the elements in the list is added to a HashTable. The missing element is found by searching into HashTable and
final missing value is returned.
Print Repeating
Example 12.13: Print the repeating integer in a list of integers.
func PrintRepeating(arr []int) {
hs := make(Set)
fmt.Print("Repeating elements are :: ")
for _, val := range arr {
if hs.Find(val) {
fmt.Print(val, " ")
} else {
hs.Add(val)
}
}
fmt.Println()
}
All the values to the hash table when some value came which is already in the hash table then that is the repeated value.
Print First Repeating
Example 12.14: Same as the above problem in this we need to print the first repeating number. Caution should be
taken to find the first repeating number. It should be the one number that is repeating. For example, 1, 2, 3, 2, 1. The
answer should be 1 as it is the first number, which is repeating.
func PrintFirstRepeating(arr []int) {
size := len(arr)
hs := make(Counter)
for i := 0; i < size; i++ {
hs.Add(arr[i])
}
for i := 0; i < size; i++ {
hs.Remove(arr[i])
if hs.Find(arr[i]) {
fmt.Println("First Repeating number is : ", arr[i])
return
}
}
}
Add values to the count map the one that is repeating will have multiple count. Now traverse the list again and see if
the count is more than one. So that is the first repeating.
Exercise
1. Design a number (ID) generator system that generate numbers between 0-99999999 (8-digits).
The system should support two functions:
a. int getNumber();
b. int requestNumber();
getNumber() function should find out a number that is not assigned, then marks it as assigned and return that
number. requestNumber() function checks the number is assigned or not. If it is assigned returns 0, else marks it as
assigned and return 1.
2. Given a large string, find the most occurring words in the string. What is the Time Complexity of the above
solution?
Hint:-
a. Create a Hashtable which will keep track of <word, frequency>
b. Iterate through the string and keep track of word frequency by inserting into Hash-Table.
c. When we have a new word, we will insert it into the Hashtable with frequency 1. For all repetition of the
word, we will increase the frequency.
d. We can keep track of the most occurring words whenever we are increasing the frequency we can see if this
is the most occurring word or not.
e. The Time Complexity is O(n) where n is the number of words in the string and Space Complexity is the
O(m) where m is the unique words in the string.
3. In the above question, What if you are given whole work of OSCAR WILDE, most popular playwrights in the early
1890s.
Hint:-
a. Who knows how many books are there, let us assume there is a lot and we cannot put everything in
memory. First, we need a Streaming Library so that we can read section by section in each document. Then
we need a tokenizer that will give words to our program. In addition, we need some sort of dictionary let us
say we will use HashTable.
b. What you need is - 1. A streaming library tokenizer, 2. A tokenizer 3. A hashmap
Method:
1. Use streamers to find a stream of the given words
2. Tokenize the input text
3. If the stemmed word is in hash map, increment its frequency count else adds a word to hash map with
frequency 1
c. We can improve the performance by looking into parallel computing. We can use the map-reduce to solve
this problem. Multiple nodes will read and process multiple documents. Once they are done with their
processing, then we can use reduce to merge them.
4.4. In the above question, What if we wanted to find the most common PHRASE in his writings.
Hint: - We can keep <phrase, frequency> Hash-Table and do the same process of the 2nd and 3rd problems.
5. Write a hashing algorithm for strings.
Hint: Use Horner's method
func hornerHash(key []rune, tableSize int) int {
size := len(key)
h := 0
for i := 0; i < size; i++ {
h = (32*h + (int)(key[i])) % tableSize
}
return h
}*/
