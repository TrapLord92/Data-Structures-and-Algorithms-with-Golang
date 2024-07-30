package main

/*Linear Search – Unsorted Input
When elements of a list are not ordered or sorted and we want to search for a particular value, we need to scan the full
list unless we find the desired value. This kind of algorithm known as unordered linear search. The major problem with
this algorithm is less performance or high Time Complexity in worst case.*/

func linearSearchUnsorted(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

/*Time Complexity: O(n). As we need to traverse the complete list in worst case. Worst case is when your desired
element is at the last position of the list. Here, ‘n’ is the size of the list.
Space Complexity: O(1). No extra memory is used to allocate the list.*/
