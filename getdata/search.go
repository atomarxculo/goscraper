package getdata

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func Search(w http.ResponseWriter, r *http.Request) {
	base := "https://jkanime.net/"
	//Verify the param "URL" exists
	URL := r.URL.Query().Get("anime")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", base+URL)

	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(base+URL, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("div.anime__item").Each(func(_ int, s *goquery.Selection) {
				log.Println(s.Find("a").Text())
			})
		},
	}).Start()
}
