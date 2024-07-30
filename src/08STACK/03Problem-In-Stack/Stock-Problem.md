Stock Span Problem
Given a list of daily stock price in a list A[i]. Find the span of the stocks for each day. A span of stock is the maximum
number of days for which the price of stock was lower than that day.
Example 8.8: Approach 1
func StockSpanRange(arr []int) []int {
n := len(arr)
SR := make([]int, n)
SR[0] = 1
for i := 1; i < len(arr); i++ {
SR[i] = 1
for j := i - 1; (j >= 0) && (arr[i] >= arr[j]); j-- {
SR[i]++
}
}
return SR
}
Example 8.9: Approach 2
func StockSpanRange2(arr []int) []int {
stk := new(Stack)
n := len(arr)
SR := make([]int, n)
stk.Push(0)
SR[0] = 1
for i := 1; i < len(arr); i++ {
for !stk.IsEmpty() && arr[stk.Top().(int)] <= arr[i] {
stk.Pop()
}
if stk.IsEmpty() {
SR[i] = (i + 1)
} else {
SR[i] = (i - stk.Top().(int))
}
stk.Push(i)
}
return SR
}
Get Max Rectangular Area in a Histogram
Given a histogram of rectangle bars of each one unit wide. Find the maximum area rectangle in the histogram.
Example 8.10: Approach 1
func GetMaxArea(arr []int) int {
size := len(arr)
maxArea := -1
currArea := 0
minHeight := 0
for i := 1; i < size; i++ {
minHeight = arr[i]
for j := i - 1; j >= 0; j-- {
if minHeight > arr[j] {
minHeight = arr[j]
}
currArea = minHeight * (i - j + 1)
if maxArea < currArea {
maxArea = currArea
}
}
}
return maxArea
}
Approach 2: Divide and conquer
Example 8.11: Approach 3
func GetMaxArea2(arr []int) int {
size := len(arr)
stk := new(Stack)
maxArea := 0
Top := 0
TopArea := 0
i := 0
for i < size {
for (i < size) && (stk.IsEmpty() || arr[stk.Top().(int)] <= arr[i]) {
stk.Push(i)
i++
}
for !stk.IsEmpty() && (i == size || arr[stk.Top().(int)] > arr[i]) {
Top = stk.Top().(int)
stk.Pop()
if stk.IsEmpty() {
TopArea = arr[Top] * i
} else {
TopArea = arr[Top] * (i - stk.Top().(int) - 1)
}
if maxArea < TopArea {
maxArea = TopArea
}
}
}
return maxArea
}
Uses of Stack
· Recursion can also be done using stack. (In place of the system stack)
· The function call is implemented using stack.
· Some problems when we want to reverse a sequence, we just push everything in stack and pop from it.
· Grammar checking, balance parenthesis, infix to postfix conversion, postfix evaluation


Exercise
1. Converting Decimal Numbers to Binary Numbers using stack data structure.
Hint: store reminders into the stack and then print the stack.
2. Convert an infix expression to prefix expression.
3. Write an HTML opening tag and closing tag-matching program.
Hint: parenthesis matching.
4. Write a function that will do Postfix to Infix Conversion
5. Write a function that will do Prefix to Infix Conversion
6. Write a palindrome matching function, which ignore characters other than English alphabet and digits. String
"Madam, I'm Adam." should return true.
7. In the Growing-Reducing Stack implementation using list. Try to figure out a better algorithm which will work
similar to Vector<> or ArrayDeque<>.