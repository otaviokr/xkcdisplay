package main

import (
	"fmt"
	"image"
	"net/http"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gocolly/colly"
	"github.com/otaviokr/go-epaper-lib"
)

var (
	img image.Image
	epd *epaper.EPaper
)

func onHtmlCallback() (func(*colly.HTMLElement)) {
	return func(e *colly.HTMLElement) {
		imgSrc := e.ChildAttr("img", "src")
		fmt.Println(imgSrc)

		res, err := http.Get("http:" + imgSrc)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
		img, _, err = image.Decode(res.Body)
		// img, err = ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
	}
}

func showImage() {
	img = epd.Rotate(img)

	if (img.Bounds().Dx() > img.Bounds().Dy()) {
		img = imaging.Resize(img, 264, 176, imaging.Lanczos)
	} else {
		img = imaging.Resize(img, 264, 176, imaging.Lanczos)
	}

	epd.AddLayer(img, 0, 0, false)
	epd.PrintDisplay()
}

func main() {
	c := colly.NewCollector()
	c.AllowURLRevisit = true
	c.OnHTML("#comic", onHtmlCallback())

	var err error
	epd, err = epaper.New(epaper.Model2in7bw)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 3; i++ {
		c.Visit("https://c.xkcd.com/random/comic/")

		epd.Init()
		epd.ClearScreen()

		showImage()

		epd.Sleep()
		time.Sleep(30 * time.Second)
	}

	epd.Init()
	epd.ClearScreen()
}