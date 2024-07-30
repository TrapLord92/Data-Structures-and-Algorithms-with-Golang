package main

/*Set implementation of Go Collections
Set is used to store only unique elements. Set is implemented using a hash table so elements are not stored in sequential
order.
Example 4.4: Set using go collection./**/
import (
	"fmt"

	"github.com/golang-collections/collections/set"

)
func main1() {
st := set.New()
st.Insert(1)
st.Insert(2)
fmt.Println(st.Has(1))
fmt.Println(st.Has(3))
}
  
//set using map interface

type Set map[interface{}]bool
func (s *Set) Add(key interface{}) {
(*s)[key] = true
}
func (s *Set) Remove(key interface{}) {
delete((*s), key)
}
func (s *Set) Find(key interface{}) bool {
return (*s)[key]
}
func main() {
	mp := make(Set)
	mp.Add("a")
	mp.Add("b")
	mp.Add("a")
	fmt.Println(mp.Find("a"))
	fmt.Println(mp.Find("b"))
	fmt.Println(mp.Find("c"))
	}