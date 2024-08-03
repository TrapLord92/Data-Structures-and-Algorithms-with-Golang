// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and sort package

import (
	"fmt"
	"sort"
)

/*The struct type
A struct type (class) can be sorted using different sets of multiple fields. In the code, we show how to sort struct types. A class called Commit
consists of the username, lang, and numlines properties. username is a string, lang is a
string, and numlines is an integer. In the following code, the Commit class is sorted by
commits and lines:*/
// A Commit is a record of code checkin
type Commit struct {
	username string
	lang     string
	numlines int
}

type lessFunc func(p1 *Commit, p2 *Commit) bool

// multiSorter class
/*The multiSorter class
The multiSorter class consists of the commits and lessFunction array properties. The
multiSorter class implements the Sort interface to sort the commits, as shown in the
following code:*/
type multiSorter struct {
	Commits      []Commit
	lessFunction []lessFunc
}

// Sort Method
/*The Sort method
In the following code snippet, the Sort method of multiSorter sorts the Commits array
by invoking sort.Sort and passing the multiSorter argument:*/
func (multiSorter *multiSorter) Sort(Commits []Commit) {
	multiSorter.Commits = Commits
	sort.Sort(multiSorter)
}

// OrderedBy Method
/*The OrderBy method
The OrderedBy method takes the less function and returns multiSorter. The
multisorter instance is initialized by the less function, as shown in the following code
snippet:*/
func OrderedBy(lessFunction ...lessFunc) *multiSorter {
	return &multiSorter{
		lessFunction: lessFunction,
	}
}

// Len Method
/*The len method
The len method of the multiSorter class returns the length of the Commits array. The
Commits array is a property of multiSorter:*/
func (multiSorter *multiSorter) Len() int {
	return len(multiSorter.Commits)
}

// Swap method
/*The Swap method
The Swap method of multiSorter takes the integers i and j as input. This method swaps
the array elements at index i and j:*/
func (multiSorter *multiSorter) Swap(i int, j int) {
	multiSorter.Commits[i] = multiSorter.Commits[j]
	multiSorter.Commits[j] = multiSorter.Commits[i]
}

// Less method
/*The less method
The Less method of the multiSorter class takes the integers i and j and compares the
element at index i to the element at index j:*/

func (multiSorter *multiSorter) Less(i int, j int) bool {

	var p *Commit
	var q *Commit
	p = &multiSorter.Commits[i]
	q = &multiSorter.Commits[j]

	var k int
	for k = 0; k < len(multiSorter.lessFunction)-1; k++ {
		less := multiSorter.lessFunction[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return multiSorter.lessFunction[k](p, q)
}

// main method
/*The main method
The main method creates a Commit array and initializes the array with values. Functions
are created for sorting by user, language, and lines. OrderedBy returns a
multiSorter, and its sort method is called by user, language, increasingLines, and
decreasingLines:*/
func main() {

	var Commits = []Commit{
		{"james", "Javascript", 110},
		{"ritchie", "python", 250},
		{"fletcher", "Go", 300},
		{"ray", "Go", 400},
		{"john", "Go", 500},
		{"will", "Go", 600},
		{"dan", "C++", 500},
		{"sam", "Java", 650},
		{"hayvard", "Smalltalk", 180},
	}

	var user func(*Commit, *Commit) bool
	user = func(c1 *Commit, c2 *Commit) bool {
		return c1.username < c2.username
	}

	var language func(*Commit, *Commit) bool
	language = func(c1 *Commit, c2 *Commit) bool {
		return c1.lang < c2.lang
	}

	var increasingLines func(*Commit, *Commit) bool
	increasingLines = func(c1 *Commit, c2 *Commit) bool {
		return c1.numlines < c2.numlines
	}

	var decreasingLines func(*Commit, *Commit) bool
	decreasingLines = func(c1 *Commit, c2 *Commit) bool {
		return c1.numlines > c2.numlines // Note: > orders downwards.
	}

	OrderedBy(user).Sort(Commits)
	fmt.Println("By username:", Commits)

	OrderedBy(user, increasingLines).Sort(Commits)
	fmt.Println("By username,asc order", Commits)

	OrderedBy(user, decreasingLines).Sort(Commits)
	fmt.Println("By username,desc order", Commits)

	OrderedBy(language, increasingLines).Sort(Commits)
	fmt.Println("By lang,asc order", Commits)

	OrderedBy(language, decreasingLines, user).Sort(Commits)
	fmt.Println("By lang,desc order", Commits)

}
