+++

order = 1200

name = "Banannagrams"
external_url = ""

short_desc = "A sub-linear time algorithm for playing Bananagrams"
tags = ["algorithms", "theory", "java"]

+++

In 2016, I had just started at Google, and had started having breakfast every day with the same group of folks - Chris, Brett, Andrew and Brandon. All of us being nerds, techno-optimists and deeply competitive, at one point Brandon described having a competition amongst us to see who could write the best AI to play the game [bananagrams](www.gradybward.com). 

My productivity dipped precipitously over the next few days as I considered how to crush my good friends.

Here is the algorithm I designed.  I realized that searching for connections was going to take too long, and that I needed to support a simpler mechanism for tile addition.  I figured out a spatial layout that assigned ownership over strips of the grid in perpetuity to "strands". Each Strand was a series of words where the ends form a two letter word and have no other interactions. In this way, each word simply becomes a starting character and finishing character, as length no longer matters, and this reduces the search costs considerably.

TODO(gradyward) put a small scale diagram here.

With this model in mind, you can immediately discard all words that dont start and end with letters that are constituent to two letter words (on the same upbeat/downbeat, too). You can further simplify out all string comparisons and just look at character integers. You only need maintain a list of all of the open docks, and a simple function to add a new doc on the end.  Throw in some logic about which strands go up and which go down, and you have a blazingly fast addition algorithm.

The code is [here](http://github.com/gbdubs/bananagrams), and the result is (in my mind) breathtaking.

![A Photo of an incredibly dense bananagrams grid.](../img/bananagrams.png)

Because all lookups are heavily ammortized, and the set of words is large enough that you only need to look where there are available keys, as the number of words on the board increases, the addition time actually decreases, though you run into memory issues at some points. 

The format Brandon used to create the server (Captn' Proto) doesn't have a Java API, so we didn't end up playing.

I learned a good amount from this project. The first is that decomposing a problem into steps that computers are good at is the key to solving hard algorithmic challenges. It doesn't matter how fancy or complete your logic is if its orthogonal to the strengths of your processor and the data structures you are using. The second is that high-level languages are going to be easier to optimize down to their core precisely because they add layers of semantic meaning. Optimizing C is not something I do well, because my mind is too stuck in the operations of the computer, where as when I am optimizing Java, I see the problem in front of me clearly. Thirdly, competition drives me to do some of my best work. There is nothing more motivating than sitting around a breakfast table, plotting victory against friends. I rise to (and love) challenges.