+++

order = 1725

name = "Sniff"

short_desc = "Why _Florida v. Harris_ should be the minimum bar for surveillance"
tags = ["4th-ammendemnt", "privacy", "law"]

+++

Justice Kagan's decision in [Florida v. Harris](https://en.wikipedia.org/wiki/Florida_v._Harris#Decision) articulates a few principles on which we can build effective privacy preserving tools for law enforcement. In the unanimous decision, the court ruled that the indication of a Drug Sniffing Dog was grounds for search under the fourth amendment. In doing so, the court held up the drug sniffing dog as an exemplar of privacy preserving law enforcement tooling. It's worthwhile extracting a few lessons from the court's reasoning in considering the design of what makes a dog's nose, in contrast to most digital technology, such a satisfying balance of rights and state interests.

A key consideration that the court addresses is that of _statistically derived reliability_ as assessed through false-positive rate. The court argues that there is likely some threshold of false positive/negative at which a search becomes unjustified. This is easy to see in the case of a drug sniffing dog because we have no other means by which to reason about the dog's judgement. Outcomes are all that we get to see. 

As I've described in my thoughts on [Terry](../terry), a balancing act of the competing harms in the case of search and seizure can't be effectively assessed by an individual case

Unfortunately we shed this common sense when applying the same logic to human judgement. Judges are not blinded by the personhood and autonomy of the dog in evaluating it's impact on competing rights claims. Dogs can't do post-hoc rationalization, a dog's thought process cannot be interrogated by the law. Kagan sees that in Harris - there's no sensible way of looking at probable cause (a question of prior knowledge) that relies solely on whether the dog is correct in their indication of drugs (knowledge derived after search). 

By talking about the dog's perception in terms of false positive rate, Kagan concedes that explanation is _not a necessary component of an acceptable test for a decision criterion that results in search or seizure_. This is in stark contrast to the way that the court looks at officer discretion in Terry stops. In that case, because we can ask a human for an explanation, we put extensive weight on the explanation, without falling back or ever examining the far more compelling indica of rights-balancing - statistically derived accuracy.

In _Harris_, Kagan is making an argument about the dog that we should be making about any tool of law enforcement discernment (whether it is a drug sniffing dog, a UV filter, or a cop's judgement). All devices used to inform choices of search and seizure should be held against a rights balancing framework that analyzes real world statistical performance. Human judgement is a force that is better inspected from the cold evidence of outcomes, rather than the squishy (and time-skewed) process of post-hoc rationalization.

The second key consideration is the concept of _targeted information isolation_. A drug sniffing dog is processing information about a wide array of olfactory inputs, but only is expressing a verdict on a narrow set of targeted criteria. It's likely that a dog can smell a wide array of contents in the baggage of passengers, but the only information it expresses is that of the query _does this bag contain drugs_. This is supremely privacy preserving, because the only data that the system outputs (or even has the _capacity_ to output) is the question of interest, and the test is one the state has a compelling interest in. If dogs could talk, they could tell us a wide array of information about the contents of passerby's suitcases. In that instance, the dog would serve as a broader means of surveillance than is proscribed by a legitimate law enforcement purpose. Good law enforcement tooling should fit this principle: limit the potential output of the system to the outcomes that are of state interest.

The third and final thought that _Harris_ brings up is the concept of interaction between tool and law. A consistent frustration for groups that try to combat stop-and-frisk policing is how officers adapt to court rulings about what constitutes articulable/reasonable suspcicion, and start to use that terminology in their reports. "The odor of Marijuana" is the classic example of this. One additional (unarticulated) benefit of a drug dog is that it is more immune from this interaction pattern. Dogs will be biased by their training, but their individualized decisions are "individualized" in the words of _Terry_ in a way that human's thought processes can never really be. An officer's awareness of the law will always color both their decisions and how they describe their decisions. Man's best friend does not suffer from either of these defects.
