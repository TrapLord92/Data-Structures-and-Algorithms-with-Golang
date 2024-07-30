Min stack
Design a stack in which get minimum value in stack should also work in O(1) Time Complexity.
Hint: Keep two stack one will be general stack, which will just keep the elements. The second will keep the min value.
1. Push: Push an element to the top of stack1. Compare the new value with the value at the top of the stack2. If the
new value is smaller, then push the new value into stack2. Or push the value at the top of the stack2 to itself once
more.
2. Pop: Pop an element from top of stack1 and return. Pop an element from top of stack2 too.
3. Min: Read from the top of the stack2, this value will be the min.
Palindrome string
Find if given string is a palindrome or not using a stack.
Definition of palindrome: A palindrome is a sequence of characters that is same backward or forward.
Eg. “AAABBBCCCBBBAAA”, “ABA” & “ABBA”
Hint: Push characters to the stack until the half-length of the string. Then pop these characters and then compare. Make
sure you take care of the odd length and even length.
Depth-First Search with a Stack
In a depth-first search, we traverse down a path until we get a dead end; then we backtrack by popping a stack to get an
alternative path.
· Create a stack
· Create a start point
· Push the start point onto the stack
· While (value searching not found and the stack is not empty)
o Pop the stack
o Find all possible points after the one which we just tried
o Push these points onto the stack
Stack using a queue
How to implement a stack using a queue. Analyze the running time of the stack operations.
See queue chapter for this.