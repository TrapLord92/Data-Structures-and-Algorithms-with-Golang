Array Interview Questions
The following section will discuss the various algorithms that are applicable to Lists and will follow by list of practice
problems with similar approaches.
Sum List
Write a method that will return the sum of all the elements of the integer list, given list as an input argument.
Example 1.28:
func SumArray(data []int) int {
size := len(data)
total := 0
for index := 0; index < size; index++ {
total = total + data[index]
}
return total
}
Sequential Search
Example 1.29: Write a method, which will search a list for some given value.
func SequentialSearch(data []int, value int) bool {
size := len(data)
for i := 0; i < size; i++ {
if value == data[i] {
return true
}
}
return false
}
Analysis:
· Since we have no idea about the data stored in the list, or if the data is not sorted then we have to search the list in
sequential manner one by one.
· If we find the value, we are looking for we return true.
· Else, we return False in the end, as we did not found the value we are looking for.
In the above example, the data are not sorted. If the data is sorted, a binary search can be used. We examine the middle
position at each step. Depending upon the data that we are searching is greater or smaller than the middle value. We
will search either the left or the right portion of the array. At each step, we are eliminating half of the search space there
by making this algorithm very efficient.
Binary Search
Example 1.30: Binary search in a sorted list.
func BinarySearch(data []int, value int) bool {
size := len(data)
var mid int
low := 0
high := size - 1
for low <= high {
mid = low + (high-low)/2
if data[mid] == value {
return true
} else {
if data[mid] < value {
low = mid + 1
} else {
high = mid - 1
}
}
}
return false
}
Analysis:
· Since we have data sorted in increasing / decreasing order, we can apply more efficient binary search. At each step,
we reduce our search space by half.
· At each step, we compare the middle value with the value we are searching. If mid value is equal to the value we
are searching for then we return the middle index.
· If the value is smaller than the middle value, we search the left half of the list.
· If the value is greater than the middle value then we search the right half of the list.
· If we find the value we are looking for then its index is returned or -1 is returned otherwise.
Rotating a list by K positions.
Given a list, you need to rotate its elements K number of times. For example, a list [10,20,30,40,50,60] rotate by 2
positions to [30,40,50,60,10,20]
Example 1.31:
func RotateArray(data []int, k int) {
n := len(data)
ReverseArray(data, 0, k-1)
ReverseArray(data, k, n-1)
ReverseArray(data, 0, n-1)
}
func ReverseArray(data []int, start int, end int) {
i := start
j := end
for i < j {
data[i], data[j] = data[j], data[i]
i++
j--
}
}
Analysis:
· Rotating list is done in two parts trick. In the first part, we first reverse elements of list first half and then second
half.
1,2,3,4,5,6,7,8,9,10 => 5,6,7,8,9,10,1,2,3,4
1,2,3,4,5,6,7,8,9,10 => 4,3,2,1,10,9,8,7,6,5 => 5,6,7,8,9,10,1,2,3,4
· Then we reverse the whole list there by completing the whole rotation.
Find the largest sum contiguous subarray.
Given a list of positive and negative integers, find a contiguous subarray whose sum (sum of elements) is maximum.
Example 1.32:
func MaxSubArraySum(data []int) int {
size := len(data)
maxSoFar := 0
maxEndingHere := 0
for i := 0; i < size; i++ {
maxEndingHere = maxEndingHere + data[i]
if maxEndingHere < 0 {
maxEndingHere = 0
}
if maxSoFar < maxEndingHere {
maxSoFar = maxEndingHere
}
}
return maxSoFar
}
func main() {
data := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
fmt.Println("Max sub array sum :", MaxSubArraySum(data))
}
Analysis:
· Maximum subarray in a list is found in a single scan. We keep track of global maximum sum so far and the
maximum sum, which include the current element.
· When we find global maximum value so far is less than the maximum value containing current value we update the
global maximum value.
· Finally return the global maximum value.