package getdata

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/atomarxculo/goscraper/array"
	"github.com/gocolly/colly"
)

/* Esta funcion consulta a la web el anime que le busques,
además de mostrarte los enlaces en orden y que no haya duplicados*/

func Anime(w http.ResponseWriter, r *http.Request) {
	base := "https://jkanime.net/"
	baseexp := "https://jkanime\\.net/"
	//Verify the param "URL" exists
	URL := r.URL.Query().Get("anime")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", base+URL)

	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(baseexp + "(|" + URL + ".+)$"),
		),
	)

	var response []string
	var result []string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
		if strings.Contains(link, URL) {
			response = append(response, link)
		}
		result = array.RemoveDuplicates(response)
		sort.Strings(result)
	})

	c.Visit(base)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	data, err := json.Marshal(result)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)

}
