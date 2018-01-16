+++

name = "Find A Clinic"
external_url = "http://www.findaclinic.org"

short_desc = "Find a Clinic is the best!"
tags = ["webapp", "collaborative"]

+++

Please note - I deeply respect various opinions on the moral wight of abortion, but personally [^1] fully support the right of a woman to make her own judgements about the matter. For that reason, I believe that women should have full knowledge of all of their options when they become pregnant, and that is the aim that this project is directed toward.

[^1] I think most of us would agree that there is a difference in value between the existence of a day-old fertilized human egg, and a pregnancy that is about to be birthed. This implies that there is a spectrum of moral culpability along this timeline. I believe that because the challenge of obtaining the right ethical point at which the act of terminating a pregancy is unacceptable, the person with the most knowledge and proximity (the woman who is preganant) is ultimately the best informed to be making the decision as to whether or not an abortion is morally acceptable.

In 2017 there was an unprecedented attack on a woman’s right to obtain a safe and legal abortion in the United States.  Many states have used arbitrary (non-health driven) regulations to try to force abortion clinics to close.  As a result, many women live in “clinic-deserts”, where no clinics exist within a reasonable radius.  These women must choose between birthing a child that they are unprepared for and finding their way to an abortion clinic which might be hundreds of miles away. 

Find a Clinic was an effort by myself and four others to try to make this situation clearer and enumerate all of the options for women in this dangerous situation.  Our goal was to collect information on the clinics in the United States, the abortion regulations in every state, the costs of procedures at different locations and pull in APIs with current flight, train, and bussing options. This data was then synthesized into a set of options - different clinic/date pairs that would minimize cost, or get around certian regulations like partner consent restrictions.

For this effort, I was in charge of building a robust backend that could handle the entry and exposition of clinic data. 

This effort ended up failing. There are many factors that led to this failure, but the main one was problems around data collection. We had overestimated the willingness of clinics to give us cost and procedure information, and the willingness of volunteers to call to ask for it during working hours.  In the end, we built a fairly robust system, with a gorgeous front end and reasonably well built back end, only to have it fail because some of our fundamental assumptions were flawed. Once the problems with data collection became apparent, our group found it signifigantly harder to concentrate on the project, and we ended up disbanding a few weeks after this problem came to light.

However, I am deeply proud of this work, and hopeful to work on things like it in the future. 