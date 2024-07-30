package main

import (
	"fmt"
	"math"
)

// Time Complexity Examples
// Example 2.1
func fun1(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m += 1
	}
	return m
}

//Time Complexity: O(n)

// Example 2.2
func fun2(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m += 1
		}
	}
	return m
}

//Time Complexity: O(n2)

// Example 2.3
func fun3(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			m += 1
		}
	}
	return m
}

//Time Complexity: O(N+(N-1)+(N-2)+...) == O(N(N+1)/2) == O(n2)

// Example 2.4
func fun4(n int) int {
	m := 0
	i := 1
	for i < n {
		m += 1
		i = i * 2
	}
	return m
}

//Each time problem space is divided into half. Time Complexity: O(log(n))

// Example 2.5
func fun5(n int) int {
	m := 0
	i := n
	for i > 0 {
		m += 1
		i = i / 2
	}
	return m
}

//Same as above each time problem space is divided into half. Time Complexity: O(log(n))

// Example 2.6
func fun6(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				m += 1
			}
		}
	}
	return m
}

/*Outer loop will run for n number of iterations. In each iteration of the outer loop, inner loop will run for n iterations of
their own. Final complexity will be n*n*n.
Time Complexity: O(n3)*/

// Example 2.7
func fun7(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m += 1
		}
	}
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			m += 1
		}
	}
	return m
}

/*These two groups of for loop are in consecutive so their complexity will add up to form the final complexity of the
program. Time Complexity: O(n2) + O(n2) = O(n2)*/

// Example 2.8
func fun8(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		sq := math.Sqrt(float64(n))
		for j := 0; j < int(sq); j++ {
			m += 1
		}
	}
	return m
}

/*Time Complexity: O(n * √n ) = O(n3/2)
Example 2.9*/

func fun9(n int) int {
	m := 0
	for i := n; i > 0; i /= 2 {
		for j := 0; j < i; j++ {
			m += 1
		}
	}
	return m
}

//Each time problem space is divided into half. Time Complexity: O(log(n))

// Example 2.10
func fun10(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			m += 1
		}
	}
	return m
}

//O(N+(N-1)+(N-2)+...) = O(N(N+1)/2) // arithmetic progression. Time //Complexity: O(n2)

// Example 2.11
func fun11(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j + 1; k < n; k++ {
				m += 1
			}
		}
	}
	return m
}

//Time Complexity: O(n3)

// Example 2.12
func fun12(n int) int {
	j := 0
	m := 0
	for i := 0; i < n; i++ {
		for ; j < n; j++ {
			m += 1
		}
	}
	return m
}

/*Think carefully once again before finding a solution, j value is not reset at each iteration.
Time Complexity: O(n)*/

// Example 2.13
func fun13(n int) int {
	m := 0
	for i := 1; i <= n; i *= 2 {
		for j := 0; j <= i; j++ {
			m += 1
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instructions O(n):: ", fun1(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", fun2(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", fun3(100))
	fmt.Println("N = 100, Number of instructions O(log(n)):: ", fun4(100))
	fmt.Println("N = 100, Number of instructions O(log(n)):: ", fun5(100))
	fmt.Println("N = 100, Number of instructions O(n^3):: ", fun6(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", fun7(100))
	fmt.Println("N = 100, Number of instructions O(n^(3/2)):: ", fun8(100))
	fmt.Println("N = 100, Number of instructions O(log(n)):: ", fun9(100))
	fmt.Println("N = 100, Number of instructions O(n^2):: ", fun10(100))
	fmt.Println("N = 100, Number of instructions O(n^3):: ", fun11(100))
	fmt.Println("N = 100, Number of instructions O(n):: ", fun12(100))
	fmt.Println("N = 100, Number of instructions O(n):: ", fun13(100))
}
