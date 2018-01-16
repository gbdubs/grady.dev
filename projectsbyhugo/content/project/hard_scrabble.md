+++

name = "Hardscrabble"
external_url = ""

short_desc = "Interactive pair programming webapp for a class of 200+ students"
tags = ["webapp", "app engine", "java"]

+++

Pair programming is a stellar way for people to learn, particularly when paired with a partner at the same level.  I saw this multiple times in action in my time as a teaching assistant for courses at Brandeis.  However, the problem of assigning pairs is fundamentally challenging - if the students are given free reign, they will likely pick their friends, and the skill level of the two will be (largely) uncorrelated.  If the teacher assigns specific pairs, this may be more effective, but is unscalable at high student volumes.

A better approach is to do it programmatically. Working for professor Antonella Di Lillo, I single handedly designed and implemented a system for pairing students up in class, and enabling them to communicate with one another via a shared document with set phases to enable live-coding excercises while in lecture.  The students would be presented with a shared problem that the professor would introduce in lecture. They would then be assigned a partner based on past performance of programs written using the software. Each student would independently try to solve the problem in a text editor that only they could see.  Then came a comment period, where each student could write comments on the code of the other.  Finally, the two students would be able to chat in a modal while they collaboratively work on coming up with a joint solution.

The student would be evaluated (for purposes of future pairing) based on their initila submission.  Basic metrics like attendance, pair history, lines of code and distance from ideal solution were reported to the professor in aggregated form. 

This was technically demanding - both figuring out a real time pairing algorithm for maximal skill matching, and figuring out the intricacies of how to write a real time code editor that the two students can pair program in. At the time, I was pretty unfamiliar with all of this stuff, so I did it poorly - using app engine and a basic state sharing algorithm.

This project was hacked together hastily as a favor for a favorite professor. Today, I would design the system quite differently - setting up peer-to-peer data channels in the browser between the lead computer and the followers, and each of the programming pairs. 

This project got signifigant use. In the year it was built, it served 200 students in the fall and 150 in the spring. In both semesters, both students and TAs were happy with the way it helped students learn.  This is a project I am proud of because of the ideas behind it, but also one where I recognize the deep and profound technical flaws. I hope someday to have the time and energy to write this code again to be (A) more technically sophistocated and robust (B) scalable (C) cost effective and (D) siloed for use in more than one classroom. 