package main

import "fmt"

/*A map is a collection of Key-Value pairs. Hash-Table is used to store elements in a Map so it is unordered.
var <variable> map[<key datatype>]<value datatype>
Maps have to be initialized using make() before they can be used.
var <variable> map[<key datatype>]<value datatype> = make(map[<key datatype>]<value datatype>)
or
<variable> := make(map[<key datatype>]<value datatype>)
Various operation on map:
1. Assignment: < variable>[<key>] = <value>
2. Delete: delete(<variable >, <key>)
3. Access: value, ok = < variable >[<key>] , the first value will have the value of key in the map. If the key is not
present in the map, it will return zero value corresponding to the value data-type. The second argument returns
whether the map contains the key.*/

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50
	for key, val := range m {
		fmt.Print("[ ", key, " -> ", val, " ]")
	}

	for _, val := range m {
		fmt.Print("[  ", val, " ]")
	}
	fmt.Println("Apple price:", m["Apple"])
	delete(m, "Apple")
	value, ok := m["Apple"]
	fmt.Println("Apple price:", value, "Present:", ok)
	value2, ok2 := m["Banana"]
	fmt.Println("Banana price:", value2, "Present:", ok2)
}
