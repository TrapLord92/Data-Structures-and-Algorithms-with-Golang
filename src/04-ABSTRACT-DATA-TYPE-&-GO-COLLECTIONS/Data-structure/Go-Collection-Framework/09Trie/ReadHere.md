Trie

Trie is a tree, in which we store only one character at each node. This final key value pair is stored in the leaves. Each
node has K children, one for each possible character. For simplicity purpose, let us consider that the character set is 26,
corresponds to different characters of English alphabets.
Trie is an efficient data structure. Using Trie, we can search the key in O(M) time. Where M is the maximum string
length. Trie is suitable for solving partial match and range query problems.

Ternary Search Trie/ Ternary Search Tree

Tries having a very good search performance of O(M) where M is the maximum size of the search string. However,
tries having very high space requirement. Every node Trie contain references to multiple nodes, each reference
corresponds to possible characters of the key. To avoid this high space requirement Ternary Search Trie (TST) is
used. A TST avoid the heavy space requirement of the traditional Trie while still keeping many of its advantages. In a
TST, each node contains a character, an end of key indicator, and three references. The three references are
corresponding to current char hold by the node (equal), characters less than and character greater than.
The Time Complexity of ternary search tree operation is proportional to the height of the ternary search tree. In the
worst case, we need to traverse up to 3 times that many links. However, this case is rare. Therefore, TST is a very good
solution for implementing Symbol Table, Partial match and range query.