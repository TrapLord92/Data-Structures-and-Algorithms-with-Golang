package main

import "fmt"

/*Maps
Maps are used to keep track of keys that are types, such as integers, strings, float, double,
pointers, interfaces, structs, and arrays. The values can be of different types. In the
following example, the language of the map type with a key integer and a value string is*/

func main() {
	var languages = map[int]string{
		3: "English",
		4: "French",
		5: "Spanish",
	}
	// Maps can be created using the make method, specifying the key type and the value type.
	// Products of a map type with a key integer and value string are shown in the following code
	// snippet:
	var products = make(map[int]string)
	products[1] = "chair"
	products[2] = "table"
	// A for loop is used for iterating through the map. The languages map is iterated as follows:
	var i int
	var value string
	for i, value = range languages {
		fmt.Println("language", i, ":", value)
	}
	fmt.Println("product with key 2", products[2])
	// Retrieving value and deleting slice operations using the products map is shown in the following code:
	fmt.Println(products[2])
	delete(products, 1)
	fmt.Println("products", products)
}
