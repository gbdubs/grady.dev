+++

order = 1400

name = "Terry v. Ohio + Aggregated Infringement"

short_desc = "Case level analysis sometimes cannot arrive at agreeable precedent."
tags = ["law", "privacy"]

+++

_Terry v. Ohio_ is known for it's implications on the probable cause requirement of the fourth amendment, but re-reading the decision over the weekend, something else jumped out at me when I put myself in the shoes of the Warren court: given the fact pattern, they were doing their best to encourage good policework, and in the case of stop and frisk the differentiation between 'good' and 'bad' policing hinges on the quality of the officer's decision making and the outcome of their choice:

The majority tries to write narrowly. I'd guess that they thought that the correct outcome was to uphold the search and they were grasping for a reason why the stop could be justified without destroying probable cause. If so, I see where they are writing from. I'm at all concerned with Terry's fate, but I definitely care about supporting the McFaddens of the world. McFadden made a judgement call to stop and frisk Terry and his coconspirators, and out of this individual decision, an agreeable outcome arose. Were McFadden less perceptive, more suspicious, or less discriminating we wouldn't be talking about _Terry_. But as it happened, the case that justifies Stop and Frisk is a _true positive_.

Zoom back for a second. An officer's decision on whether to Stop and Frisk (S+F) a civilian has four outcomes, based on whether the officer decides to stop (STOP) or not (PASS), and whether the civilian is actually in possession of illegal goods or arms (GUILTY) or not (INNOCENT): 

(1) True Negative (PASS, INNOCENT) 
(2) False Negative (PASS, GUILTY) 
(3) False Positive (STOP, INNOCENT) 
(4) True Positive (STOP, ARREST)

There are clear downsides to False Positives and False Negatives. __False positives__ seriously intrude on basic liberties to be secure in person and posession, tend to be racially biased, and create mistrust between law enforcement and civilians. __False negatives__ likely increase crime rates and embolden criminals. A city's choice on how and if to conduct S+F __should probably be based on a concious tradeoff between the relative incidence of false positives and false negatives__.

But thats now how courts operate. In _Terry_, we only see a single instance, a True Positive. Were the circumstances in front of the Warren court not that of McFadden and Terry, but, say, Sandra Bland and Brian Encinia (a rage inducing _False Positive_), the court might arrive at a different outcome. Heck, if Officer McFadden had stopped 100 innocent people before stopping Terry, I'd argue that should change the outcome, but we don't get to know that context in deciding this case. The problem with _Terry v. Ohio_ is not (as some argue) that the fact pattern sits at the peak of the watershed 'probable' divide (though it does). Instead, the problem with _Terry_ is that it is a single case.

When constitutional rights are in conflict, courts perform our balancing. When the parties with competing constitutional claims are on opposite benches, performing this balancing is easier than when no single case can capture both of the claims. In thinking about S+F, the fact pattern in front of us heavily influences our intuition: true positives show that there are occasions when we want to defer to officer discretion, and false positives show that officer discretion should not be unlimited.

It would be impractical and unwise to conclude based on this sentiment that case by case analysis here can yield just outcomes. Repairing 4th amendment violations is fraught. Instead, we should take a wholistic approach to justice, outside of the confines of a fact pattern to bias our intuition. It is feasible to evaluate the actual numbers of the four outcomes above in any given city. The Positives just need record keeping, and the Negatives  need a few dozen (fully random) stops per day as controls (ignoring the clear constitutional problems with that practice). From there we could evaluate the harms of both groups in direct comparison (consider targeting a 4:1 FalsePositive to FalseNegative ratio). 

What if a city is above this target ratio? Then officers should exercice more discretion before making a stop: instruct officers to have a higher degree of confidence before chosing to S+F! Rotate cops off of S+F if their ratio is bad! Provide more training on differentiating between suspicious and atypical behavior! Provide pay incentives for good ratios! Similar remedies could be imagined for a ratio below the target.

These are all specific recommendations to a very general and very hard problem: how do you improve the results (say S+F ratio) of a blackbox system (such as the brain of an officer) that can't be reliably inspected? __You define a cost function, and instruct the system to self-optimize__. Officers know how to police far better than I (or most judges) do. Rather than defining a set of operational rules (that would be near impossible to enforce or evaluate) for how an officer should decide whether to perform a S+F, I'd rather we be honest about the aggregated harms of over or under stopping, and encode our relative thoughts about both in a function that can then be actualized by the experts.

More broadly, re-reading _Terry_ I'm reminded that cases are instructive, but single-case rulings probably wont consistently yield just outcomes. This has me thinking about a ton of topics: incremental rights and incremental infringement, actor driven optimization, and the boundaries of case law for achieving social justice.
