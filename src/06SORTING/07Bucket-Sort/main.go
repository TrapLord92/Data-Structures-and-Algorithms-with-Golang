package main

import "fmt"

/*Bucket Sort
Bucket sort is the simplest and most efficient type of sorting. Bucket sort has a strict requirement of a predefined range
of data.
Like, sort how many people are in which age group. We know that the age of people can vary between 0 and 130.*/

func BucketSort(data []int, lowerRange int, upperRange int) {
	rng := upperRange - lowerRange
	size := len(data)
	count := make([]int, rng)
	for i := 0; i < rng; i++ {
		count[i] = 0
	}
	for i := 0; i < size; i++ {
		count[data[i]-lowerRange]++
	}
	j := 0
	for i := 0; i < rng; i++ {
		for ; count[i] > 0; count[i]-- {
			data[j] = i + lowerRange
			j++
		}
	}
}
func main() {
	data := []int{23, 24, 22, 21, 26, 25, 27, 28, 21, 21}
	BucketSort(data, 20, 30)
	fmt.Println(data)
}

/*Analysis:
· We have created a count list to store counts.
· Count list elements are initialized to zero.
· Index corresponding to input list is incremented.
· Finally, the information stored in count list is saved in the list.
Complexity Analysis:
Data structure List
Worst case performance O(n+k)
Average case performance O(n+k)
Worst case Space Complexity O(k)
k - number of distinct elements.
n - total number of elements in list.*/
