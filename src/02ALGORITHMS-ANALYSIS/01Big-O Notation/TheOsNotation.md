Big-O Notation
Definition: “f(n) is big-O of g(n)” or f(n) = O(g(n)), if there are two +ve constants c and n0 such that
f(n) ≤ c g(n) for all n ≥ n0,
In other words, c g(n) is an upper bound for f(n) for all n ≥ n0
The function f(n) growth is slower than c g(n)
We can simply say that after a sufficient large value of input N the (c.g(n)) will always be greater than f(n).
Example: n2 + n = O(n2)
Omega-Ω Notation
Definition: “f(n) is omega of g(n).” or f(n) = Ω(g(n)) if there are two +ve constants c and n0 such that
c g(n) ≤ f(n) for all n ≥ n0
In other words, c g(n) is lower bound for f(n)
Function f(n) growth is faster than c g(n)
Find relationship of f(n) = nc and g(n) = cn
f(n) = Ω(g(n))
Theta-Θ Notation
Definition: “f(n) is theta of g(n).” or f(n) = Θ(g(n)) if there are three +ve constants c1, c2 and n0 such that c1 g(n) ≤
f(n) ≤ c2 g(n) for all n ≥ n0
Function g(n) is an asymptotically tight bound on f(n). Function f(n) grows at the same rate as g(n).
Example: n3 + n2 + n = Ɵ(n3)
Example: n2 + n = Ɵ(n2)
Find relationship of f(n) = 2n2 + n and g(n) = n2
f(n) = O(g(n))
f(n) = Ɵ(g(n))
f(n) = Ω(g(n))
Note:- Asymptotic Analysis is not perfect, but that is the best way available for analyzing algorithms.
For example, say there are two sorting algorithms first take f(n) = 10000*n*log(n) and second f(n) = n2 time. The
asymptotic analysis says that the first algorithm is better (as it ignores constants) but actually for a small set of data
when n is small then 10000, the second algorithm will perform better. To consider this drawback of asymptotic analysis
case analysis of the algorithm is introduced.
Complexity analysis of algorithms
1. Worst Case complexity: It is the complexity of solving the problem for the worst input of size n. It provides the
upper bound for the algorithm. This is the most common analysis done.
2. Average Case complexity: It is the complexity of solving the problem on an average. We calculate the time for all
the possible inputs and then take an average of it.
3. Best Case complexity: It is the complexity of solving the problem for the best