Problems in Queue
Queue using a stack
How to implement a queue using a stack. You can use more than one stack.
Solution: We can use two stack to implement queue. We need to simulate first in first our using stack.
1. Enqueue Operation: new elements are added to the top of first stack.
2. Dequeue Operation: elements are popped from the second stack. When second stack is empty then all the elements
of first stack are popped one by one and pushed into second stack.
Example 9.5:
type QueueUsingStack struct {
stk1 Stack
stk2 Stack
}
func (que *QueueUsingStack) Add(value int) {
que.stk1.Push(value)
}
func (que *QueueUsingStack) Remove() int {
var value int
if que.stk2.IsEmpty() == false {
value = que.stk2.Pop().(int)
return value
}
for que.stk1.IsEmpty() == false {
value = que.stk1.Pop().(int)
que.stk2.Push(value)
}
value = que.stk2.Pop().(int)
return value
}
Analysis: All add() happens to stack 1. When remove() is called remove happens from stack 2. When the stack 2 is
empty then stack 1 is popped and pushed into stack 2. This popping from stack 1 and pushing into stack 2 revert the
order of retrieval there by making queue behavior out of two stacks.
Stack using a Queue
Implement stack using a queue.
Solution 1: use two queue
Push: add new elements to queue1.
Pop: while size of queue1 is bigger than 1. Push all items from queue 1 to queue 2 except the last item. Switch the
name of queue 1 and queue 2. Then return the last item.
Push operation is O(1) and Pop operation is O(n)
Solution 2: This same can be done using just one queue.
Push: add the element to queue.
Pop: find the size of queue. If size is zero then return error. Else, if size is positive then remove size- 1 elements from
the queue and again add to the same queue. At last, remove the next element and return it.
Push operation is O(1) and Pop operation is O(n)
Solution 3: In the above solutions the push is efficient and pop is un efficient can we make pop efficient O(1) and push
inefficient O(n)
Push: add new elements to queue2. Then add all the elements of queue 1 to queue 2. Then switch names of queue1 and
queue 2.
Pop: remove from queue1
Reverse a stack
Reverse a stack using a queue
Solution 1:
· Pop all the elements of stack and add them into a queue.
· Then remove all the elements of the queue into stack
· We have the elements of the stack reversed.
Solution 2:
· Since dynamic list or [ ] list is used to implement stack in Go, we can iterate from both the directions of the list and
swap the elements.
Reverse a queue
Reverse a queue-using stack
Solution:
· Dequeue all the elements of the queue into stack ( append to the Go list [ ] )
· Then pop all the elements of stack and add them into a queue. (pop the elements from the list )
· We have the elements of the queue reversed.
Breadth-First Search with a Queue
In breadth-first search, we explore all the nearest nodes first by finding all possible successors and add them to a queue.
· Create a queue
· Create a start point
· Enqueue the start point onto the queue
· while (value searching not found and the queue is not empty)
o Dequeue from the queue
o Find all possible points after the last one tried
o Enqueue these points onto the queue
Josephus problem
There are n people standing in a queue waiting to be executed. The counting begins at the front of the queue. In each
step, k number of people are removed and again added one by one from the queue. Then the next person is executed.
The execution proceeds around the circle until only the last person remains, who is given freedom.
Find that position where you want to stand and gain your freedom.
Solution:
· Just insert integer for 1 to k in a queue. (corresponds to k people)
· Define a Kpop() function such that it will remove and add the queue k-1 times and then remove one more time.
(This man is dead.)
· Repeat second step until size of queue is 1.
· Print the value in the last element. This is the solution.
Exercise
1. Implement queue using dynamic memory allocation. Such that the implementation should follow the following
constraints.
a. The user should use memory allocation from the heap using new operator. In this, you need to take care of
the max value in the queue.
b. Once you are done with the above exercise and you are able to test your queue. Then you can add some
more complexity to your code. In add() function when the queue is full, in place of printing, “Queue is full”
you should allocate more space using new operator.
c. Once you are done with the above exercise. Now in remove function once you are below half of the capacity
of the queue, you need to decrease the size of the queue by half. You should add one more variable "min"
to queue so that you can track what is the original value capacity passed at initialization() function.
Moreover, the capacity of the queue will not go below the value passed in the initialization.
(If you are not able to solve the above exercise, and then have a look into stack chapter, where we have done
similar for stack)
2. Implement the below function for the queue:
a. IsEmpty: This is left as an exercise for the user. Take a variable, which will take care of the size of a queue
if the value of that variable is zero, isEmpty should return 1 (true). If the queue is not empty, then it should
return 0 (false).
b. Size: Use the size variable to be used under size function call. Size() function should return the number of
elements in the queue.
3. Implement stack using a queue. Write a program for this problem. You can use just one queue.
4. Write a program to Reverse a stack using queue
5. Write a program to Reverse a queue using stack
6. Write a program to solve Josephus problem (algo already discussed.). There are n people standing in a queue
waiting to be executed. The counting begins at the front of the queue. In each step, k number of people are removed
and again added one by one from the queue. Then the next person is executed. The elimination proceeds around the
circle until only the last person remains, who is given freedom. Find that position where you want to stand and gain
your freedom.
7. Write a CompStack() function which takes reference to two stack as an argument and return true or false depending
upon whether all the elements of the stack are equal or not. You are given isEqual(int, int) which will compare and
return 1 if both values are equal and 0 if they are different.