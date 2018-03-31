+++

order = 1100

name = "Parley"
external_url = "http://www.parley.us"

short_desc = "Recursive annotations as a counter-proposal to the comment thread"
tags = ["webapp", "design", "appengine", "ux"]

+++

Comment threads are an inherently flawed way to have a discussion, because they are strictly linear.  In recent years, large organizations like facebook have realized this and have tried to reconsile it with a binary tree pattern of _deeper_ or _next_. This only partially solves the problem. Far more frequently, if I want to respond to a comment, there are multiple peices of it that I want to respond to, and some may already be addressed further down the thread.  The heart of the matter is that a meaningful static discussion is not best facilitated by a linear progression of strings.

Parley was my idea on how to change this. It views a document as something which is not responed to in whole, but which can be annotated with comments by different authors.  The comments (in turn) become top level objects which are just as open to annotation and further comment. The recursive, n-ary tree nature of this structure has multiple advantages, but the big one is that you can dive deep on a given line of argument naturally, by clicking through and understanding the points as made in isolation.

This idea in and of itself is not revolutionary, but I wanted to see if I could make a demo where this is implemented in a way that is comprehensible and intuitive.  I half succeeded here.  The parley interface used the Draw Me CSS pattern to create a right-left layout where the content is displayed with underlines corresponding to annotations, and the text of the annotation is displayed in a scrolling left hand column.  The interface is flexible, relatively intuitive, and fairly simple.

Where I ran into problems was in the minutia -- what about really long annotations? What about annotations stacked 15 high? How do you link effectively (as links cannot then be annotated)? 

The backend is lame (an app engine hack), but I stand behind the fundamental idea (recursive comment sections, rather than linear ones), and I stand behind the UI for what it is - a proof of concept. 