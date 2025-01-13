+++

order = 400

name = "GradyDocs"
external_url = "http://www.gradydocs.com"

short_desc = "GradyDocs - a shittier version of Google Docs"
tags = ["webapp", "nodejs", "aws", "mongodb"]

+++

I one time made the mistake of looking at how the Google Docs DOM is structured. I wondered if I could make a functional front end that looked similar without half a meg of divs, so I tried!

Then I wrote a backend for it (node-js, socketio), and got the project to the point where two users could edit the document together with sub-second lag. 

The underlying data model for this was not only flawed, it is abysmal.  Each movement, keypress and modification that a user makes was logged as a change event. Every change event was sent through a server on AWS-beanstalk in Node, via socket-io to every other web client subscribed to the same document.

It was honestly expensive to run this server, so I shut it down a while ago. When 10+ people were on a document it became expensive. 

Additionally, I had no server side compilation logic, so all of the docs were just stored as a series of event modifications.  Horrifying. 

I hope to revisit this project in the future, because it was really fun to write.  I wrote it in two weeks over my christmas break, at a time when I didn’t know much about much.  I would love to take another stab at it and set up a mesh of peer-to-peer connections with a sane document reconsiliaiton model in the future.

