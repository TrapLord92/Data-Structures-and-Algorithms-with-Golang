Introduction
A queue is a basic data structure that organized items in first-in-first-out (FIFO) manner. First element inserted into a
queue will be the first to be removed. It is also known as "first-come-first-served".
The real life analogy of queue is typical lines in which we all participate time to time.
· We wait in a line of railway reservation counter.
· We wait in the cafeteria line.
· We wait in a queue when we call to some customer case.
The elements, which are at the front of the queue, are the one that stayed in the queue for the longest time.
Computer science also has common examples of queues. We issue a print command from our office to a single printer
per floor. The print task are lined up in a printer queue. The print command that was issued first will be printed before
the next commands in line.
In addition to printing queues, operating system is also using different queues to control process scheduling. Processes
are added to processing queue, which is used by an operating system for various scheduling algorithms.
Soon we will be reading about graphs and will come to know about breadth-first traversal, which uses a queue.
The Queue Abstract Data Type
Queue abstract data type is defined as a class whose object follows FIFO or first-in-first-out for the elements, added to
it.
Queue should support the following operation:
1. add(): Which add a single element at the back of a queue
2. remove(): Which remove a single element from the front of a queue.
3. isEmpty(): Returns 1 if the queue is empty
4. size(): Returns the number of elements in a queue.