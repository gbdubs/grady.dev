---
aliases: ["/project/senior_thesis/"]
title: "Senior Thesis"
subtitle: "Senior thesis on the graph isomorphism (GI) problem"
date: 2024-01-01T00:00:00Z
draft: false

# Project classification
type: "projects"
tags: ["algorithms", "education", "tools"]
isPriority: false
sortPriority: 1300
isSimple: false

# Call to action
CTA: "Read Thesis"
CTALink: "https://github.com/gbdubs/thesis"
CTAPreamble: "Want to read the full academic work"

# Optional image
featured_image: ""

# Summary (used in lists)
summary: "Senior thesis on the graph isomorphism (GI) problem"
---

I wrote my undergraduate thesis on the [Graph Isomorphism](http://mathworld.wolfram.com/IsomorphicGraphs.html) (GI) problem, and several topics within that space. My work was [version controlled on github](https://github.com/gbdubs/thesis), and the final result is [here](https://github.com/gbdubs/thesis/blob/master/Documentation%20%2B%20Reports/THESIS/thesis.pdf). 

I was first interested in the “Paths” invariant, which is incredibly powerful test for discerning between potentially isomorphic graphs in cubic (and fully paralellizable) time, and was originally theorized to be a solution to GI. The primary findings of my thesis were that:

* [Computational Wrangling] Unlike was theorized, copaths graphs do exist, and finding them is easily done brute force calculations on today's standard GPUs.
* [Theoretical Proof] Paths uniquely determines the chromatic polynomial of a graph.
* [Analytical Algorithms] Measuring discriminatory power to value ratio of invariants within graph isomorphism is possible and can help us devise better algorithms for practical isomorphism detection, and the paths invariant is well weighted in an analysis of discriminatory power.
* [Theoretical Proof] Paths as an invariant does not need to be calculated past N vertices.

Since I answered most questions I was interested in  paths stuff earlier than expected, I started working on Random Graph theory and found that:

* [Literature Review] Random Graph generators almost all are random matrix generators. This understates the frequency of highly auto-isomorphic graphs by one-to-factorial odds.
* [Coding] We can create a better random graph generator to evaluate random graph algorithms.
* [Analytical Proof] Many algorithms are evaluated across "matrix random" graphs, and their amortized running time can be exponentially higher when using a true random graph generator.

I defended this thesis was successfully defended in May of 2016, and recieved highest honors in Computer Science.


