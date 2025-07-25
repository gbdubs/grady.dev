---
aliases: ["/project/sniff/"]
title: "Sniff"
subtitle: "Why _Florida v. Harris_ should be the minimum bar for surveillance"
date: 2024-01-01T00:00:00Z
draft: false

# Project classification
type: "projects"
tags: ["privacy", "writing", "government"]
isPriority: false 
sortPriority: 1725
isSimple: false

# Call to action
CTA: ""
CTALink: ""
CTAPreamble: ""

# Optional image
featured_image: ""

# Summary (used in lists)
summary: "Why _Florida v. Harris_ should be the minimum bar for surveillance"
---

In _[Florida v. Harris](https://en.wikipedia.org/wiki/Florida_v._Harris#Decision)_ (2012) a unanimous Supreme Court held that statistically derived evidence of efficacy of a drug sniffing dog was not a requirement for establishing probable cause for evaluating fourth amendment claims. I think this misses the mark from one of my favorite cases _[US v. Place](https://en.wikipedia.org/wiki/United_States_v._Place)_ (1983), in which O'Connor issued this gem of straightforward and balanced reasoning:

"A "canine sniff" by a well-trained narcotics detection dog, however, does not require opening the luggage. It does not expose noncontraband items that would otherwise remain hidden from public view, as does, for example, an officer's rummaging through the contents of the luggage. Thus, the manner in which information is obtained through this investigative technique is much less intrusive than a typical search. Moreover the sniff discloses only the presence or absense of narcotics, a contraband item. Thus, despite the fact that the sniff tells the authorities about the contents of the luggage, the information obtained is limited. ... We are aware of no other investigative procedure that is so limited both in the manner in which the information is obtained and in the content of the information revealed by the procedure."

I really like this passage both because it provides an aspirational goal for law-enforcement, and because it articulates a clear requirement: that the decision procedure used must be found to be accurate to be reasonable (a fact not contended at the Supreme Court, but critical to O'Connor's reasoning.) It's worthwhile extracting a few lessons from the court's reasoning in considering the design of what makes a dog's nose, in contrast to most digital technology, such a satisfying balance of rights and state interests.

A key consideration that the court tangentially addresses in _Place_ is that of _statistically derived reliability_ as assessed through false-positive rate. There must be some threshold of false positive/negative at which a search becomes unjustified (or at which the decision procedure falls out of constructive use toward totality of the circumstances). This is easy to see in the case of a drug sniffing dog because we have no other means by which to reason about the dog's judgement. Outcomes are all that we get to see for a dog, so the question of "well-trained" is a throwaway line, rather than a core holding.

This ties into my thoughts on [Terry](../terry): a balancing act of competing harms in 4th amendment law can't be effectively assessed by an individual case.

Unfortunately we shed this common sense when answering the same questions about to human judgement. Judges are not blinded by the personhood and autonomy of the dog in evaluating it's impact on competing rights claims. Dogs can't perform post-hoc rationalization, a dog's thought process cannot be interrogated by the law. There's no sensible way of looking at probable cause (a question of prior knowledge) that relies solely on whether the dog is correct in their indication of drugs (knowledge derived after search). 

By talking about the dog's perception in terms of accuracy, O'Conner implies that explanation is _not a necessary component of an acceptable test for a decision criterion that results in search or seizure_. This is in stark contrast to the way that the court looks at officer discretion in Terry stops. In that case, because we can ask a human for an explanation, we put extensive weight on the explanation, without falling back or ever examining the far more compelling indication of rights-balancing - statistically derived accuracy.

In _Harris_, Kagan misses an argument about the dog that we should be making about any tool of law enforcement discernment (whether it is a drug sniffing dog, a thermal imager, or a cop's judgement). All decision criteria used to inform choices of search and seizure should be held against a rights balancing framework that analyzes real world statistical performance. Human judgement is a force that is better inspected from the cold evidence of outcomes, rather than the squishy (and time-skewed) process of post-hoc rationalization.

The second component that makes me love O'Connor and this decision is the concept of _targeted information isolation_. A drug sniffing dog is processing information about a wide array of olfactory inputs, but only is expressing a verdict on a narrow set of targeted criteria. It's likely that a dog can smell a wide array of contents in the baggage of passengers, but the only information it expresses is that of the query _does this bag contain drugs_. This is supremely privacy preserving, because the only data that the system outputs (or even has the _capacity_ to output) is the question of interest, and the test is one the state has a compelling interest in. If dogs could talk, they could tell us a wide array of information about the contents of passerby's suitcases. In that instance, the dog would serve as a broader means of surveillance than is proscribed by a legitimate law enforcement purpose. Good law enforcement tooling should fit this principle: limit the potential output of the system to the outcomes that are of explicit and compelling state interest.

The third and final thought that _Place_ and _Harris_ bring up is the concept of interaction between tool and law. A consistent frustration for groups that try to combat stop-and-frisk policing is how officers adapt to court rulings about what constitutes articulable/reasonable suspcicion, and start to use that terminology in their reports. "The odor of Marijuana" is a powerful phrase. One additional (unarticulated) benefit of a drug dog (or more generally, a non-human system) is that it is more immune from this interaction pattern. Dogs will be biased by their training, but their individualized decisions are "individualized" in the words of _Terry_ in a way that human's thought processes can never really be. An officer's awareness of the law will always color both their decisions and how they describe their decisions. Man's best friend does not suffer from either of these defects.
