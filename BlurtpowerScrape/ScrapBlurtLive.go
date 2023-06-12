package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Product struct {
	Name    string
	PostAge string
	Url     string
}

func main() {

	url := []string{
		"https://ecosynthesizer.com/blurt/@oswaldotorres88",
		"https://ecosynthesizer.com/blurt/@outofthematrix",
		"https://ecosynthesizer.com/blurt/@perceptualflaws",
		"https://ecosynthesizer.com/blurt/@petrapurple",
		"https://ecosynthesizer.com/blurt/@phasewalker",
	}

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.SetRequestTimeout(100 * time.Second)
	products := make([]Product, 0)

	// Callbacks

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			e.ForEach("td", func(i int, g *colly.HTMLElement) {
				item := Product{}
				item.Name = h.Text
				item.PostAge = g.Text
				item.Url = g.ChildAttr("a", "href")
				products = append(products, item)
			})
		})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(products, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("products.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}
	})
	for _, i := range url {
		c.Visit(i)
	}
	c.Wait()
}
