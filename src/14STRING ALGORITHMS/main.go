package main

//TODO
/*Introduction
  A String is a sequence of Unicode character. String is an immutable type variable. Strings are declared using double
  quotes.
  s := "hello, World!"
  Strings are immutable so you cannot change its content once created. You need to first convert into a slice of rune then
  do the changes and in the end convert it back to string.
  Example 14.1:
  func main() {
  s := "hello, World!"
  r := []rune(s)
  r[0] = 'H'
  s2 := string(r)
  fmt.Println(s2)
  }
  Below is table, which explain some of the operations on string.
  mystring := “Hello World!”
  Expression Value Explanation
  len(mystring) 12 Used to find the number of characters in mystring
  “hello”+“world” “helloworld” Two strings are concatenated into a single string
  “world” == “hello”
  “world” == “world”
  False
  True Equality can be tested using “==” sign
  “a” < “b”
  “b” < “a”
  True
  False Unicode value can also be used to compare
  mystring[0] “h” Indexing: String are indexed same as array.
  mystring[1:4] "ell" Slicing
  String Matching
  Every word processing program has a search function in which you can search all occurrences of any particular word in
  a long text file. For this, we need string-matching algorithms.
  Brute Force Search
  We have a pattern that we want to search in the text. The pattern is of length m and the text is of length n. Where m < n.
  The brute force search algorithm will check the pattern at all possible value of “i” in the text where the value of “i”
  range from 0 to n-m. The pattern is compared with the text, character by character from left to right. When a mismatch
  is detected, then pattern is compared by shifting the compare window by one character.
  Example 14.2:
  func BruteForceSearch(text string, pattern string) int {
  j := 0
  n := len(text)
  m := len(pattern)
  for i := 0; i <= n-m; i++ {
  j = 0
  for j < m && pattern[j] == text[i+j] {
  j++
  }
  if j == m {
  return i
  }
  }
  return -1
  }
  Worst case Time Complexity of the algorithm is O(m*n), we got the pattern at the end of the text or we did not get the
  pattern at all.
  Best case Time Complexity of this algorithm is O(m), The average case Time Complexity of this algorithm is O(n)
  Robin-Karp algorithm
  Robin-Karp algorithm is somewhat similar to the brute force algorithm. Because the pattern is compared to each
  portion of text of length m. Instead of pattern at each position a hash code is compared, only one comparison is
  performed. The hash code of the pattern is compared with the hash code of the text window. We try to keep the hash
  code as unique as possible.
  The two features of good hash code are:
  · The collision should be excluded as much as possible. A collision occurs when hash code matches, but the pattern
  does not.
  · The hash code of text must be calculated in constant time.
  Hash value of text of length m is calculated. Each time we exclude one character and include next character. The
  portion of text that will be compared move as a window of characters. For each window calculation of hash is done in
  constant time, one member leaves the window and a new number enters a window.
  Multiplication by 2 is same as left shift operation. Multiplication by 2m-1 is same as left shift m-1 times. We want this
  multiple times so just store it in variable pow(m) = 2m-1
  We do not want to do big multiplication operations so modular operation with a prime number is used.
  Example 14.3:
  func RobinKarp(text string, pattern string) int {
  n := len(text)
  m := len(pattern)
  prime := 101
  powm := 1
  TextHash := 0
  PatternHash := 0
  i, j := 0, 0
  if m == 0 || m > n {
  return -1
  }
  for i = 0; i < m-1; i++ {
  powm = (powm << 1) % prime
  }
  for i = 0; i < m; i++ {
  PatternHash = ((PatternHash << 1) + (int)(pattern[i])) % prime
  TextHash = ((TextHash << 1) + (int)(text[i])) % prime
  }
  for i = 0; i <= (n - m); i++ {
  if TextHash == PatternHash {
  for j = 0; j < m; j++ {
  if text[i+j] != pattern[j] {
  break
  }
  }
  if j == m {
  return i
  }
  }
  TextHash = (((TextHash - (int)(text[i])*powm) << 1) + (int)(text[i+m])) % prime
  if TextHash < 0 {
  TextHash = (TextHash + prime)
  }
  }
  return -1
  }
  Knuth-Morris-Pratt algorithm
  There is an inefficiency in the brute force method of string matching. After a shift of the pattern, the brute force
  algorithm forgotten all the information about the previous matched symbols. This is because of which its worst case
  Time Complexity is O(mn).
  The Knuth-Morris-Pratt algorithm make use of this information that is computed in the previous comparison. It never
  re compares the whole text. It uses preprocessing of the pattern. The preprocessing takes O(m) time and whole
  algorithm is O(n)
  Preprocessing step: we try to find the border of the pattern at a different prefix of the pattern.
  A prefix is a string that comes at the start of a string.
  A proper prefix is a prefix that is not the complete string. Its length is less than the length of the string.
  A suffix is a string that comes at the end of a string.
  A proper suffix is a suffix that is not the complete string. Its length is less than the length of the string.
  A border is a string that is both proper prefix and a proper suffix.
  Example 14.4:
  func KMPPreprocess(pattern string, ShiftArr []int) {
  m := len(pattern)
  i := 0
  j := -1
  ShiftArr[i] = -1
  for i < m {
  for j >= 0 && pattern[i] != pattern[j] {
  j = ShiftArr[j]
  }
  i++
  j++
  ShiftArr[i] = j
  }
  }
  We have to loop outer loop for the text and inner loop for the pattern when we have matched the text and pattern
  mismatch, we shift the text such that the widest border is considered and then the rest of the pattern matching is
  resumed after this shift. If again a mismatch happens then the next mismatch is taken.
  Example 14.5:
  func KMP(text string, pattern string) int {
  i, j := 0, 0
  n := len(text)
  m := len(pattern)
  ShiftArr := make([]int, m+1)
  KMPPreprocess(pattern, ShiftArr)
  for i < n {
  for j >= 0 && text[i] != pattern[j] {
  j = ShiftArr[j]
  }
  i++
  j++
  if j == m {
  return (i - m)
  }
  }
  return -1
  }
  Example 14.6: Use the same KMP algorithm to find the number of occurrences of the pattern in a text.
  func KMPFindCount(text string, pattern string) int {
  i, j := 0, 0
  count := 0
  n := len(text)
  m := len(pattern)
  ShiftArr := make([]int, m+1)
  KMPPreprocess(pattern, ShiftArr)
  for i < n {
  for j >= 0 && text[i] != pattern[j] {
  j = ShiftArr[j]
  }
  i++
  j++
  if j == m {
  count++
  j = ShiftArr[j]
  }
  }
  return count
  }
  Dictionary / Symbol Table
  A symbol table is a mapping between a string (key) and a value that can be of any type. A value can be an integer such
  as occurrence count, dictionary meaning of a word and so on.
  Binary Search Tree (BST) for Strings
  Binary Search Tree (BST) is the simplest way to implement symbol table. Simple strcmp() function can be used to
  compare two strings. If all the keys are random, and the tree is balanced. Then on an average key lookup can be done in
  O(logn) time.
  Below is an implementation of binary search tree to store string as key. This will keep track of the occurrence count of
  words in a text.
  Example 14.7:
  type Node struct {
  value string
  count int
  lChild *Node
  rChild *Node
  }
  type StringTree struct {
  root *Node
  }
  func (t *StringTree) print() {
  t.printUtil(t.root)
  }
  func (t *StringTree) printUtil(curr *Node) {
  if curr != nil {
  fmt.Println(" value is ::", curr.value)
  fmt.Println(" count is :: ", curr.count)
  t.printUtil(curr.lChild)
  t.printUtil(curr.rChild)
  }
  }
  func (t *StringTree) Insert(value string) {
  t.root = t.insertUtil(value, t.root)
  }
  func (t *StringTree) insertUtil(value string, curr *Node) *Node {
  var compare int
  if curr == nil {
  curr = new(Node)
  curr.value = value
  } else {
  compare = strings.Compare(curr.value, value)
  if compare == 0 {
  curr.count++
  } else if compare == 1 {
  curr.lChild = t.insertUtil(value, curr.lChild)
  } else {
  curr.rChild = t.insertUtil(value, curr.rChild)
  }
  }
  return curr
  }
  func (t *StringTree) freeTree() {
  t.root = nil
  }
  func (t *StringTree) Find(value string) bool {
  ret := t.findUtil(t.root, value)
  fmt.Println("Find ", value, " Return ", ret)
  return ret
  }
  func (t *StringTree) findUtil(curr *Node, value string) bool {
  var compare int
  if curr == nil {
  return false
  }
  compare = strings.Compare(curr.value, value)
  if compare == 0 {
  return true
  }
  if compare == 1 {
  return t.findUtil(curr.lChild, value)
  }
  return t.findUtil(curr.rChild, value)
  }
  func (t *StringTree) frequency(value string) int {
  return t.frequencyUtil(t.root, value)
  }
  func (t *StringTree) frequencyUtil(curr *Node, value string) int {
  var compare int
  if curr == nil {
  return 0
  }
  compare = strings.Compare(curr.value, value)
  if compare == 0 {
  return curr.count
  }
  if compare > 0 {
  return t.frequencyUtil(curr.lChild, value)
  }
  return t.frequencyUtil(curr.rChild, value)
  }
  func main() {
  tt := new(StringTree)
  tt.Insert("banana")
  tt.Insert("apple")
  tt.Insert("mango")
  fmt.Println("Search results for apple, banana, grapes and mango :")
  tt.Find("apple")
  tt.Find("banana ")
  tt.Find("grapes")
  tt.Find("mango")
  }
  Output:
  Find apple Return true
  Find banana Return true
  Find grapes Return false
  Find mango Return true
  Hash-Table
  The Hash-Table is another data structure that can be used for symbol table implementation. Below Hash-Table
  diagram, we can see the name of that person is taken as the key, and their meaning is the value of the search. The first
  key is converted into a hash code by passing it to appropriate hash function. Inside hash function the size of Hash-Table
  is also passed, which is used to find the actual index where values will be stored. Finally, the value, which is meaning
  of name is stored in the Hash-Table, or you can store a reference to the string which store meaning can be stored into
  the Hash-Table.
  Hash-Table has an excellent lookup of O(1).
  Let us suppose we want to implement autocomplete the box feature of Google search. When you type some string to
  search in google search, it propose some complete string even before you have done typing. BST cannot solve this
  problem as related strings can be in both right and left subtree.
  The Hash-Table is also not suited for this job. One cannot perform a partial match or range query on a Hash-Table.
  Hash function transforms string to a number. Moreover, a good hash function will give a distributed hash code even for
  partial string and there is no way to relate two strings in a Hash-Table.
  Trie and Ternary Search tree are a special kind of tree that solves partial match and range query problem well.
  Trie
  Trie is a tree, in which we store only one character at each node. The final key value pair is stored in the leaves. Each
  node has R children, one for each possible character. For simplicity purpose, let us consider that the character set is 26,
  corresponds to different characters of English alphabets.
  Trie is an efficient data structure. Using Trie, we can search the key in O(M) time. Where M is the maximum string
  length. Trie is also suitable for solving partial match and range query problems.
  Example 14.8:
  package main
  import (
  "fmt"
  "strings"

  )
  type Node struct {
  isLastChar bool
  child [26](*Node)
  }
  type Trie struct {
  root *Node
  }
  func (t *Trie) Insert(s string) {
  if s == "" {
  return
  }
  str := strings.ToLower(s)
  t.root = t.InsertUtil(t.root, str, 0)
  }
  func (t *Trie) InsertUtil(curr *Node, str string, index int) *Node {
  if curr == nil {
  curr = new(Node)
  }
  if len(str) == index {
  curr.isLastChar = true
  } else {
  curr.child[str[index]-'a'] = t.InsertUtil(curr.child[str[index]-'a'], str, index+1)
  }
  return curr
  }
  func (t *Trie) Remove(s string) {
  if s == "" {
  return
  }
  str := strings.ToLower(s)
  t.RemoveUtil(t.root, str, 0)
  }
  func (t *Trie) RemoveUtil(curr *Node, str string, index int) {
  if curr == nil {
  return
  }
  if len(str) == index {
  if curr.isLastChar {
  curr.isLastChar = false
  }
  return
  }
  t.RemoveUtil(curr.child[str[index]-'a'], str, index+1)
  }
  func (t *Trie) Find(s string) bool {
  if s == "" {
  return false
  }
  str := strings.ToLower(s)
  return t.FindUtil(t.root, str, 0)
  }
  func (t *Trie) FindUtil(curr *Node, str string, index int) bool {
  if curr == nil {
  return false
  }
  if len(str) == index {
  return curr.isLastChar
  }
  return t.FindUtil(curr.child[str[index]-'a'], str, index+1)
  }
  func main() {
  t := new(Trie)
  a := "apple"
  b := "app"
  c := "appletree"
  d := "tree"
  t.Insert(a)
  t.Insert(d)
  fmt.Println(t.Find(a))
  fmt.Println(t.Find(b))
  fmt.Println(t.Find(c))
  fmt.Println(t.Find(d))
  }
  Output:
  true
  false
  false
  true
  Ternary Search Trie/ Ternary Search Tree
  Tries having a very good search performance of O(M) where M is the maximum size of the search string. However,
  tries having very high space requirement. Every node Trie contain references to multiple nodes, each reference
  corresponds to possible characters of the key. To avoid this high space requirement Ternary Search Trie (TST) is used.
  A TST avoid the heavy space requirement of the traditional Trie while still keeping many of its advantages. In a TST,
  each node contains a character, an end of key indicator, and three references. The three references are corresponding to
  current char hold by the node(equal), characters less than and character greater than.
  The Time Complexity of ternary search tree operation is proportional to the height of the ternary search tree. In the
  worst case, we need to traverse up to 3 times that many links. However, this case is rare. Therefore, TST is a very good
  solution for implementing Symbol Table, Partial match and range query.
  Example 14.9:
  type Node struct {
  data byte
  isLastChar bool
  left, equal, right *Node
  } type TST struct {
  root *Node
  }
  func (t *TST) Insert(word string) {
  t.root = t.insertUtil(t.root, word, 0)
  }
  func (t *TST) insertUtil(curr *Node, word string, wordIndex int) *Node {
  if curr == nil {
  curr = new(Node)
  curr.data = word[wordIndex]
  }
  if word[wordIndex] < curr.data {
  curr.left = t.insertUtil(curr.left, word, wordIndex)
  } else if word[wordIndex] > curr.data {
  curr.right = t.insertUtil(curr.right, word, wordIndex)
  } else {
  if wordIndex < len(word)-1 {
  curr.equal = t.insertUtil(curr.equal, word, wordIndex+1)
  } else {
  curr.isLastChar = true
  }
  }
  return curr
  }
  func (t *TST) findUtil(curr *Node, word string, wordIndex int) bool {
  if curr == nil {
  return false
  }
  if word[wordIndex] < curr.data {
  return t.findUtil(curr.left, word, wordIndex)
  } else if word[wordIndex] > curr.data {
  return t.findUtil(curr.right, word, wordIndex)
  } else {
  if wordIndex == len(word)-1 {
  return curr.isLastChar
  }
  return t.findUtil(curr.equal, word, wordIndex+1)
  }
  }
  func (t *TST) Find(word string) bool {
  ret := t.findUtil(t.root, word, 0)
  fmt.Print(word, " :: ")
  if ret {
  fmt.Println(" Found ")
  } else {
  fmt.Println("Not Found ")
  }
  return ret
  }
  func main() {
  tt := new(TST)
  tt.Insert("banana")
  tt.Insert("apple")
  tt.Insert("mango")
  fmt.Println("Search results for apple, banana, grapes and mango :")
  tt.Find("apple")
  tt.Find("banana")
  tt.Find("mango")
  tt.Find("grapes")
  }
  Output:
  Search results for apple, banana, grapes and mango :
  apple :: Found
  banana :: Found
  mango :: Found
  grapes :: Not Found
  Problems in String
  Regular Expression Matching
  Implement regular expression matching with the support of ‘?’ and ‘*’ special character.
  ‘?’ Matches any single character.
  ‘*’ Matches zero or more of the preceding element.
  Example 14.10:
  func matchExpUtil(exp string, str string, i int, j int) bool {
  if i == len(exp) && j == len(str) {
  return true
  }
  if (i == len(exp) && j != len(str)) || (i != len(exp) && j == len(str)) {
  return false
  }
  if exp[i] == '?' || exp[i] == str[j] {
  return matchExpUtil(exp, str, i+1, j+1)
  }
  if exp[i] == '*' {
  return matchExpUtil(exp, str, i+1, j) || matchExpUtil(exp, str, i, j+1) || matchExpUtil(exp, str, i+1, j+1)
  }
  return false
  }
  func matchExp(exp string, str string) bool {
  return matchExpUtil(exp, str, 0, 0)
  }
  Order Matching
  Given a long text string and a pattern string. Find if the characters of pattern string are in the same order in text string.
  Eg. Text String: ABCDEFGHIJKLMNOPQRSTUVWXYZ Pattern string: JOST
  Example 14.11:
  // Match if the pattern is present in the source text.
  func match(source string, pattern string) int {
  iSource := 0
  iPattern := 0
  sourceLen := len(source)
  patternLen := len(pattern)
  for iSource = 0; iSource < sourceLen; iSource++ {
  if source[iSource] == pattern[iPattern] {
  iPattern++
  }
  if iPattern == patternLen {
  return 1
  }
  }
  return 0
  }
  Unique Characters
  Write a function that will take a string as input and return 1 if it contain all unique characters else return 0.
  Example 14.12:
  func isUniqueChar(str string) bool {
  mp := make(map[byte]int)
  size := len(str)
  for i := 0; i < size; i++ {
  c := str[i]
  if mp[c] != 0 {
  fmt.Println("Duplicate detected!")
  return false
  }
  mp[c] = 1
  }
  fmt.Println("No duplicate detected!")
  return true
  }
  Permutation Check
  Example 14.13: Function to check if two strings are permutation of each other.
  func isPermutation(s1 string, s2 string) bool {
  count := make(map[byte]int)
  length := len(s1)
  if len(s2) != length {
  fmt.Println(s1, "&", s2, "are not permutation")
  return false
  }
  for i := 0; i < length; i++ {
  ch := s1[i]
  count[ch]++
  ch = s2[i]
  count[ch]--
  }
  for i := 0; i < length; i++ {
  ch := s1[i]
  if count[ch] != 0 {
  fmt.Println(s1, "&", s2, "are not permutation")
  return false
  }
  }
  fmt.Println(s1, "&", s2, "are permutation")
  return true
  }
  Palindrome Check
  Example 14.14: Find if the string is a palindrome or not
  func isPalindrome(str string) bool {
  i := 0
  j := len(str) - 1
  for i < j && str[i] == str[j] {
  i++
  j--
  }
  if i < j {
  fmt.Println("String is not a Palindrome")
  return false
  }
  fmt.Println("String is a Palindrome")
  return true
  }
  Time Complexity is O(n) and Space Complexity is O(1)
  Power function
  Example 14.15: Function which will calculate xn, Taking x and n as argument.
  func pow(x int, n int) int {
  var value int
  if n == 0 {
  return 1
  } else if n%2 == 0 {
  value = pow(x, n/2)
  return (value * value)
  } else {
  value = pow(x, n/2)
  return x * value * value
  }
  }
  String Compare function
  Write a function strcmp() to compare two strings. The function return values should be:
  a) The return value is 0 indicates that both first and second strings are equal.
  b) The return value is negative indicates the first string is less than the second string.
  c) The return value is positive indicates that the first string is greater than the second string.
  Example 14.16:
  func strcmp(a string, b string) int {
  index := 0
  len1 := len(a)
  len2 := len(b)
  minlen := len1
  if len1 > len2 {
  minlen = len2
  }
  for index < minlen && a[index] == b[index] {
  index++
  }
  if index == len1 && index == len2 {
  return 0
  } else if len1 == index {
  return -1
  } else if len2 == index {
  return 1
  }
  return (int)(a[index]) - (int)(b[index])
  }
  Reverse String
  Example 14.17: Reverse all the characters of a string.
  func reverseString(a string) string {
  chars := []rune(a)
  reverseStringUtil(chars)
  return string(chars)
  }
  func reverseStringUtil(a []rune) {
  lower := 0
  upper := len(a) - 1
  for lower < upper {
  a[lower], a[upper] = a[upper], a[lower]
  lower++
  upper--
  }
  }
  Reverse Words
  Example 14.18: Reverse order of words in a string sentence.
  func reverseStringRange(a []rune, lower int, upper int) {
  for lower < upper {
  a[lower], a[upper] = a[upper], a[lower]
  lower++
  upper--
  }
  }
  func reverseWords(str string) string {
  length := len(str)
  upper := -1
  lower := 0
  arr := []rune(str)
  for i := 0; i < length; i++ {
  if arr[i] == ' ' {
  reverseStringRange(arr, lower, upper)
  lower = i + 1
  upper = i
  } else {
  upper++
  }
  }
  reverseStringRange(arr, lower, upper)
  reverseStringRange(arr, 0, length-1)
  return string(arr)
  }
  Print Anagram
  Example 14.19: Given a string as character list, print all the anagram of the string.
  func printAnagram(a string) {
  n := len(a)
  printAnagramUtil([]rune(a), n, n)
  }
  func printAnagramUtil(a []rune, max int, n int) {
  if max == 1 {
  fmt.Println(string(a))
  }
  for i := -1; i < max-1; i++ {
  if i != -1 {
  a[i], a[max-1] = a[max-1], a[i]
  }
  printAnagramUtil(a, max-1, n)
  if i != -1 {
  a[i], a[max-1] = a[max-1], a[i]
  }
  }
  }
  Shuffle String
  Example 14.20: Write a program to convert list ABCDE12345 to A1B2C3D4E5
  func shuffle(arr string) string {
  ar := []rune(arr)
  n := len(ar) / 2
  count := 0
  k := 1
  var temp rune
  for i := 1; i < n; i = i + 2 {
  k = i
  temp = ar[i]
  for true {
  k = (2 * k) % (2*n - 1)
  temp, ar[k] = ar[k], temp
  count++
  if i == k {
  break
  }
  }
  if count == (2*n - 2) {
  break
  }
  }
  return string(ar)
  }
  Exercise
  1. Given a string, find the longest substring without reputed characters.
  2. The function memset() copies ch into the first 'n' characters of the string
  3. Serialize a collection of string into a single string and de serializes the string into that collection of strings.
  4. Write a smart input function, which take 20 characters as input from the user. Without cutting some word.
  User input: “Harry Potter must not go”
  First 20 chars: “Harry Potter must no”
  Smart input: “Harry Potter must”
  5. Write a code that returns if a string is palindrome and it should return true for below inputs too.
  Stella won no wallets.
  No, it is open on one position.
  Rise to vote, Sir.
  Won't lovers revolt now?
  6. Write an ASCII to integer function, which ignore the non-integral character and give the integer. For example, if the
  input is “12AS5” it should return 125.
  7. Write code that would parse a Bash brace expansion.
  Example: the expression "(a, b, c) d, e" and would output all the possible strings: ad, bd, cd, e
  8. Given a string write a function to return the length of the longest substring with only unique characters
  9. Replace all occurrences of "a" with "the"
  10. Replace all occurrences of %20 with ' '.
  E.g. Input: www.Hello%20World.com
  Output: www.Hello World. com
  11. Write an expansion function that will take an input string like "1..5,8,11..14,18,20,26..30" and will print
  "1,2,3,4,5,8,11,12,13,14,18,20,26,27,28,29,30"
  12. Suppose you have a string like "Thisisasentence". Write a function that would separate these words. Moreover, will
  print whole sentence with spaces.
  13. Given three string str1, str2 and str3. Write a complement function to find the smallest sub-sequence in str1 which
  contains all the characters in str2 and but not those in str3.
  14. Given two strings A and B, find whether any anagram of string A is a sub string of string B.
  For eg: If A = xyz and B = afdgzyxksldfm then the program should return true.
  15. Given a string, find whether it contains any permutation of another string. For example, given "abcdefgh" and "ba",
  the function should return true, because "abcdefgh" has substring "ab", which is a permutation of the given string
  "ba".
  16. Give an algorithm which removes the occurrence of “a” by “bc” from a string? The algorithm must be in-place.
  17. Given a string "1010101010" in base2 convert it into string with base4. Do not use an*/
