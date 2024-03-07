package utils

import (
	"log"
	"strconv"
	"strings"
)

// takes a price as a string and returns it as a float32
func FormatPrice(price string) float32 {
	parsed, err := strconv.ParseFloat(price, 32)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return 0
	}
	return float32(parsed)
}

// takes a bedrooms string and returns it as an int
func FormatBedBathSqft(bedBathSqft string) int {
	parsed, err := strconv.Atoi(bedBathSqft)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return 0
	}
	return parsed
}

// removes the "for sale" from the category string along with any hyphens and extra spaces
func FormatCategory(category string) string {
	return strings.TrimSpace(strings.Replace(strings.Replace(category, "-", "", -1), "for sale", "", -1))
}
