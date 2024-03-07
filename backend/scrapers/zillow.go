package scrapers

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/ttamre/go.home/utils"
)

type ZillowListing struct {
	price      float32
	bedrooms   int
	bathrooms  int
	squareFeet int
	category   string
	address    string
	link       string
}

func ScrapeZillow() {
	// Create a slice of ZillowListing structs
	var listings []ZillowListing

	// Visit Zillow
	fmt.Println("\nScraping data from Zillow...")
	c := colly.NewCollector()


	c.OnRequest(func(r *colly.Request) {
		fmt.Println("REQUEST:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("ERROR:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("RESPONSE:", r.Request.URL)
	})

	// Called after OnResponse() if content is HTML
	c.OnHTML("article", func(e *colly.HTMLElement) {
		listing := ZillowListing{}

		// Get the price
		fmt.Println("- grabbing price")
		e.ForEach(`span[data-test="property-card-addr"]`, func(_ int, el *colly.HTMLElement) {
			listing.price = utils.FormatPrice(el.Text)
		})

		// Get the category, bedrooms, bathrooms, and square feet
		// optional class selector: div.StyledPropertyCardDataArea-c11n-8-84-3__sc-yipmu-0
		fmt.Println("- grabbing category")
		e.ForEach("div.dbDWjx", func(_ int, el *colly.HTMLElement) {
			// category
			listing.category = utils.FormatCategory(el.Text)

			// bedrooms, bathrooms, and square feet
			fmt.Println("- grabbing bed bath sqft")
			e.ForEach("ul.eYPFID", func(_ int, el *colly.HTMLElement) {
				e.ForEach("li", func(i int, child *colly.HTMLElement) {
					switch i {
					case 0:
						listing.bedrooms = utils.FormatBedBathSqft(child.Text)
					case 1:
						listing.bathrooms = utils.FormatBedBathSqft(child.Text)
					case 2:
						listing.squareFeet = utils.FormatBedBathSqft(child.Text)
					}
				})
			})
		})

		// Get the address from an address element with the class "address"
		fmt.Println("- grabbing addr")
		e.ForEach(`address[data-test="property-card-addr"]`, func(_ int, el *colly.HTMLElement) {
			listing.address = el.Text
		})

		// Get the link from an element with the class "link"
		fmt.Println("- grabbing link")
		e.ForEach(`a[data-test="property-card-link"]`, func(_ int, el *colly.HTMLElement) {
			listing.link = el.Attr("href")
		})

		listings = append(listings, listing)
	})
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	c.Visit("https://www.zillow.com/homes/Edmonton,-AB_rb/")
	fmt.Println(listings)
}
