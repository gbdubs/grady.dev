---
aliases: ["/project/grady_in/"]
title: "Grady.In"
subtitle: "Go-Links for everybody"
date: 2024-01-01T00:00:00Z
draft: false

# Project classification
type: "projects"
tags: ["websites", "tools"]
isPriority: false 
sortPriority: 700
isSimple: false

# Call to action
CTA: "Try Grady.In"
CTALink: "http://www.grady.in"
CTAPreamble: "Want to create your own short links"

# Optional image
featured_image: "/img/grady_in.png"

# Summary (used in lists)
summary: "Go-Links for everybody"
---

Go Links are a system used interally at google to create semantically meaningful short links. Grady.In replicates this functionality for use outside of the company, allowing anyone to construct short and human readable links. 

This was my first forray into a serverless architechure. The system is built using Google Cloud Functions. The primary function validates a request and writes the new link (as a simple redirect) to a CDN. The client (when looking at the grady.in/xxx address) is pointed toward the CDN, and the static content served from the CDN redirects them to the pre-specified link, if it exists.

Notably absent from the implementation is any kind of rate limiting or authentication. This was really just meant as a fun way to play around with GCP’s cloud functions, and the lack of sophistication shows through in these obvious flaws.
