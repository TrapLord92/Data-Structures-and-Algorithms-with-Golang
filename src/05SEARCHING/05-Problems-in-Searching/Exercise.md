TODO

# Find max, appearing element in a list
Given a list of n numbers, find the element, which appears maximum number of times.
First approach: Exhaustive search or Brute force, for each element in array find how many times this particular value
appears in array. Keep track of the maxCount and when some element count is greater than maxCount then update the
maxCount. This is done using two for loop, first loop to select the element and second loop to count the occurrence of
that element.
The Time Complexity is and Space Complexity is
Example 5.9
func getMax(data []int) int {
size := len(data)
max := data[0]
count := 1
maxCount := 1
for i := 0; i < size; i++ {
count = 1
for j := i + 1; j < size; j++ {
if data[i] == data[j] {
count++
}
}
if count > maxCount {
max = data[i]
maxCount = count
}
}
return max
}
Second approach: Sorting, Sort all the elements in the list and after this in a single scan, we can find the counts. Sorting
algorithms take O(n.logn) time and single scan take O(n) time. The Time Complexity of an algorithm is O(n.logn) and
Space Complexity is O(1)
Example 5.10
func getMax2(data []int) int {
size := len(data)
max := data[0]
maxCount := 1
curr := data[0]
currCount := 1
sort.Ints(data) // Sort(data,size)
for i := 1; i < size; i++ {
if data[i] == data[i-1] {
currCount++
} else {
currCount = 1
curr = data[i]
}
if currCount > maxCount {
maxCount = currCount
max = curr
}
}
return max
}
Third approach: Counting, This approach is only possible if we know the range of the input. If we know that, the
elements in the list are in the range 0 to n-1. We can reserve a list of length n and when we see an element, we can
increase its count. In just one single scan, we know the duplicates. If we know the range of the elements, then this is the
fastest way to find the max count.
Counting approach just uses list so to increase count take constant time O(1) so the total Time Complexity of the
algorithm is O(n) time. Space Complexity for creating count list is also O(n)
Example 5.11
func getMax3(data []int, dataRange int) int {
max := data[0]
maxCount := 1
size := len(data)
count := make([]int, dataRange)
for i := 0; i < size; i++ {
count[data[i]]++
if count[data[i]] > maxCount {
maxCount = count[data[i]]
max = data[i]
}
}
return max
}
Majority element in a list
Given a list of n elements. Find the majority element, which appears more than n/2 times. Return 0 in case there is no
majority element.
First approach: Exhaustive search or Brute force, for each element in array find how many times this particular value
appears in array. Keep track of the maxCount and when some element count is greater than maxCount then update the
maxCount. This is done using two for loop, first loop to select the element and second loop to count the occurrence of
that element.
Once we have the final, maxCount we can see if it is greater than n/2, if it is greater than we have a majority if not we
do not have any majority.
The Time Complexity is O(n2) + O(1) = O(n2) and Space Complexity is O(1)
Example 5.12
func getMajority(data []int) (int, bool) {
size := len(data)
max := 0
count := 0
maxCount := 0
for i := 0; i < size; i++ {
for j := i + 1; j < size; j++ {
if data[i] == data[j] {
count++
}
}
if count > maxCount {
max = data[i]
maxCount = count
}
}
if maxCount > size/2 {
return max, true
}
fmt.Println("MajorityDoesNotExist")
return 0, false
}
Second approach: Sorting, Sort all the elements in the list. If there is a majority then the middle element at the index n/2
must be the majority number. So just single scan can be used to find its count and see if the majority is there or not.
Sorting algorithms take O(n.logn) time and single scan take O(n) time.
The Time Complexity of an algorithm is O(n.logn) and Space Complexity is O(1)
Example 5.13
func getMajority2(data []int) (int, bool) {
size := len(data)
majIndex := size / 2
sort.Ints(data) // Sort(data,size)
candidate := data[majIndex]
count := 0
for i := 0; i < size; i++ {
if data[i] == candidate {
count++
}
}
if count > size/2 {
return data[majIndex], true
}
fmt.Println("MajorityDoesNotExist")
return 0, false
}
Third approach: This is a cancelation approach (Moore’s Voting Algorithm), if all the elements stand against the
majority and each element is cancelled with one element of majority if there is majority then majority prevails.
· Set the first element of the list as majority candidate and initialize the count to be 1.
· Start scanning the list.
o If we get some element whose value same as a majority candidate, then we increase the count.
o If we get an element whose value is different from the majority candidate, then we decrement the count.
o If count become 0, that means we have a new majority candidate. Make the current candidate as majority
candidate and reset count to 1.
o At the end, we will have the only probable majority candidate.
· Now scan through the list once again to see if that candidate we found above have appeared more than n/2 times.
Counting approach just scans through list two times. The Time Complexity of the algorithm is O(n) time. Space
Complexity for creating count list is also O(1)
Example 5.14
func getMajority3(data []int) (int, bool) {
majIndex := 0
count := 1
size := len(data)
for i := 1; i < size; i++ {
if data[majIndex] == data[i] {
count++
} else {
count--
}
if count == 0 {
majIndex = i
count = 1
}
}
candidate := data[majIndex]
count = 0
for i := 0; i < size; i++ {
if data[i] == candidate {
count++
}
}
if count > size/2 {
return data[majIndex], true
}
fmt.Println("MajorityDoesNotExist")
return 0, false
}
Find the missing number in a list
Given a list of n-1 elements, which are in the range of 1 to n. There are no duplicates in the list. One of the integer is
missing. Find the missing element.
First approach: Exhaustive search or Brute force, for each value in the range 1 to n, find if there is some element in list
which have the same value. This is done using two for loop, first loop to select value in the range 1 to n and the second
loop to find if this element is in the list or not.
The Time Complexity is O(n2) and Space Complexity is O(1)
Example 5.15
func findMissingNumber(data []int) (int, bool) {
var found int
size := len(data)
for i := 1; i <= size; i++ {
found = 0
for j := 0; j < size; j++ {
if data[j] == i {
found = 1
break
}
}
if found == 0 {
return i, true
}
}
fmt.Println("NoNumberMissing")
return 0, false
}
Second approach: Sorting, Sort all the elements in the list and after this in a single scan, we can find the duplicates.
Sorting algorithms take O(n.logn) time and single scan take O(n) time.
The Time Complexity of an algorithm is O(n.logn) and Space Complexity is O(1)
Third approach: Hash-Table, using Hash-Table, we can keep track of the elements we have already seen and we can
find the missing element in just one scan.
Hash-Table insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time. Space
Complexity is also O(n)
Forth approach: Counting, we know the range of the input so counting will work. As we know that, the elements in the
list are in the range 0 to n-1. We can reserve a list of length n and when we see an element, we can increase its count. In
just one single scan, we know the missing element.
Counting approach just uses a list so insert and find take constant time O(1) so the total Time Complexity of the
algorithm is O(n) time. Space Complexity for creating count list is also O(n)
Fifth approach: You are allowed to modify the given input list. Modify the given input list in such a way that in the
next scan you can find the missing element.
When you scan through the list. When at index “index”, the value stored in the list will be arr[index] so add the number
“n + 1” to arr[ arr[ index]]. Always read the value from the list using a reminder operator “%”. When you scan the list
for the first time and modified all the values, then one single scan you can see if there is some value in the list which is
smaller than “n+1” that index is the missing number.
In this approach, the list is scanned two times and the Time Complexity of this algorithm is O(n). Space Complexity is
O(1)
Sixth approach: Summation formula to find the sum of n numbers from 1 to n. Subtract the values stored in the list and
you will have your missing number.
The Time Complexity of this algorithm is O(n). Space Complexity is O(1)
Seventh approach: XOR approach to find the sum of n numbers from 1 to n. XOR the values stored in the list and you
will have your missing number.
The Time Complexity of this algorithm is O(n). Space Complexity is O(1)
Example 5.16
func findMissingNumber2(data []int, dataRange int) int {
size := len(data)
xorSum := 0
// get the XOR of all the numbers from 1 to dataRange
for i := 1; i <= dataRange; i++ {
xorSum ^= i
}
// loop through the array and get the XOR of elements
for i := 0; i < size; i++ {
xorSum ^= data[i]
}
return xorSum
}
Note: Same problem can be asked in many forms (sometimes you have to do the xor of the range sometime you do
not):1. There are numbers in the range of 1-n out of which all appears single time but one that appear two times.
2. All the elements in the range 1-n are appearing 16 times and one element appear 17 times. Find the element that
appears 17 times.
Find Pair in a list
Given a list of n numbers, find two elements such that their sum is equal to “value”
First approach: Exhaustive search or Brute force, for each element in list find if there is some other element, which
sum up to the desired value. This is done using two for loop, first loop to select the element and second loop to find
another element.
The Time Complexity is O(n2) and Space Complexity is O(1)
Example 5.17
func FindPair(data []int, value int) bool {
size := len(data)
ret := false
for i := 0; i < size; i++ {
for j := i + 1; j < size; j++ {
if (data[i] + data[j]) == value {
fmt.Println("The pair is : ", data[i], ",", data[j])
ret = true
}
}
}
return ret
}
Second approach: Sorting, Steps are as follows:
1. Sort all the elements in the list.
2. Take two variable first and second. Variable first= 0 and second = size -1
3. Compute sum = arr[first]+arr[second]
4. If the sum is equal to the desired value, then we have the solution
5. If the sum is less than the desired value, then we will increase first
6. If the sum is greater than the desired value, then we will decrease the second
7. We repeat the above process until we get the desired pair or we get first >= second
Sorting algorithms take O(n.logn) time and single scan take O(n) time.
The Time Complexity of an algorithm is O(n.logn) and Space Complexity is O(1)
Example 5.18
func FindPair2(data []int, value int) bool {
size := len(data)
first := 0
second := size - 1
ret := false
sort.Ints(data) // Sort(data, size)
for first < second {
curr := data[first] + data[second]
if curr == value {
fmt.Println("The pair is ", data[first], ",", data[second])
ret = true
}
if curr < value {
first++
} else {
second--
}
}
return ret
}
Third approach: Hash-Table, using Hash-Table, we can keep track of the elements we have already seen and we can
find the pair in just one scan.
1. For each element, insert the value in Hashtable. Let say current value is arr[index]
2. If value - arr[index] is in the Hashtable then we have the desired pair.
3. Else, proceed to the next entry in the list.
Hash-Table insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time. Space
Complexity is also O(n)
Example 5.19
func FindPair3(data []int, value int) bool {
s := make(Set)
size := len(data)
ret := false
for i := 0; i < size; i++ {
if s.Find(value - data[i]) {
fmt.Println(i, "The pair is:", data[i],",",(value - data[i]))
ret = true
} else {
s.Add(data[i])
}
}
return ret
}
Forth approach: Counting, This approach is only possible if we know the range of the input. If we know that, the
elements in the list are in the range 0 to n-1. We can reserve a list of length n and when we see an element, we can
increase its count. In place of the Hashtable in the above approach, we will use this list and will find out the pair.
Counting approach just uses a list so insert and find take constant time O(1) so the total Time Complexity of the
algorithm is O(n) time. Space Complexity for creating count list is also O(n)
Find the Pair in two Lists
Given two list X and Y. Find a pair of elements (xi, yi) such that xi∈X and yi∈Y where xi+yi=value.
First approach: Exhaustive search or Brute force, loop through element xi of X and see if you can find (value – xi) in
Y. Two for loop.
The Time Complexity is O(n2) and Space Complexity is O(1)
Second approach: Sorting, Sort all the elements in the second list Y. For each element if X you can see if that element
is there in Y by using binary search.
Sorting algorithms take O(m.logm) and searching will take O(n.logm) time.
The Time Complexity of an algorithm is O(n.logm) or O(m.logm) and Space Complexity is O(1)
Third approach: Sorting, Steps are as follows:
1. Sort the elements of both X and Y in increasing order.
2. Take the sum of the smallest element of X and the largest element of Y.
3. If the sum is equal to value, we got our pair.
4. If the sum is smaller than value, take next element of X
5. If the sum is greater than value, take the previous element of Y
Sorting algorithms take O(n.logn) + O(m.logm) for sorting and searching will take O(n+m) time.
The Time Complexity of an algorithm is O(n.logn) Space Complexity is O(1)
Forth approach: Hash-Table, Steps are as follows:
1. Scan through all the elements in the list Y and insert them into Hashtable.
2. Now scan through all the elements of list X, let us suppose the current element is xi see if you can find (value -
xi) in the Hashtable.
3. If you find the value, you got your pair.
4. If not, then go to the next value in the list X.
Hash-Table insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time.
Space Complexity is also O(n)
Fifth approach: Counting, This approach is only possible if we know the range of the input. Same as Hashtable
implementation just use a simple list in place of Hashtable and you are done.
Counting approach just uses a list so insert and find take constant time O(1) so the total Time Complexity of the
algorithm is O(n) time. Space Complexity for creating count list is also O(n)
Two elements whose sum is closest to zero
Given a List of integers, both +ve and -ve. You need to find the two elements such that their sum is closest to zero.
First approach: Exhaustive search or Brute force, for each element in list find the other element whose value when
added will give minimum absolute value. This is done using two for loop, first loop to select the element and second
loop to find the element that should be added to it so that the absolute of the sum will be minimum or close to zero.
The Time Complexity is O(n2) and Space Complexity is O(1)
Example 5.20
func minAbsSumPair(data []int) {
var sum int
size := len(data)
// Array should have at least two elements
if size < 2 {
fmt.Println("InvalidInput")
}
// Initialization of values
minFirst := 0
minSecond := 1
minSum := abs(data[0] + data[1])
for l := 0; l < size-1; l++ {
for r := l + 1; r < size; r++ {
sum = abs(data[l] + data[r])
if sum < minSum {
minSum = sum
minFirst = l
minSecond = r
}
}
}
fmt.Println(" The two elements with minimum sum are : ", data[minFirst], " , ", data[minSecond])
}
Second approach: Sorting
Steps are as follows:
1. Sort all the elements in the list.
2. Take two variable firstIndex = 0 and secondIndex = size -1
3. Compute sum = arr[firstIndex]+arr[secondIndex]
4. If the sum is equal to the 0 then we have the solution
5. If the sum is less than the 0 then we will increase first
6. If the sum is greater than the 0 then we will decrease the second
7. We repeat the above process 3 to 6, until we get the desired pair or we get first >= second
Example 5.21
func minAbsSumPair2(data []int) {
var sum int
size := len(data)
// Array should have at least two elements
if size < 2 {
fmt.Println("InvalidInput")
}
sort.Ints(data) // Sort(data, size)
// Initialization of values
minFirst := 0
minSecond := size - 1
minSum := abs(data[minFirst] + data[minSecond])
for l, r := 0, (size - 1); l < r; {
sum = (data[l] + data[r])
if abs(sum) < minSum {
minSum = abs(sum) /// just corrected......hemant
minFirst = l
minSecond = r
}
if sum < 0 {
l++
} else if sum > 0 {
r--
} else {
break
}
}
fmt.Println(" The two elements with minimum sum are : ", data[minFirst], " , ", data[minSecond])
}
Find maxima in a bitonic list
A bitonic list comprises of an increasing sequence of integers immediately followed by a decreasing sequence of
integers.
Since the elements are sorted in some order, we should go for algorithm similar to binary search. The steps are as
follows:
1. Take two variable for storing start and end index. Variable start=0 and end=size-1
2. Find the middle element of the list.
3. See if the middle element is the maxima. If yes, return the middle element.
4. Alternatively, If the middle element in increasing part, then we need to look for in mid+1 and end.
5. Alternatively, if the middle element is in the decreasing part, then we need to look in the start and mid-1.
6. Repeat step 2 to 5 until we get the maxima.
Example 5.22
func SearchBitonicArrayMax(data []int) (int, bool) {
size := len(data)
start := 0
end := size - 1
mid := (start + end) / 2
maximaFound := 0
if size < 3 {
fmt.Println("InvalidInput")
return 0, false
}
for start <= end {
mid := (start + end) / 2
if data[mid-1] < data[mid] && data[mid+1] < data[mid] { //maxima
maximaFound = 1
break
} else if data[mid-1] < data[mid] && data[mid] < data[mid+1] { // increasing
start = mid + 1
} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] { // decreasing
end = mid - 1
} else {
break
}
}
if maximaFound == 0 {
fmt.Println("NoMaximaFound")
return 0, false
}
return mid, true
}
Search element in a bitonic list
A bitonic list comprises of an increasing sequence of integers immediately followed by a decreasing sequence of
integers. To search an element in a bitonic list:
1. Find the index or maximum element in the list. By finding the end of increasing part of the list, using modified
binary search.
2. Once we have the maximum element, search the given value in increasing part of the list using binary search.
3. If the value is not found in increasing part, search the same value in decreasing part of the list using binary search.
Example 5.23
func SearchBitonicArray(data []int, key int) int {
size := len(data)
maxIndex, _ := FindMaxBitonicArray(data)
k := BinarySearch(data, 0, maxIndex, key, true)
if k != -1 {
return k
} else {
return BinarySearch(data, maxIndex+1, size-1, key, false)
}
}
func FindMaxBitonicArray(data []int) (int, bool) {
size := len(data)
start := 0
end := size - 1
mid := 0
if size < 3 {
fmt.Println("InvalidInput")
return -1, false
}
for start <= end {
mid = (start + end) / 2
if data[mid-1] < data[mid] && data[mid+1] < data[mid] /* maxima */ {
return mid, true
} else if data[mid-1] < data[mid] && data[mid] < data[mid+1] /* increasing */ {
start = mid + 1
} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] /* decreasing */ {
end = mid - 1
} else {
break
}
}
fmt.Println("error")
return -1, false
}
func BinarySearch(data []int, start int, end int, key int, isInc bool) int {
if end < start {
return -1
}
mid := (start + end) / 2
if key == data[mid] {
return mid
}
if isInc != false && key < data[mid] || isInc == false && key > data[mid] {
return BinarySearch(data, start, mid-1, key, isInc)
} else {
return BinarySearch(data, mid+1, end, key, isInc)
}
}
Occurrence counts in sorted List
Given a sorted list arr[] find the number of occurrences of a number.
First approach: Brute force, Traverse the list and in linear time we will get the occurrence count of the number. This is
done using one loop.
The Time Complexity is O(n) and Space Complexity is O(1)
Example 5.24
func findKeyCount(data []int, key int) int {
count := 0
size := len(data)
for i := 0; i < size; i++ {
if data[i] == key {
count++
}
}
return count
}
Second approach: Since we have sorted list, we should think about some binary search.
1. First, we should find the first occurrence of the key.
2. Then we should find the last occurrence of the key.
3. Take the difference of these two values and you will have the solution.
Example 5.25
func findKeyCount2(data []int, key int) int {
size := len(data)
firstIndex := findFirstIndex(data, 0, size-1, key)
lastIndex := findLastIndex(data, 0, size-1, key)
return (lastIndex - firstIndex + 1)
}
func findFirstIndex(data []int, start int, end int, key int) int {
if end < start {
return -1
}
mid := (start + end) / 2
if key == data[mid] && (mid == start || data[mid-1] != key) {
return mid
}
if key <= data[mid] {
return findFirstIndex(data, start, mid-1, key)
}
return findFirstIndex(data, mid+1, end, key)
}
func findLastIndex(data []int, start int, end int, key int) int {
if end < start {
return -1
}
mid := (start + end) / 2
if key == data[mid] && (mid == end || data[mid+1] != key) {
return mid
}
if key < data[mid] {
return findLastIndex(data, start, mid-1, key)
}
return findLastIndex(data, mid+1, end, key)
}
Separate even and odd numbers in List
Given a list of even and odd numbers, write a program to separate even numbers from the odd numbers.
First approach: allocate a separate list, then scan through the given list, and fill even numbers from the start and odd
numbers from the end.
Second approach: Algorithm is as follows.
1. Initialize the two variable left and right. Variable left=0 and right= size-1.
2. Keep increasing the left index until the element at that index is even.
3. Keep decreasing the right index until the element at that index is odd.
4. Swap the number at left and right index.
5. Repeat steps 2 to 4 until left is less than right.
Example 5.26
func seperateEvenAndOdd(data []int) {
size := len(data)
left := 0
right := size - 1
for left < right {
if data[left]%2 == 0 {
left++
} else if data[right]%2 == 1 {
right--
} else {
data[left], data[right] = data[right], data[left] // swap
left++
right--
}
}
}
Stock purchase-sell problem
Given a list, in which nth element is the price of the stock on nth day. You are asked to buy once and sell once, on what
date you will be buying and at what date you will be selling to get maximum profit.
Or
Given a list of numbers, you need to maximize the difference between two numbers, such that you can subtract the
number, which appear before form the number that appear after it.
First approach: Brute force, for each element in list find if there is some other element whose difference is maximum.
This is done using two for loop, first loop to select, buy date index and the second loop to find its selling date entry.
The Time Complexity is O(n2) and Space Complexity is O(1)
Second approach: Another clever solution is to keep track of the smallest value seen so far from the start. At each point,
we can find the difference and keep track of the maximum profit. This is a linear solution.
The Time Complexity of the algorithm is O(n) time. Space Complexity for creating count list is also O(1)
Example 5.27
func maxProfit(stocks []int) {
size := len(stocks)
buy := 0
sell := 0
curMin := 0
currProfit := 0
maxProfit := 0
for i := 0; i < size; i++ {
if stocks[i] < stocks[curMin] {
curMin = i
}
currProfit = stocks[i] - stocks[curMin]
if currProfit > maxProfit {
buy = curMin
sell = i
maxProfit = currProfit
}
}
fmt.Println("Purchase day is- ", buy, " at price ", stocks[buy])
fmt.Println("Sell day is- ", sell, " at price ", stocks[sell])
fmt.Println("Max Profit :: ", maxProfit)
}
Find a median of a list
Given a list of numbers of size n, if all the elements of the list are sorted then find the element, which lie at the index
n/2.
First approach: Sort the list and return the element in the middle.
Sorting algorithms take O(n.logn).
The Time Complexity of an algorithm is O(n.logn) and Space Complexity is O(1)
Example 5.28
func getMedian(data []int) int {
size := len(data)
sort.Ints(data)
return data[size/2]
}
Second approach: Use QuickSelect algorithm. This algorithm we will look into the next chapter. In QuickSort
algorithm just skip the recursive call that we do not need.
The average Time Complexity of this algorithm will be O(1)
Find median of two sorted Lists.
First approach: Keep track of the index of both the list, say the index are i and j. keep increasing the index of the list
which ever have a smaller value. Use a counter to keep track of the elements that we have already traced.
The Time Complexity of an algorithm is O(n) and Space Complexity is O(1)
Example 5.29
func findMedian(dataFirst []int, dataSecond []int) int {
sizeFirst := len(dataFirst)
sizeSecond := len(dataSecond)
// cealing function.
medianIndex := ((sizeFirst+sizeSecond)+(sizeFirst+sizeSecond)%2)/2
i := 0
j := 0
count := 0
for count < medianIndex-1 {
if i < sizeFirst-1 && dataFirst[i] < dataSecond[j] {
i++
} else {
j++
}
count++
}
if dataFirst[i] < dataSecond[j] {
return dataFirst[i]
}
return dataSecond[j]
}
Search 01 List
Given a list of 0’s and 1’s. All the 0’s come before 1’s. Write an algorithm to find the index of the first 1.
Or
You are given a list which contains either 0 or 1, and they are in sorted order Ex. a [] = {1, 1, 1, 1, 0, 0, 0} How will
you count no of 1`s and 0's?
First approach: Binary Search, since the list is sorted using binary search to find the desired index.
The Time Complexity of an algorithm is O(logn) and Space Complexity is O(1)
Example 5.30
func BinarySearch01(data []int) int {
size := len(data)
if size == 1 && data[0] == 1 {
return -1
}
return BinarySearch01Util(data, 0, size-1)
} func BinarySearch01Util(data []int, start int, end int) int {
if end < start {
return -1
}
mid := (start + end) / 2
if 1 == data[mid] && 0 == data[mid-1] {
return mid
}
if 0 == data[mid] {
return BinarySearch01Util(data, mid+1, end)
}
return BinarySearch01Util(data, start, mid-1)
}
Search in sorted rotated List
Given a sorted list s of n integer. s is rotated an unknown number of times. Find an element in the list.
First approach: Since the list is sorted, we can use modified binary search to find the element.
The Time Complexity of an algorithm is O(logn) and Space Complexity is O(1)
Example 5.31
func BinarySearchRotateArray(data []int, key int) int {
size := len(data)
return BinarySearchRotateArrayUtil(data, 0, size-1, key)
}
func BinarySearchRotateArrayUtil(data []int, start int, end int, key int) int {
if end < start {
return -1
}
mid := (start + end) / 2
if key == data[mid] {
return mid
}
if data[mid] > data[start] {
if data[start] <= key && key < data[mid] {
return BinarySearchRotateArrayUtil(data, start, mid-1, key)
}
return BinarySearchRotateArrayUtil(data, mid+1, end, key)
} else {
if data[mid] < key && key <= data[end] {
return BinarySearchRotateArrayUtil(data, mid+1, end, key)
}
return BinarySearchRotateArrayUtil(data, start, mid-1, key)
}
}
First Repeated element in the list
Given an unsorted list of n elements, find the first element, which is repeated.
First approach: Exhaustive search or Brute force, for each element in list find if there is some other element with the
same value. This is done using two for loop, first loop to select the element and second loop to find its duplicate entry.
The Time Complexity is O(n2) and Space Complexity is O(1)
Example 5.32
func FirstRepeated(data []int) int {
size := len(data)
for i := 0; i < size; i++ {
for j := i + 1; j < size; j++ {
if data[i] == data[j] {
return data[i]
}
}
}
return 0
}
Second approach: Hash-Table, using Hash-Table, we can keep track of the number of times a particular element came
in the list. First scan just populate the Hashtable. In the second, scan just look the occurrence of the elements in the
Hashtable. If occurrence is more for some element, then we have our solution and the first repeated element. Hash-
Table insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time. Space
Complexity is also O(n) for maintaining hash.
Transform List
How would you swap elements of a list like [a1 a2 a3 a4 b1 b2 b3 b4] to convert it into [a1 b1 a2 b2 a3 b3 a4 b4]?
Approach:
· First swap elements in the middle pair
· Next swap elements in the middle two pairs
· Next swap elements in the middle three pairs
· Iterate n-1 steps.
Example 5.33
func transformArrayAB1(str string) string {
data := []rune(str)
size := len(data)
N := size / 2
for i := 1; i < N; i++ {
for j := 0; j < i; j++ {
data[N-i+2*j], data[N-i+2*j+1] = data[N-i+2*j+1], data[N-i+2*j]
}
}
return string(data)
}
Find 2nd largest number in a list with minimum comparisons
Suppose you are given an unsorted list of n distinct elements. How will you identify the second largest element with
minimum number of comparisons?
First approach: Find the largest element in the list. Then replace the last element with the largest element. Then search
the second largest element int the remaining n-1 elements.
The total number of comparisons is: (n-1) + (n-2)
Second approach: Sort the list and then give the (n-1) element. This approach is still more inefficient.
Third approach: Using priority queue / Heap. This approach we will look into heap chapter. Use buildHeap() function
to build heap from the list. This is done in n comparisons. Arr[0] is the largest number, and the greater among arr[1]
and arr[2] is the second largest.
The total number of comparisons are: (n-1) + 1 = n
Check if two Lists are permutation of each other
Given two integer Lists. You have to check whether they are permutation of each other.
First approach: Sorting, Sort all the elements of both the Lists and Compare each element of both the Lists from
beginning to end. If there is no mismatch, return true. Otherwise, false.
Sorting algorithms take O(n.logn) time and comparison take O(n) time.
The Time Complexity of an algorithm is O(n.logn) and Space Complexity is O(1)
Example 5.34
func CheckPermutation(data1 []int, data2 []int) bool {
size1 := len(data1)
size2 := len(data2)
if size1 != size2 {
return false
}
sort.Ints(data1)
sort.Ints(data2)
for i := 0; i < size1; i++ {
if data1[i] != data2[i] {
return false
}
}
return true
}
Second approach: Hash-Table (Assumption: No duplicates).
1. Create a Hash-Table for all the elements of the first list.
2. Traverse the other list from beginning to the end and search for each element in the Hash-Table.
3. If all the elements are found in, the Hash-Table return true, otherwise return false.
Hash-Table insert and find take constant time O(1) so the total Time Complexity of the algorithm is O(n) time. Space
Complexity is also O(n)
Time Complexity = O(n) (For creation of Hash-Table and look-up),
Space Complexity = O(n) (For creation of Hash-Table).
Example 5.35
func checkPermutation2(data1 []int, data2 []int) bool {
size1 := len(data1)
size2 := len(data2)
h := make(map[int]int)
if size1 != size2 {
return false
}
for i := 0; i < size1; i++ {
h[data1[i]]++
h[data2[i]]--
}
for i := 0; i < size1; i++ {
if h[data1[i]] != 0 {
return false
}
}
return true
}
Remove duplicates in an integer list
First approach: Sorting, Steps are as follows:
1. Sort the list.
2. Take two references. A subarray will be created with all unique elements starting from 0 to the first reference
(The first reference points to the last index of the subarray). The second reference iterates through the list
from 1 to the end. Unique numbers will be copied from the second reference location to first reference location
and the same elements are ignored.
Time Complexity calculation:
Time to sort the list = O(nlogn). Time to remove duplicates = O(n).
Overall Time Complexity = O(nlogn).
No additional space is required so Space Complexity is O(1).
Example 5.36
func removeDuplicates(data []int) int {
j := 0
size := len(data)
if size == 0 {
return 0
}
sort.Ints(data)
for i := 1; i < size; i++ {
if data[i] != data[j] {
j++
data[j] = data[i]
}
}
return j + 1
}
Searching for an element in a 2-d sorted list
Given a 2 dimensional list. Each row and column are sorted in ascending order. How would you find an element in it?
The algorithm works as:
1. Start with element at last column and first row
2. If the element is the value we are looking for, return true.
3. If the element is greater than the value we are looking for, go to the element at previous column but same row.
4. If the element is less than the value we are looking for, go to the element at next row but same column.
5. Return false, if the element is not found after reaching the element of the last row of the first column. Condition
(row < r && column >= 0) is false.
Example 5.37
func FindElementIn2DArray(data [][]int, r int, c int, value int) bool {
row := 0
column := c - 1
for row < r && column >= 0 {
if data[row][column] == value {
return true
} else if data[row][column] > value {
column--
} else {
row++
}
}
return false
}
Running time = O(N).
Exercise
1. Given a list of n elements, find the first repeated element.
2. Given a list of n elements, write an algorithm to find three elements in a list whose sum is a given value.
Hint: Try to do this problem using a brute force approach. Then try to apply the sorting approach along with a
brute force approach. The Time Complexity will be O(n2)
3. Given a list of –ve and +ve numbers, write a program to separate –ve numbers from the +ve numbers.
4. Given a list of 1’s and 0’s, write a program to separate 0’s from 1’s.
Hint: QuickSelect, counting
5. Given a list of 0’s, 1’s and 2’s, write a program to separate 0’s , 1’s and 2’s.
6. Given a list whose elements is monotonically increasing with both negative and positive numbers. Write an
algorithm to find the point at which list becomes positive.
7. Given a sorted list, find a given number. If found return the index if not, find the index of that number if it is
inserted into the list.
8. Find max in sorted rotated list.
9. Find min in the sorted rotated list.
10. Find kth Smallest Element in the Union of Two Sorted Lists