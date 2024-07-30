/*Queue
A queue is a First-In-First-Out (FIFO) kind of data structure. The element that is added to the queue first will be the
first to be removed and so on.
Queue ADT Operations:
· Add(): Add a new element to the back of the queue.
· Remove(): remove an element from the front of the queue and return its value.
· Front(): return the value of the element at the front of the queue.
· Size(): returns the number of elements inside the queue.
· IsEmpty(): returns 1 if the queue is empty otherwise return 0
Note: All the above Queue operations are implemented in O(1) Time Complexity.
Queue implementation in Go Collection
Deque is the class implementation of doubly ended queue. If we use append(), popleft() it will behave as a queue.*/
package main

import "github.com/golang-collections/collections/queue"

import "fmt"

func main() {
	q := queue.New()
	//adding elements to the queue
	//enqueue Put an item on the end of a queue
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	for q.Len() != 0 {
		val := q.Dequeue() //Take the next item off the front of the queue
		fmt.Print(val, " ")
	}
}
