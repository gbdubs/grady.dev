+++

name = "DumbDB"
external_url = ""

short_desc = "Dumb databases make for good prototypes"
tags = ["webapp", "functional", "javascript"]

+++

While protyping front end interfaces for responsive web applications, I frequently have to mock out the data interactions with static constants in my javascript that mock server requests and responses. This has several downsides: you cannot change the data in a demo situation, inserting real requests requires you to construct the remoteand the number of mocks can explode if you are trying to make a CRUD user interface.

DumbDB is a better way.  DumbDB is a remote server/database that accepts arbitrary post and get messages, and maintains an object store that maps the requested paths to the objects that have previously been stored at each.  This mapping between path and object is the simplest form of an object-store (dumb), but is a functionally complete database (db).

DumbDB has already helped me accelerate my prototyping on several projects.

It is implemented as a MongoDB database on GCP with a thin layer of GCP cloud functions on top. One of the features of DumbDB is that it dosen’t have authentication, but does have rate limiting by IP. This prevents eggregiously abusive behavior, but requires no configuration on the client side.
