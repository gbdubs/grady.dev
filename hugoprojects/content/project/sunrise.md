+++

order = 3998 

name = "Sunrise"
external_url = "https://grady.dev"

short_desc = "Capturing the best time of the day"
tags = ["art", "golang"]

+++

I love mornings - they're the best time of the day.
As a new day starts, all is possible, nothing has yet been decided.
Nature hums with unspoken anticipation - tension and potential bound up in stillness.

I also love sunrises - you get to lay claim to one bookend of the day,
and if seeing it being born doesn't allow you to sieze it, I can't imagine what could.

For several years I had the idea of trying to build a webpage that would capture my
love for this moment. I tinkered around with building gradients for it in Inkscape (
an SVG editor). I tried to build an algorithm for it based off of photonic refraction
for [Abound](../abound). But to no use. I could never get the colors right - even when
I was sampling from an image of a sunrise, the colors always felt either sterile or
gaudy. I realized that to emulate the real thing I needed data from the real thing.

## Data

I needed to get a timelapse of the sky on the morning of a perfect sunrise.

It took me four tries. Three times I found myself up before
dawn, perched somewhere in the foothills in Colordao, looking westward into the black.
But the sunrise is tricky - you have to get up far before Nautical Twilight to get in
position in order to catch it from black (which nescessarily implies using a day off) -
you have to have a timelapse setup that isn't too shaky - and critically, you get at
most one shot per day. I messed up the timelapse setup with one that got blown over once.

But the problem that thwarted me twice was clouds. Clouds are tricky. They
totally skew the data on colors because they introduce huge dips in saturation and
skew toward bluer hues even at sunrise. They're also ubiquitious.
If you look out to the horizon from a high vantage, only very rarely will there
truly be no clouds in the sky. When visibility is good, you need sometimes hundreds
of miles of clear sky in front of you to fully capture the color of the true horizon
during a sunrise. Twice I found myself thwarted in the wee hours of the morning by
clouds.

But then, on Junteenth 2025 (the Federal holiday most worthy of celebration for other reasons)
it happened. I set an alarm for 12:55 AM. I packed a bag and hit the road by 1:30 AM.
By 2:40AM I was in position at Rainbow Curve of Rocky Mountain National Park
(after a brief delay where the automated machines at the park entry wouldn't accept my money).
I setup the ersatz tripod (a laptop stand with a phone ductaped to it, reinforced by books),
using the direction of the
crescent on the moon to point toward the sun, which was coming from the northwest, instead of
the west as I'd planned. Rainbow Curve points due west through the foothills, which is why
I chose it, but this meant the sunrise would happen over the mountains. I went to sleep in my
car at 3:00 AM, and fitfully slept through howling winds until 6:30, when I opened the door to
bask in the warm glow of morning (the best time of the day). The video was perfect - I've included
it below

![The original video off of which this animation is based](/img/sunrise-video.mp4)

## Building the Gradient

First, I used Google Photos to stabilize the result, which really made subsequent analysis much easier.

Then, I created a mask which only included the sky (by hand in GIMP), and ran a bash script
to extract each frame of the image and apply that mask to it. I also put any of the frames
that happened after the sun poked over the horizon into a different folder, it was going to be
hard to solve both the pre-and-post sunrise with the same algorithm (I quickly realized.)

![An example of how I constructed a mask pixel by pixel to capture only the sunrise.](/img/sunrise-mask-1.png)

Using the data to create the output ended up being relatively simple, it took about 4-8h total.
Using Claude Code I tried out 4-5 different algorithms that I'd sketched out before to try to
achieve the subtle curvature you see in sunrises, including:

- Using Iso-Color bands to try to identify the center of a radial gradient based on curvature
- Using a loss function (of comparison with an orginal frame) to try to come up with an idealized gradient for each frame
- Using a vertical color sample from the middle and edges of the image to estimate the eccentricity and center of an elipse we could use as a radial gradient.

All of these proved not effective, and while I got some good failure modes, I wasn't happy with
progress toward them.

![A failed experiment](/img/sunrise-radial.png)

But the more I played around with it, the more it felt like maybe a linear gradient (a far simpler approach)
could capture it if I was able to use the mountains to break up the effect.

So the actual algorithm is dead simple: <b>for every Nth row in a frame of the video, average the colors,
and use that as the break point for that point in the gradient in the animation.</b> Easy.

There are only two additional effects I added after that analysis. The first was recognizing that the
sampling of individual rows meant that the matrix it generated wasn't "smooth" - both in the time
dimension and in the vertical dimension there was a lot of "roughness" in that the colors would jitter about
and you could see the bands clearly at keyframes. I solved this through a basic smoothing algorithm - just
average each keyframe value with the 8 adjacent keyframes (2 in the same frame, 3 in the next and 3 in the prev).

The second realization was that averaging RGB values yields colors that tend toward greys - think about the
3-space of the RGB cube to reason about why. HSL on the other hand doesnt suffer from this proprety, so I converted
the logic to average based on HSL, and that dramatically improved the color fidelity. In the image below,
the true image is on the left, the RGB averaged image is in the center, and the HSL averaged image (no other changes)
is on the right - what a difference that makes eh!?

![Three side by side images](/img/sunrise-color-averaging.png)

After figuring that out, it was just the small matter of tricking CSS into being willing to render an animated gradient,
because all the specs will tell you `linear gradients aren't animatable`. Because of this I started out by using SVG
animations, which worked great, but wasn't pausable/replayable, and that felt really important to get build in.

Well, it turns out if you use CSS' `@property` declarations with type color, and then animate
**all of the properties that impact the linear gradient in the same set of keyframes**,
CSS is willing to render a linear gradient as if it is animated. Many thanks to
[Temani Afif](https://dev.to/afif/we-can-finally-animate-css-gradient-kdk) for this hack.

My largest takeaway from this project was how effective Claude was at helping me build the image processing
algorithms. It would have taken me much more time and I would have had many more bugs (though Claude had 3-5)
I had to debug. The end result is gorgeous physically, but under the hood it's quite grotesque - code that
carries the scars of dozens of substantial revisions without much continuity of thought. An instructive co-programming
session.

## The End Result

I like it. I like watching it. It feels thoughtful - not overdone, not sloppy.

I have a list of things I'd like to try to get working in the future, including a true radial gradient
(in particular to capture the 3-5 second range of the video), getting the same view later in the year
(when the sun will rise in the gap in the mountains instead of over them), and trying a different view
to capture the impact of alpenglow.

I also feel ambivalent - I really liked the [old site](/oldindex/), and I haven't yet had a good
idea on how to translate the sunrise theme through to these project pages.

But this is good for now. The day is young and there is much to do!
