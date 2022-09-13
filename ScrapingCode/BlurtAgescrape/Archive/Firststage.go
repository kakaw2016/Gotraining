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
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	products2 := make([]Product, 0)

	// Callbacks

	c.OnHTML("div.PostFull__header", func(e *colly.HTMLElement) {
		e.ForEach("h1.entry-title", func(i int, h *colly.HTMLElement) {
			e.ForEach("div.right-side", func(i int, g *colly.HTMLElement) {
				item := Product{}
				item.Name = h.Text
				item.PostAge = g.Text
				item.Url = "https://blurt.one" + g.ChildAttr("a", "href")
				products2 = append(products2, item)
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
		js, err := json.MarshalIndent(products2, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("products.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}

	})

	c.Visit("https://blurt.one/wintness/@zahidsun/how-to-deposit-and-withdraw-blurt-on-pancakeswap")
}
