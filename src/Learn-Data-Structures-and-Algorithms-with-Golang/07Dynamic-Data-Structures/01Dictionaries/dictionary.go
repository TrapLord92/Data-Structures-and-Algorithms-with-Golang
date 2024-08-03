// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and sync package
import (
	"fmt"
	"sync"
)

/*Dictionaries
A dictionary is a collection of unique key and value pairs. A dictionary is a broadly useful
data structure for storing a set of data items. It has a key, and each key has a solitary item
associated with it. When given a key, the dictionary will restore the item associated with
that key. These keys can be of any type: strings, integers, or objects. Where we need to sort
a list, an element value can be retrieved utilizing its key. Add, remove, modify, and lookup
operations are allowed in this collection. A dictionary is similar to other data structures,
such as hash, map, and HashMap. The key/value store is used in distributed caching and in
memory databases. Arrays differ from dictionaries in how the data is accessed. A set has
unique items, whereas a dictionary can have duplicate values.
Dictionary data structures are used in the following streams:
Phone directories
Router tables in networking
Page tables in operating systems
Symbol tables in compilers
Genome maps in biology
The following code shows how to initialize and modify a dictionary. In this snippet, the
dictionary has the key DictKey and is a string:*/

// DictKey is the key of the dictionary

/*
DictVal type
The dictionary has the value DictVal of type string mapped to DictKey:
*/
type DictKey string

// DictVal type
type DictVal string

// Dictionary class
/*Dictionary class
The dictionary in the following code is a class with dictionary elements, with DictKey
as the key and DictVal as the value. It has a sync.RWMutex property, lock:*/

type Dictionary struct {
	elements map[DictKey]DictVal
	lock     sync.RWMutex
}

// Put method
/*Put method
A has a Put method, as shown in the following example, that takes the key and value
parameters of the DictKey and DictVal types respectively. The Lock method of the
dictionary's lock instance is invoked, and the Unlock method is deferred. If there are
empty map elements in the dictionary, elements are initialized using make.
The map elements are set with a key and a value if they are not empty:*/
func (dict *Dictionary) Put(key DictKey, value DictVal) {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	if dict.elements == nil {
		dict.elements = make(map[DictKey]DictVal)
	}
	dict.elements[key] = value
}

// Remove method
/*Remove method
A dictionary has a remove method, as shown in the following code, which has a key
parameter of the DictKey type. This method returns a bool value if the value associated
with Dictkey is removed from the map:*/

func (dict *Dictionary) Remove(key DictKey) bool {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	var exists bool
	_, exists = dict.elements[key]
	if exists {
		delete(dict.elements, key)
	}
	return exists
}

// Contains method
/*Contains method
In the following code, the Contains method has an input parameter, key, of the
DictKey type, and returns bool if key exists in the dictionary:*/
func (dict *Dictionary) Contains(key DictKey) bool {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	var exists bool
	_, exists = dict.elements[key]
	return exists
}

// Find method
/*Find method
The Find method takes the key parameter of the DictKey type and returns the
DictVal type associated with the key. The following code snippet explains the Find
method:*/
func (dict *Dictionary) Find(key DictKey) DictVal {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return dict.elements[key]
}

// Reset method
/*Reset method
The Reset method of the Dictionary class is presented in the following snippet. The
Lock method of the dictionary's lock instance is invoked and Unlock is deferred. The
elements map is initialized with a map of the DictKey key and the DictVal value:*/
func (dict *Dictionary) Reset() {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	dict.elements = make(map[DictKey]DictVal)
}

// NumberOfElements method
/*NumberOfElements method
The NumberOfElements method of the Dictionary class returns the length of the
elements map. The RLock method of the lock instance is invoked. The RUnlock method
of the lock instance is deferred before returning the length; this is shown in the following
code snippet:*/
func (dict *Dictionary) NumberOfElements() int {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return len(dict.elements)
}

// GetKeys method
/*GetKeys method
The GetKeys method of the Dictionary class is shown in the following code snippet. The
method returns the array of the DictKey elements. The RLock method of the lock instance
is invoked, and the RUnlock method is deferred. The dictionary keys are returned by
traversing the element's map:*/
func (dict *Dictionary) GetKeys() []DictKey {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	var dictKeys []DictKey
	dictKeys = []DictKey{}
	var key DictKey
	for key = range dict.elements {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

// GetValues method
/*GetValues method
The GetValues method of the Dictionary class returns the array of the DictVal
elements. In the following code snippet, the RLock method of the lock instance is invoked
and the RUnlock method is deferred. The array of dictionary values is returned after
traversing the element's map:*/
func (dict *Dictionary) GetValues() []DictVal {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	var dictValues []DictVal
	dictValues = []DictVal{}
	var key DictKey
	for key = range dict.elements {
		dictValues = append(dictValues, dict.elements[key])
	}
	return dictValues
}

// main method
/*The main method
The following code shows the main method, where the dictionary is initialized and printed:*/
func main() {

	var dict *Dictionary = &Dictionary{}

	dict.Put("1", "1")
	dict.Put("2", "2")
	dict.Put("3", "3")
	dict.Put("4", "4")

	fmt.Println(dict)

}
