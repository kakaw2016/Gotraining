package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gocolly/colly"
)

type Product struct {
	Name    string
	PostAge string
	Url     string
}

func writeJSON(data []Product) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Unable to create the JSON file.")
	}
	_ = ioutil.WriteFile("products.json", file, 0644)
	fmt.Println("Scraping and Writing successful. Go for Good!")
}

func main() {

	url := []string{
		"https://blurt.blog/blurtconnect/@anitagbemudu/why-do-teenagers-use-drugs",
		"https://blurt.blog/blurtlife/@animal-shelter/gender-roles-of-the-wheel-of-the-year-6-landlady-and-landlord",
		"https://blurt.blog/blurtlive/@blurt.live/blurt-live-curation-report-05-or-08-or-2022",
		"https://blurt.blog/blurtpower/@kamranrkploy/blurt-milestone-400k-blurtpower-reached",
		"https://blurt.blog/colorchallenge/@gduran/4koory-blurt-colorchallenge-thursday-green",
		"https://blurt.blog/epistem/@famigliacurione/self-empowerment-are-you-concerned-about-who-loves-and-who-hates-you",
		"https://blurt.blog/iduvts/@ahlawat/mobile-photoshoot-of-the-day-32",
		"https://blurt.blog/powerclub/@powerclub.curate/let-your-dream-shine-powerclub-curation-report-15",
		"https://blurt.blog/r2cornell/@alim236/let-us-come-under-the-shade-of-islam-and-protect-our-lives-and-families",
		"https://blurt.blog/r2cornell/@mahmud552/beautiful-photography-of-a-grass-flower-setaria-viridis",
		"https://blurt.blog/r2cornell/@mahmud552/it-is-not-right-to-block-religious-activities",
		"https://blurt.blog/r2cornell/@tanweeralam/brinjal",
		"https://blurt.blog/techclub/@techclub/daily-curation-report-53-for-techclub-community",
	}

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.SetRequestTimeout(100 * time.Second)
	products := make([]Product, 0)

	// Callbacks

	c.OnHTML(".PostFull__header", func(e *colly.HTMLElement) {
		//e.ForEach("h1.entry-title", func(i int, h *colly.HTMLElement) {
		//e.ForEach("a[href]", func(i int, g *colly.HTMLElement) {
		//test := e.DOM.Find("entry-title").Eq(0)

		item := Product{}

		item.Name = e.ChildText("h1.entry")
		item.PostAge = e.ChildAttr("#right-side", "title")
		item.Url = "https://blurt.blog/@" + e.ChildText("#Author > span.author > strong > a")
		products = append(products, item)
	})
	//})
	//})

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
		/*js, err := json.MarshalIndent(products, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file")
		if err := os.WriteFile("products.json", js, 0664); err == nil {
			fmt.Println("Data written to file successfully")
		}*/
		writeJSON(products)
	})
	for _, i := range url {
		c.Visit(i)
	}
	c.Wait()
}
