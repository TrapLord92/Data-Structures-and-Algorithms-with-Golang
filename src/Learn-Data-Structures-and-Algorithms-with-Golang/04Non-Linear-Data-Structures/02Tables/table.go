// /main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

/*As we already know, tables are used in data management and other areas. A table has a
name and a header with the column names. Let's take a look at the different classes in
tables such as the Table class, the Row class, the Column class, and the PrintTable
method in the following sections.*/

// Table
/*A Table class has an array of rows and column names. The table's Name is a string property
in the struct class, as shown here:*/
type Table struct {
	Rows        []Row
	Name        string
	ColumnNames []string
}

// Row Class
/*The Row class has an array of columns and an Id integer, as shown in the following code.
The Id instance is a unique identifier for a row:*/
type Row struct {
	Columns []Column
	Id      int
}

// Column Class
/*A Column class has an Id integer and a Value string that's identified by a unique
identifier, as presented in the following code snippet:*/
type Column struct {
	Id    int
	Value string
}

// printTable method
/*In the following code snippet, the printTable method prints the rows and columns of a
table. Rows are traversed, and then for every row the columns are printed:*/
func printTable(table Table) {

	var rows []Row = table.Rows
	fmt.Println(table.Name)

	for _, row := range rows {

		var columns []Column = row.Columns

		for i, column := range columns {

			fmt.Println(table.ColumnNames[i], column.Id, column.Value)
		}

	}

}

// main method
/*The main method
In this main method, we will instantiate the classes such as Table, Row, and Column, which
we just took a look at. The main method creates a table and sets the name and column
names. Columns are created with values. The columns are set on the rows after the rows
are created. The table is printed by invoking the printTable method, as shown here:*/

func main() {

	var table Table = Table{}
	table.Name = "Customer"
	table.ColumnNames = []string{"Id", "Name", "SSN"}

	var rows []Row = make([]Row, 2)
	rows[0] = Row{}
	var columns1 []Column = make([]Column, 3)
	columns1[0] = Column{1, "323"}
	columns1[1] = Column{1, "John Smith"}
	columns1[2] = Column{1, "3453223"}
	rows[0].Columns = columns1

	rows[1] = Row{}
	var columns2 []Column = make([]Column, 3)
	columns2[0] = Column{2, "223"}
	columns2[1] = Column{2, "Curran Smith"}
	columns2[2] = Column{2, "3223211"}
	rows[1].Columns = columns2

	table.Rows = rows

	fmt.Println(table)

	printTable(table)

}
