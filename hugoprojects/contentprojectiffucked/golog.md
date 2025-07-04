+++

order = 1000

name = "GoLog"

short_desc = "a visual post-hoc logger for golang"
tags = ["art", "css"]

+++

I've been writing a lot more golang of late, so I wrote a (gorgeous) visual debugger for helping to do post-hoc debugging. It's been hugely helpful in my lambda based projects, where one often wants the ability to quickly view "what happened during this trace" without running locally with a traditional step based debugger.

![A screenshot of the Golog interface for visual debugging](https://raw.githubusercontent.com/gbdubs/golog/master/golog-stack-trace.png)

Unlike many of my projects, this one is pretty well supported, and I've continued to add features as I need them.

You can check it out [on github here](https://github.com/gbdubs/golog).

