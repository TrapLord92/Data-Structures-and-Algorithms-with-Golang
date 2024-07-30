Dictionary / Symbol Table
A symbol table is a mapping between a string (key) and a value, which can be of any data type. A value can be an
integer such as occurrence count, dictionary meaning of a word and so on.
Binary Search Tree (BST) for Strings
Binary Search Tree (BST) is the simplest way to implement symbol table. Simple string compare function can be used
to compare two strings. If all the keys are random, and the tree is balanced. Then on an average key lookup can be done
in logarithmic time.
Hash-Table
The Hash-Table is another data structure, which can be used for symbol table implementation. Below Hash-Table
diagram, we can see the name of that person is taken as the key, and their meaning is the value of the search. The first
key is converted into a hash code by passing it to appropriate hash function. Inside hash function the size of Hash-Table
is also passed, which is used to find the actual index where values will be stored. Finally, the value that is meaning of
name is stored in the Hash-Table, or you can store a reference to the string which store meaning can be stored into the
Hash-Table.
Hash-Table has an excellent lookup of constant time.
Let us suppose we want to implement autocomplete the box feature of Google search. When you type some string to
search in google search, it proposes some complete string even before you have done typing. BST cannot solve this
problem as related strings can be in both right and left subtree.
The Hash-Table is also not suited for this job. One cannot perform a partial match or range query on a Hash-Table.
Hash function transforms string to a number. Moreover, a good hash function will give a distributed hash code even for
partial string and there is no way to relate two strings in a Hash-Table.
Trie and Ternary Search tree are a special kind of tree, which solves partial match, and range query