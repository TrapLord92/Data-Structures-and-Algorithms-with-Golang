// main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

/*A queue consists of elements to be processed in a particular order or based on priority. A
priority-based queue of orders is shown in the following code, structured as a heap.
Operations such as enqueue, dequeue, and peek can be performed on queue. A queue is a
linear data structure and a sequential collection. Elements are added to the end and are
removed from the start of the collection. Queues are commonly used for storing tasks that
need to be done, or incoming HTTP requests that need to be processed by a server. In real
life, handling interruptions in real-time systems, call handling, and CPU task scheduling
are good examples for using queues.*/

// Queue  - Array of Orders Type
type Queue []*Order

// Order class
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

/*The New method
The New method on the Order class assigns the properties from the priority, quantity,
and product parameters for name and customerName. The method initializes the
properties of the order as follows:*/
// New method initialises with Order with priority, quantity, product, customerName
func (order *Order) New(priority int, quantity int, product string, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}

// Add method adds the order to the queue
/*The Add method
In the following code snippet, the Add method on the Queue class takes the order
parameter and adds it to Queue based on the priority. Based on this, the location of the
order parameter is found by comparing it with the priority parameter:*/
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {

		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

/*The example output after the add method is invoked with the order parameter is as
follows. The order is checked to see whether or not it exists in the queue. The order is then
appended to the queue:*/

// main method
func main() {
	/*The main method creates two orders, and the priority of the orders is set to 2 and 1. In the
	following code, the queue will first process the order with the higher number on the
	priority value:*/
	var queue Queue
	queue = make(Queue, 0)

	var order1 *Order = &Order{}

	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "Computer"
	var customerName1 string = "Greg White"

	order1.New(priority1, quantity1, product1, customerName1)

	var order2 *Order = &Order{}

	var priority2 int = 1
	var quantity2 int = 10
	var product2 string = "Monitor"
	var customerName2 string = "John Smith"

	order2.New(priority2, quantity2, product2, customerName2)

	queue.Add(order1)

	queue.Add(order2)

	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
