+++

order = 60

name = "Birth Certificate Scraping"
external_url = "https://github.com/gbdubs/birthcertificatescraping"

short_desc = "Scraping data from birth certificates between 1900 and 1950 for an Economics paper."
tags = ["scraping", "OCR", "matlab"]

+++

Professor Linda Bui is a gem who really helped me enjoy the last class of my economics degree at Brandeis.  So when she asked me if I could scrape all birth certificates from 1900 to 1950 in Massachusetts, I said “maybe”, rather than “no”.

I ended up using Matlab to do OCR with a combination of models for segmenting the certificates, and extracting the text for each field.  The forms were either hand written or were machine printed, and I generated seperate models for each.  The full accuracy of the system (having zero errors on the 21 fields) was about 40%.  80% got 18 of 21 fields correct.

Ultimately the algorithm I designed was too resource intensive to be runnable on the full set of data. I hadn’t considered that using a proprietary codebase like Matlab would have the drawback that I wouldn’t be able to scale it on rented server time.

I ended up passing this project off to a research assistant of Prof. Bui’s, and I don’t know what ended up happening to it.

One of the main things I learned from the project was that choice of the software the forms the building blocks of your system really matters. If you choose correctly, you are able to reap benefits in both implementation and scale, and if you choose for short term gains (i.e. Matlab's solid OCR library), you might end up paying the cost later.