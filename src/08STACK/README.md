Introduction
A stack is a basic data structure that organized items in last-in-first-out (LIFO) manner. Last element inserted in a stack
will be the first to be removed from it.
The real-life analogy of the stack is "stack of plates". Imagine a stack of plates in a dining area everybody takes a plate
at the top of the stack, thereby uncovering the next plate for the next person.
Stack allow to only access the top element. The elements that are at the bottom of the stack are the one that is going to
stay in the stack for the longest time.
Computer science also has the common example of a stack. Function call stack is a good example of a stack. Function
main() calls function foo() and then foo() calls bar(). These function calls are implemented using stack first bar() exists,
then foo() and then finally main().
As we navigate from web page to web page, the URL of web pages are kept in a stack, with the current page URL at
the top. If we click back button, then each URL entry is popped one by one.