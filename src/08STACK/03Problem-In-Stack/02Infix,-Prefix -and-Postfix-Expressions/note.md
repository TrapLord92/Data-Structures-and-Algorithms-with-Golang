Infix, Prefix and Postfix Expressions
When we have an algebraic expression like A + B then we know that the variable is being added to variable B. This
type of expression is called infix expression because the operator “+” is between operands A and operand B.
Now consider another infix expression A + B * C. In the expression there is a problem that in which order + and *
works. Does A and B are added first and then the result is multiplied. Alternatively, B and C are multiplied first and
then the result is added to A. This makes the expression ambiguous. To deal with this ambiguity we define the
precedence rule or use parentheses to remove ambiguity.
So if we want to multiply B and C first and then add the result to A. Then the same expression can be written
unambiguously using parentheses as A + (B * C). On the other hand, if we want to add A and B first and then the sum
will be multiplied by C we will write it as (A + B) * C. Therefore, in the infix expression to make the expression
unambiguous, we need parenthesis.
Infix expression: In this notation, we place operator in the middle of the operands.
< Operand > < operator > < operand >
Prefix expressions: In this notation, we place operator at the beginning of the operands.
< Operator > < operand > < operand >
Postfix expression: In this notation, we place operator at the end of the operands.
< Operand > < operand > < operator >
Infix Expression Prefix Expression Postfix Expression
A + B + A B A B +
A + (B * C) + A * B C A B C * +
(A + B) * C * + ABC A B + C *
Now comes the most obvious question why we need so unnatural Prefix or Postfix expressions when we already have
infix expressions which words just fine for us. The answer to this is that infix expressions are ambiguous and they need
parenthesis to make them unambiguous. While postfix and prefix notations do not need any parenthesis.