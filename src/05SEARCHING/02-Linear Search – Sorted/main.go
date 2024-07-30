package main

/*Linear Search – Sorted
If elements of the list are sorted either in increasing order or in decreasing order, searching for a desired element will be
much more efficient than unordered linear search. In many cases, we do not need to traverse the complete list.
Following example explains when you encounter a greater value element from the increasing sorted list, you stop
searching further. This is how this algorithm saves the time and improves the performance.
Example 5.2*/

func linearSearchSorted(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		} else if value < data[i] {
			return false
		}
	}
	return false
}

/*Time Complexity: O(n). As we need to traverse the complete list in worst case. Worst case is when your desired
element is at the last position of the sorted list. However, in the average case this algorithm is more efficient even
though the growth rate is same as unsorted.
Space Complexity: O(1). No extra memory is used to allocate the list.*/
