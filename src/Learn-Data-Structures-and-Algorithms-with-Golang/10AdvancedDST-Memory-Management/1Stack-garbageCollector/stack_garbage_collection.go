// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
	"sync"
)

/*Garbage collection
Garbage collection is a type of programmed memory management in which memory,
currently occupied by objects that will never be used again, is gathered. John McCarthy was
the first person to come up with garbage collection for managing Lisp memory
management. This technique specifies which objects need to be de-allocated, and then
discharges the memory. The strategies that are utilized for garbage collection are stack
allocation and region interference. Sockets, relational database handles, user window
objects, and file resources are not overseen by garbage collectors.
Garbage collection algorithms help reduce dangling pointer defects, double-free defects,
and memory leaks. These algorithms are computing-intensive and cause decreased or
uneven performance. According to Apple, one of the reasons for iOS not having garbage
collection is that garbage collection needs five times the memory to match explicit memory
management. In high-transactional systems, concurrent, incremental, and real-time garbage
collectors help manage memory collection and release.
Garbage collection algorithms depend on various factors:
GC throughput
Heap overhead
Pause times
Pause frequency
Pause distribution
Allocation performance
Compaction
Concurrency
Scaling
Tuning
Warm-up time
Page release
Portability
Compatibility
That's simple, deferred, one-bit, weighted reference counting, mark-and-sweep, and
generational collection algorithms discussed in the following sections.*/

// Reference Counter Class
/*The ReferenceCounter class
The following code snippet shows how references to created objects are maintained in the
stack. */

/*
The ReferenceCounter class has the number of references, including the pool of
references and removed references, as properties:
*/
type ReferenceCounter struct {
	num     *uint32
	pool    *sync.Pool
	removed *uint32
}

// new Reference Counter method
/*The newReferenceCounter method
The newReferenceCounter method initializes a ReferenceCounter instance and returns
a pointer to ReferenceCounter. This is shown in the following code:*/
func newReferenceCounter() *ReferenceCounter {
	return &ReferenceCounter{
		num:     new(uint32),
		pool:    &sync.Pool{},
		removed: new(uint32),
	}
}

// New method of Stack class
func (stack *Stack) New() {
	stack.references = make([]*ReferenceCounter, 0)
}

// Stack class
type Stack struct {
	references []*ReferenceCounter
	Count      int
}

// Push method
func (stack *Stack) Push(reference *ReferenceCounter) {
	stack.references = append(stack.references[:stack.Count], reference)
	stack.Count = stack.Count + 1
}

// Pop method
func (stack *Stack) Pop() *ReferenceCounter {
	if stack.Count == 0 {
		return nil
	}

	var length int = len(stack.references)
	var reference *ReferenceCounter = stack.references[length-1]
	if length > 1 {
		stack.references = stack.references[:length-1]

	} else {
		stack.references = stack.references[0:]

	}
	stack.Count = len(stack.references)
	return reference
}

// main method
func main() {
	var stack *Stack = &Stack{}
	stack.New()
	var reference1 *ReferenceCounter = newReferenceCounter()
	var reference2 *ReferenceCounter = newReferenceCounter()
	var reference3 *ReferenceCounter = newReferenceCounter()
	var reference4 *ReferenceCounter = newReferenceCounter()

	stack.Push(reference1)
	stack.Push(reference2)
	stack.Push(reference3)
	stack.Push(reference4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
