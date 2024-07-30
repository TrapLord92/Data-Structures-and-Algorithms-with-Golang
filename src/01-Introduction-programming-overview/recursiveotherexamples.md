lowest Nth disk from source to destination. Then will move N-1 disks from temp to destination.
Greatest common divisor (GCD)
Example 1.36: Find greatest common divisor.
func GCD(m int, n int) int {
if m < n {
return GCD(n, m)
}
if m%n == 0 {
return n
}
return GCD(n, m%n)
}
Analysis: Euclid’s algorithm is used to find gcd. GCD(n, m) == GCD(m, n mod m).
Fibonacci number
Example 1.37: Given N find the Nth number in the Fibonacci series. .
func fibonacci(n int) int {
if n <= 1 {
return n
}
return fibonacci(n-1) + fibonacci(n-2)
}
Analysis: Fibonacci number are calculated by adding sum of the previous two number.
Note:- There is an inefficiency in the solution we will look better solution in coming chapters.
All permutations of an integer list
Example 1.38: Generate all permutations of an integer list.
func Permutation(data []int, i int, length int) {
if length == i {
PrintSlice(data)
return
}
for j := i; j < length; j++ {
swap(data, i, j)
Permutation(data, i+1, length)
swap(data, i, j)
}
}
func swap(data []int, x int, y int) {
data[x], data[y] = data[y], data[x]
}
func main() {
var data [5]int
for i := 0; i < 5; i++ {
data[i] = i
}
Permutation(data[:], 0, 5)
}
Analysis: In permutation method at each recursive call number at index, “i” is swapped with all the numbers that are
right of it. Since the number is swapped with all the numbers in its right one by one it will produce all the permutation
possible.


Binary search using recursion
Example 1.39: Search a value in an increasing order sorted list of integers.
func BinarySearchRecursive(data []int, low int, high int, value int) int {
mid := low + (high-low)/2 // To afunc the overflow
if data[mid] == value {
return mid
} else if data[mid] < value {
return BinarySearchRecursive(data, mid+1, high, value)
} else {
return BinarySearchRecursive(data, low, mid-1, value)
}
}
Analysis: Similar iterative solution we had already seen. Now let us look into the recursive solution of the same
problem. In this solution, we are diving the search space into half and discarding the rest. This solution is very efficient
as each step we are rejecting half the search space/ list.