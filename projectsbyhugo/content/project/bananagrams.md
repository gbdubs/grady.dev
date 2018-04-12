+++

order = 1200

name = "Banannagrams"
external_url = ""

short_desc = "A beautiful algorithm for playing Bananagrams"
tags = ["algorithms", "theory", "java"]

+++

In 2016, when I began working at Google, I started having breakfast with the same group of folks every morning - Chris, Brett, Andrew and Brandon. All of us being nerds, techno-optimists and deeply competitive, at one point Brandon described having a competition to see who could write the best AI to play the game [bananagrams](https://en.m.wikipedia.org/wiki/Bananagrams). 

My productivity dipped precipitously over the next few days as I considered how to destroy my good friends in tiled combat.

Here is the algorithm I designed.  I realized that searching for connections in any spatial model was going to dominate running time, and that I tile addition was going to be near-impossible if you had to maintain state over colliding words. To optimize so I didn't have to maintain a model of the board, I figured out a layout that assigned ownership over strips of the grid in perpetuity to "strands". 

![An Diagram of an example Strand](../img/bananagrams_strand.png)

![An Diagram of how strands merge together to fill the full playing space.](../img/bananagrams_dock.png)

Each Strand was a series of words where the ends form a two letter word and have no other interactions. In this way, each word simply becomes a starting character and finishing character, as length no longer matters, and this reduces the search costs considerably.

With this model in mind, you can immediately discard all words that dont start and end with letters that are constituent to two letter words (on the same upbeat/downbeat, too). You can further simplify out all string comparisons and just look at character integers. You only need maintain a list of all of the open docks, and a simple function to add a new doc on the end.  Throw in some logic about which strands go up and which go down, and you have a blazingly fast addition algorithm.

The code is [here](http://github.com/gbdubs/bananagrams), and the result is (in my mind) breathtaking.

![A Photo of an incredibly dense bananagrams grid.](../img/bananagrams.png)

Because all lookups are heavily ammortized, and the set of words is large enough that you only need to look where there are available keys, as the number of words on the board increases, the addition time actually decreases, though you run into memory issues at some points. 

The format Brandon used to create the server (Captn' Proto) doesn't have a Java API, so we didn't end up playing.

I learned a good amount from this project:

* Decomposing a problem into steps that computers are good at is the key to solving hard algorithmic challenges. It doesn't matter how fancy or complete your logic is if its orthogonal to the strengths of your processor and the data structures you are using. 
* High-level languages are going to be easier to optimize down to their core precisely because they add layers of semantic meaning. Optimizing C is not something I do well, because my mind is too stuck in the operations of the computer, where as when I am optimizing Java, I see the problem in front of me clearly. 
* Competition drives me to do some of my best work. There is nothing more motivating than sitting around a breakfast table, plotting victory against friends. I rise to (and love) challenges.