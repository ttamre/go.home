package scrapers

import (
	"fmt"

	"github.com/gocolly/colly"
)

type RealtorListing struct {
	price      float32
	bedrooms   int
	bathrooms  int
	squareFeet int
	category   string
	address    string
	link       string
}

func ScrapeRealtor() {
	// Create a slice of RealtorListing structs
	var listings []RealtorListing

	fmt.Println("\nScraping data from Realtor...")
	c := colly.NewCollector()


	c.OnRequest(func(r *colly.Request) {
		fmt.Println("REQUEST:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("ERROR:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("RESPONSE:", string(r.Body))
	})

	// Called after OnResponse() if content is HTML
	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	url := "https://www.realtor.ca/map#ZoomLevel=9&Center=53.527308%2C-113.492490&LatitudeMax=54.38621&LongitudeMax=-110.43554&LatitudeMin=52.65062&LongitudeMin=-116.54944&Sort=6-D&PGeoIds=g30_c3x29657&GeoName=Edmonton%2C%20AB&PropertyTypeGroupID=1&TransactionTypeId=2&PropertySearchTypeId=0&Currency=CAD&HiddenListingIds=&IncludeHiddenListings=false"
	c.Visit(url)
	fmt.Println(listings)
}
