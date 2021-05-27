package getdata

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

/* Muestra los enlaces para descargar del anime que le indiques */

func Download(w http.ResponseWriter, r *http.Request) {
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

	c.OnHTML("td", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		if link != "" {
			response = append(response, link)
		}
	})

	c.Visit(base + URL)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	data, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
