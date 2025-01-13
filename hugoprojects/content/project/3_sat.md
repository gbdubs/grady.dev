+++

order = 150

name = "3-SAT decomposition"
external_url = "https://github.com/gbdubs/2001"

short_desc = "Using exponential space to perform 3-SAT calculations in polynomial time"
tags = ["theory", "java", "algorithms"]

+++

Haven’t you ever looked at [SAT](https://en.wikipedia.org/wiki/Satisfiability) (the problem of finding a solution to an arbitrary expression of boolean predicates), and, realizing it can be reduced to 3-SAT, thought it must be solvable in sub-polynomial time? My intuition says so too, but this is clearly a function of naievite. 

The thing that struck me about SAT generally is that it is fundamentally trying to explore a bunch of different roads at once, and reporting back the ones that don’t hit dead ends.  You can do this in exponential time on a single processor, or you can do it in (near) constant time on an exponential number of processors. 

I was interested in exploring how this could be used semi-practically. I focused on **variable-small** satisfiability problems (where you might have a large set of operands, but a small set of variables), if you express each 3-SAT statement as an integer, where every bit of the integer represents a combination of states (i.e. 0 -> {A=F, B=F, C=F}, 1 -> {A=T, B=F, C=F}, etc.),
the binary OR operator can combine satisfiability statements in constant time if the number of variables is small enough (2 ^ 2 ^ N grows pretty quickly, but 4 is trivial, 5 is still useful), while retaining information on the full set of value sets that might be able to satisfy the predicate.

This is not a means you can use to solve the SAT problem, BUT it actually can be really useful in the computation if used alongside other optimizations and compressions.

I liken SAT to untying a giant knot. straightforward pulling through via simple compression algorithms get you most of the way, but having a pick to work through small, localized component problems dramatically speeds up overall solution time.

This was really jsut an academic endeavour, but it was fun to design and run, and was my first exploration of the world of computational complexity. The end algorithm scaled up really nicely with the number of predicates, but faced growing pains on the number of variables.

From this project I learned about how composite algorithms are frequently more effective than pure ones.  A problem may sound uniform, but different problem sub-configurations are likely to have different performance tradeoffs, and optimizing the whole means paying attention to the drawbacks and strengths of each constituent.