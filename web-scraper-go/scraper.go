package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

/*
OnRequest(): Called before performing an HTTP request with Visit().
OnError(): Called if an error occurred during the HTTP request.
OnResponse(): Called after receiving a response from the server.
OnHTML(): Called right after OnResponse() if the received content is HTML.
OnScraped(): Called after all OnHTML() callback executions.
*/

type PokemonProduct struct {
	url, image, name, price string
}

func main() {

	var pokemonProducts []PokemonProduct

	c := colly.NewCollector()

	//  Note that the e parameter of the callback function represents a single li.product HTMLElement
	// target li HTML element has the .product class
	c.OnHTML("li.product", func(e *colly.HTMLElement) {

		pokemonProduct := PokemonProduct{}

		// get the href attribute of the a element
		pokemonProduct.url = e.ChildAttr("a", "href")

		// get the src attribute of the img element
		pokemonProduct.image = e.ChildAttr("img", "src")

		// get the text of the h2 element
		pokemonProduct.name = e.ChildText("h2")

		// get the text with the class .price
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)

	})

	c.Visit("https://scrapeme.live/shop/")

	// opening the csv file
	file, err := os.Create("pokemon.csv")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}
	defer file.Close() // Close closes the File, rendering it unusable for I/O

	// initializing a file writter
	writer := csv.NewWriter(file)

	// defining the CSV headers
	headers := []string{"url", "image", "name", "price"}

	// writting the colume headers
	writer.Write(headers)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}

		writer.Write(record)
	}
	defer writer.Flush() // Flush writes any buffered data to the underlying io.Writer.

}
