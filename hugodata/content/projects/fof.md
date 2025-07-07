---
aliases: ["/project/fof/"]
title: "Focus on Features"
subtitle: "A resource for regulators and product managers to minimize digital harms"
summary: "A resource for regulators and product managers to minimize digital harms"
type: "projects"
isSimple: false
isPriority: true
sortPriority: 6000
draft: false
date: 2023-01-01
tags: ["government", "online spaces"]
CTALink: "https://fellowship.ii.grady.dev"
CTA: "Browse Focus on Features"
CTAPreamble: "Explore the resource:"
featured_image: "/img/fof_2.png"
---

When I worked in Counter-Abuse at Google, I found the largest issue we faced
wasn't malicious users, but poorly thought out product decisions, and the 
constraints that the business model placed on what we could do to keep users safe. 
It's very hard to clean up a kitchen if you're only empowered
to mop, rather than trying to understand or change why things keep spilling on the floor.

So when the Integrity Institutute (a small thinktank on content moderation issues, 
which I've been a member of for several years) sent out applications for visiting
fellowships, I figured it was my chance to try to do something about it. I built a 
resource for product managers and regulators to think through the connections between
harms they see on platforms, and the features that enable that harm, and the set of
changes that could minimize that connectivity.
The best explanation I have is the ["Why I built this" recursive essay here](https://fellowship.ii.grady.dev/about/why).

![The "Self-Harm amplification" page on the FOF website.](/img/fof_1.png)

The site as a whole took a while to build because there was just a lot of writing to do.
As of this writing, there are 14 intervention categories, 47 harms, 64 interventions, and 9 features.

The project ultimately became a group project of the Integrity Institute. Because of that
every page needed to be fully-editable by a nontechnical audience, which presented a fun set of technical
challenges.

![The editing menu the FOF website.](/img/fof_3.png)

The technical elements of this project are pretty slick. It's backed by a postgres
database running on neon.tech, which connects to Strapi, an open source CMS.
The frontend is made in Nuxt 3 with typescript, and then hosted via firebase cloud functions.

However, that's the editing-path. Since the content is static, I set up SSG to
pre-create every page on the site, and then back-hydrate with a flat file after 
first load. So the page as a whole is incredibly snappy, and costs nothing to host
since it's just flat files for almost everyone using it, and scale-to-zero for editors.

I've gotten more traction on this project than I expected - it's started many conversations
with PMs, regulators, and other thinktanks, and still (years later) has a steady 
stream of traffic, mostly from the EU.

![A filter view tailored for folks trying to find the most helpful/least costly interventions](/img/fof_2.png)

Browse for yourself [here](https://features.integrityinstitute.org).
