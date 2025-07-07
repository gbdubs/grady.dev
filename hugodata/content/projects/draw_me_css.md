---
aliases: ["/project/draw_me_css/"]
title: "Draw Me CSS"
subtitle: "Simple framework for making websites that magically appear"
summary: "Simple framework for making websites that magically appear"
type: "projects"
isSimple: false
isPriority: false
sortPriority: 500
draft: false
date: 2017-01-01
tags: ["art", "websites", "animation", "tools"]
CTALink: "https://github.com/gbdubs/draw-me"
CTA: "View on GitHub"
CTAPreamble: "Check out the framework code:"
featured_image: ""
---

In the third Harry Potter book, Harry is given the “Maurauder’s Map”, and one of the many images I like about this map is how the ink spreads outward from the point that it is awoken. 

I wondered if you could do this in CSS, so I poked around, learned about CSS Transitions, and came up with a generalized framework for making a set of boxes animate on to the page. The final product takes in a source file of HTML, and generates minimized CSS that includes all of the style classes needed to animate the page according to the class specification.

DrawMeCSS was used to create several other sites, including Parley (which I worked on co-concurrently) and V2 of my personal website (which was the main test case). 

I have ultimately abaondoned DMCSS because it generates too much code, requires too much thought and logic to be placed in the DOM, and leads to lots of empty and meaningless divs. Version 3 of my portfolio abandoned it for all of these reasons.
