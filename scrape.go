package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

//web scraping function, taking facts from website and returning fact slice (ID, fact)
func ScrapeFacts() []Fact {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)
	//telling scaper where to look
	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}

		dogFact := element.Text

		fact := Fact{
			ID:   factId,
			Fact: dogFact,
		}
		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	collector.Visit("https://www.factretriever.com/dog-facts")
	return allFacts
}
