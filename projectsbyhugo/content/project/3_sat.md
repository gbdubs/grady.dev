+++

name = "3-SAT decomposition"
external_url = "https://github.com/gbdubs/2001"

short_desc = "Using 64 bit architecture to perform fast 3-SAT calculations"
tags = ["theory", "java"]

+++

Haven’t you ever looked at SAT (the problem of finding a solution to an arbitrary expression of boolean predicates), and thought it must be solvable in sub-polynomial time? My intuition says so too, but this is clearly a function of naievite. 

The thing that struck me about 3-SAT is that it is fundamentally trying to explore a bunch of different roads at once, and reporting back the ones that don’t hit dead ends.  You can do this in exponential time on a single processor, or you can do it in (near) constant time on an exponential number of processors.  However I was wondering if you could do it in linear time with a single processor by using an exponential amount of memory.  The answer, it turns out, is yes!

By using the OR bitwise operator, you can do SAT in linear time and exponential space (which for problems with fewer than 8 boolean variables, is constant on a 64 bit processor). 

This was really jsut an academic endeavour, but it was fun to design and run, and was my first exploration of the world of computational complexity.

