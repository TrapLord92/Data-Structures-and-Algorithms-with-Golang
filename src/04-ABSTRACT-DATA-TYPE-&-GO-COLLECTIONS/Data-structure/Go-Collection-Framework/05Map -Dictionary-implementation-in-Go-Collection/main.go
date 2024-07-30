package main

import "fmt"

/*Map / Dictionary implementation in Go Collection
A Dictionary is a data structure that maps keys to values. A Dictionary is using a hash table so the key value pairs are
not stored in sorted order. Dictionary does not allow duplicate keys buts values can be duplicate.*/

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50
	for key, val := range m {
		fmt.Print("[ ", key, " -> ", val, " ]")
	}
	fmt.Println("Apple price:", m["Apple"])
	delete(m, "Apple")
	fmt.Println("Apple price:", m["Apple"])
	v, ok := m["Apple"]
	fmt.Println("Apple price:", v, "Present:", ok)
	v2, ok2 := m["Banana"]
	fmt.Println("Banana price:", v2, "Present:", ok2)
}
