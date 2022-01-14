# xkcdisplay
Fetch a random XKCD strip and show it on an e-paper display.

This is a fun little demonstration on how to use [go-epaper-lib](github.com/otaviokr/go-epaper-lib) and, at a much smaller scale, [colly](github.com/gocolly/colly).

This app will fetch a [XKCD](xkcd.com) random strip, try to make it fit the display and show it for a small period... and then repeat it.

I used a 2.7 e-paper display, but that is painfully too small to actually read the strip. I suggest you tried it with a bigger screen (some adjustments in the code are required, of course).

If I get a bigger screen, I'll run the tests and adapt this code to use it correctly - for the time being, I would say this is a successful proof-of-concept.
