package scrapers

import (
	"fmt"
	"github.com/gocolly/colly"
)

type RedfinListing struct {
	price      string
	bedrooms   string
	bathrooms  string
	squareFeet string
	category   string
	address    string
	link       string
}


func ScrapeRedfin(city string, province string) {
	// Create a slice of RedfinListing structs
	fmt.Println("\nScraping data from Redfin...")

	var listings []RedfinListing
	host := "https://www.redfin.com"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("REQUEST:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("ERROR:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("RESPONSE:", r.StatusCode)
	})

	// Called after OnResponse() if content is HTML
	c.OnHTML("body", func(e *colly.HTMLElement) {
		listing := RedfinListing{}

		// link
		e.ForEach(`a.link-and-anchor[href]`, func(_ int, el *colly.HTMLElement) {
			listing.link = host + el.Attr("href")
		})
		// price
		e.ForEach(`span.bp-Homecard__Price--value`, func(_ int, el *colly.HTMLElement) {
			listing.price = el.Text
		})

		// bed bath sqft
		e.ForEach(`div.bp-Homecard__Stats`, func(_ int, el *colly.HTMLElement) {
			listing.bedrooms = el.ChildText(".bp-Homecard__Stats--beds")
			listing.bedrooms = el.ChildText(".bp-Homecard__Stats--baths")
			listing.bedrooms = el.ChildText(".bp-Homecard__Stats--sqft")
		})

		// address
		e.ForEach(`div.bp-Homecard__Address`, func(_ int, el *colly.HTMLElement) {
			listing.address = el.Text
		})

		listings = append(listings, listing)
	})

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	url := fmt.Sprintf("%s/%s/%s", host, province, city)
	c.Visit(url)
	fmt.Println(listings)
}
