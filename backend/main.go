package main

import "github.com/ttamre/go.home/scrapers"

func main() {
	scrapers.ScrapeZillow()
	scrapers.ScrapeRedfin()
	scrapers.ScrapeRealtor()
}
