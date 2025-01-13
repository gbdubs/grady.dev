+++

order = 3999 

name = "Custom Board Games"

short_desc = "Themed gifts for friends and family"
tags = []

+++

I love making board games for friends and family as gifts, often
themed around a shared interest or an inside joke. While these were
originally seperate posts, I've condensed them here for brevity.

## 2024 - FRONTIER

I (and my boyfriend) both have a love-hate relationship with
Frontier Airlines. The fares are impossibly cheap, and if you have their
credit card and Discount Den memberhsip, things can be legitimately a 
fantastic bargin. However, they're often late, the staff isn't the
friendliest, and I've had some of my strangest experiences flying with
them. We also both really like the game PanAm, about the expansion
of the PanAm corporation - a gorgeous game with great mechanics like stock
pricing and technological upgrades. 

So, when my boyfriend mentioned offhandedly that I should make a version of
PanAm based off of FRONTIER, I was struck by the jealousy you have
when someone comes up with a great idea that you should have had. While I 
initially wrote off the idea for that Christmas (because this comment was 
made in early November), I took the cancellation of a flight on FRONTIER a 
week later as a sign/opportunity to build out a full version of the game.

![A very rough sketch of the board that I'd end up creating, with only the cities laid out and circles between them representing approximate distances](/img/frontier-rough.png)

The first sketch (above), the final board, 24h later, below:

![The map from the board game, an enormously complex SVG with colorful routes of airlines criscrossing the United States](/img/frontier-map.png)

This project was very low-fi - there were only two technical elements,
simulations to ensure the stock price deviation and expansion frequency was
distributionally good and within normal bounds - a property the real game 
fails to live up to . 


![The output of a simulation from the board game](/img/frontier-simulation.png)

The majority of the work on this project was just learning to use InkScape 
effectively - practicing my (minimal) SVG skills on a broader canvas and
learning the keyboard shortcuts, automation tools (like cloning), and becoming
diligent about layer and group naming, selectors, and tags.

![An event card from the frontier board game, giving the history of the spirit aquisition attempt from 2022](/img/frontier-event.png)

Only the destination cards' images were generated using AI - and even then 
there was a lot of subjective post processing. 

![A screenshot of a bunch of cards - each corresponding to an airport in the united states, with the airport code shown on top of an iconic landmark from the location of the airport](/img/frontier-dest.png)

The gift has yet to be given - but thankfully I don't think anyone reads my blog!

## 2024 - Puzzle Generation

You can check out my laser-cut puzzles [here](/project/laser_puzzles).

## 2023 - Tail Tally (Point Salad)

In 2023 I made a board game copy of Point Salad for my boyfriend.
He loves the game Point Salad and loves his dogs even more.

![The front of a card from the Tail Tally deck](/img/tail-tally-1.png)

![The back of a card from the Tail Tally deck](/img/tail-tally-2.png)

The code for it is written in Go,
generates HTML files (mostly to take advantage of flexbox positioning),
and then automatically screenshots those at 500dpi, with bleeding/cutoff
that matches the specifications from my favorite [board game printer](https://www.makeplayingcards.com/).

This one
was particularly fun because I made it fully generalizable - you just
input six photos and select a few things about the color scheme and sizing you
want, and the deck pops out. I've printed (but not yet gifted) another version
for friends - so beware, friends.

## 2022 - Wingspan

A writeup of the Wingspan project is [here](/project/wingspan/), but a few bonus images are below.

![A card depicting a blue jay from the auto-generated wingspan deck project](/img/wingspan-jay.png)

![A card depicting a bluebird from the auto-generated wingspan deck project](/img/wingspan-bluebird.png)

![A card depicting an osprey from the auto-generated wingspan deck project](/img/wingspan-osprey.png)

## 2020-2021 - Ticket to Ride

### 2020 - The Original Gift 

My mom loves the board game Ticket to Ride, perhaps to a fault.
Starting in 2018 (and apparently continuing to this day) she regularly plays it by
herself late into the night - with two distinct players (her right and
left hand) competing to see which one would win. 

> "I think everyone needs something to unwind at the end of a day, some people need to watch silly TV some people like to have a drink, for me it's board games against myself." - My mother

In 2020, we'd been planning on celebrating her 60th birthday as a family,
but we were unable to do so because of the pandemic. My sister and I, keenly
aware that this was a really big loss, decided to pour our energies into 
a great gift, and so took a weekend to create this board game. 

Each location is one that she has been to and has fond memories of. The photo
cards we created were a few select ones from family lore. The trains were 
dowels that I'd cut to length and dyed, and the tokens my sister made from
clay, each a fun inside joke.

![An early Ticket to Ride Card, featuring my parents on a green background, representing a green train card](/img/ttr-card1.png)

This was the first custom game I'd made, so looking back on it, the build
was fairly primitive - we sent back and forth PPT files in ultra HD.
The background is from the CINDI project. The files
are large enough that they've been lost to time, but the board(s) itself 
lives on, as do some artifacts from the early days of its creation.

![the original sketch of the map for the board](/img/ttr-map-1.png)

![the revised map for the board, including distances](/img/ttr-map-2.png)

![the final layout for the board](/img/ttr-map-3.png)

![the final board, well used and loved](/img/ttr-map-final.jpg)

My mom loved it, and it has further extended her TTR obsession.

### 2021 - Ticket To Ride Cards

When a year had gone by, the pandemic had continued, and my mother was still
obsessed with TTR (see below), we decided to build out a nicer deck of the
cards, with each card capturing a different, unique photo from our family.

This project was mostly about selecting the right set of photos from our
library, though I used some very basic facial detection APIs from Google to
make sure they were approximately centered within the boundaries of the frame.

![A new ticket to ride card, featuring myself and my sister on a purple background, representing a purple train card](/img/ttr-card2.png)

![Another TTR train card.](/img/ttr-card3.png)

My mother still refuses to use these cards - to my chagrin - because "she hasn't fully worn through the other deck yet" - in case you thought I was overstating the degree of her obsession.

### 2021 - Webapp 

One last time I dipped into the TTR well by creating an interactive version of the
custom board game using the simplest thing you can imagine: global 
state within a Java binary, running on a server on GCP with open connections,
once-per-second polling, and start/stop mechanics. Building out the bizlogic for 
turns and state proved challenging as I stepped on every rake that someone does
when creating a state based multiplayer game (validation, transition, etc.). A few
years later I'd read a book on generalized gameplay schemas and be struck by how
I'd made every mistake in the book. You live and learn though!

The end result was something that was fun for our family to play a few times, 
actually a decently responsive application, and a great reminder: building 
something hacky can actually get you reasonably far as long as you don't intend
for it to be around long term.

It was a good mother's day gift - though probably the time we spent on the phone while
we played it was the real one.

