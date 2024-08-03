// /main package has examples shown
// in Go Data Structures and algorithms book
package main

/*Ordered lists
Lists in Go can be sorted in two ways:
Ordered list: By creating a group of methods for the slice data type and calling
sort
Unordered list: The other way is to invoke sort.Slice with a custom less
function
The only difference between an ordered list and an unordered list is that, in an ordered list,
the order in which the items are displayed is mandatory.
An ordered list in HTML starts with an <ol> tag. Each item in the list is written in <li>
tags. Here's an example:
<ol>
<li>Stones</li>
<li>Branches</li>
<li>Smoke</li>
</ol>*/

// importing fmt and sort package
import (
	"fmt"
	"sort"
)

/*An example of an ordered list using Golang is shown in the following code snippet. The
Employee class has Name, ID, SSN, and Age properties:*/
// class Employee
type Employee struct {
	Name string
	ID   string
	SSN  int
	Age  int
}

// ToString method
/*The ToString method
The ToString method of the Employee class returns a string version of employee. The
string version consists of a comma-separated Name, Age, ID, and SSN. This is shown in the
following code snippet:*/
func (employee Employee) ToString() string {
	return fmt.Sprintf("%s: %d,%s,%d", employee.Name, employee.Age, employee.ID, employee.SSN)
}

// SortByAge type
/*The SortByAge type
The SortByAge method sorts the elements concerned by Age. The SortByAge interface
operates on the Employee array. This is shown in the following code snippet:*/
type SortByAge []Employee

func (sortIntf SortByAge) Len() int               { return len(sortIntf) }
func (sortIntf SortByAge) Swap(i int, j int)      { sortIntf[i], sortIntf[j] = sortIntf[j], sortIntf[i] }
func (sortIntf SortByAge) Less(i int, j int) bool { return sortIntf[i].Age < sortIntf[j].Age }

// main method
func main() {
	var employees = []Employee{
		{"Graham", "231", 235643, 31},
		{"John", "3434", 245643, 42},
		{"Michael", "8934", 32432, 17},
		{"Jenny", "24334", 32444, 26},
	}

	fmt.Println(employees)

	sort.Sort(SortByAge(employees))
	fmt.Println(employees)

	sort.Slice(employees, func(i int, j int) bool {
		return employees[i].Age > employees[j].Age
	})
	fmt.Println(employees)

}
