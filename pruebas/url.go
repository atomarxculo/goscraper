package pruebas

import (
	"fmt"

	"github.com/gocolly/colly"
)

func url() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only root url and urls which start with "e" or "h" on httpbin.org
		colly.URLFilters(
		//regexp.MustCompile("http://zippyshare\\.com/(|.+)$"),
		//regexp.MustCompile("http://httpbin\\.org/h.+"),
		),
	)

	// On every a element which has href attribute call callback
	/* c.OnHTML("a [href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		fmt.Println("Enlace:" + link)
		// Visit link found on page
		// Only those links are visited which are matched by  any of the URLFilter regexps
		c.Visit(e.Request.AbsoluteURL(link))
	}) */

	var response []string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			response = append(response, link)
			fmt.Println(link)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on http://httpbin.org
	c.Visit("https://jkanime.net/dragon-ie-wo-kau/")
}