package getdata

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

var fullURL string
var response []string

func Search(w http.ResponseWriter, r *http.Request) {
	base := "https://jkanime.net/"
	URL := r.URL.Query().Get("anime")
	fullURL = base + URL + "/"
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", base+URL)
	started()

	data, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
	response = nil
}

func started() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(fullURL, g.Opt.ParseFunc)
		},
		ParseFunc: parser,
	}).Start()
}

func parser(g *geziyor.Geziyor, r *client.Response) {
	enlace := r.Request.URL.String()
	if enlace != fullURL {
		log.Println("URL equivocada")
		started()
	} else {
		r.HTMLDoc.Find("div.anime__item").Each(func(_ int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			response = append(response, link)
		})
	}
}
