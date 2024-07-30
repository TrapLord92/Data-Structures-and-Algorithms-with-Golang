package binarysearch

/*Binary Search
How do we search a word in a dictionary? In general, we go to some approximate page (mostly middle) and start
searching from that point. If we see the word that we are searching is same then we are done with the search. Else, if we
see the page is before the selected pages, then apply the same procedure for the first half otherwise to the second half.
Binary Search also works in the same way. We get to the middle point from the sorted list and start comparing with the
desired value.
Note: Binary search requires the list to be sorted otherwise binary search cannot be applied.*/

// Binary Search Algorithm - Iterative Way
func Binarysearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	mid := 0
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

/*Time Complexity: O(logn). We always take half input and throwing out the other half. So the recurrence relation for
binary search is T(n) = T(n/2) + c. Using master theorem (divide and conquer), we get T(n) = O(logn)
Space Complexity: O(1)*/
/*Example 5.4: Binary search implementation using recursion.
// Binary Search Algorithm - Recursive Way*/
func BinarySearchRecursive(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)/2
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}

/*Time Complexity: O(logn). Space Complexity: O(logn) for system stack in recursion*/
