+++

name = "Birth Certificate Scraping"
external_url = "https://github.com/gbdubs/birthcertificatescraping"

short_desc = "Scraping data from birth certificates between 1900 and 1950 for an Economics paper."
tags = ["scraping", "OCR", "matlab"]

+++

Professor Linda Bui is a gem who really helped me enjoy my economics degree at Brandeis.  So when she asked me if I could scrape all birth certificates from 1900 to 1950 in Massachusetts, I said “maybe”, rather than “no”.

I ended up using matlab to do OCR with a combination of models.  The forms were either hand written or were machine printed, and I generated seperate models for each.  The full accuracy of the system (having zero errors on the 21 fields) was about 40%.  80% got 18 of 21 fields correct.

Ultimately the algorithm I designed was too resource intensive to be runnable on the full set of data. I hadn’t considered that using a proprietary codebase like matlab would have the drawback that I wouldn’t be able to scale it on rented server time. 

I ended up passing this project off to a research assistant of Prof. Bui’s, and I don’t know what ended up happening to it.  
