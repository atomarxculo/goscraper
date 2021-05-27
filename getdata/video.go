package getdata

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/atomarxculo/goscraper/array"
	"github.com/gocolly/colly"
)

/* Muestra los enlaces para descargar del anime que le indiques */

func Video(w http.ResponseWriter, r *http.Request) {
	base := "https://jkanime.net/"
	//Verify the param "URL" exists
	URL := r.URL.Query().Get("anime")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", base+URL)

	c := colly.NewCollector()

	var response []string
	var result []string

	c.OnHTML("iframe [src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		c.Visit(e.Request.AbsoluteURL(link))
		if link != "" {
			response = append(response, link)
		}
		result = array.RemoveDuplicates(response)
		sort.Strings(result)
		log.Println(response)
	})

	c.Visit(base + URL)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	b, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}
