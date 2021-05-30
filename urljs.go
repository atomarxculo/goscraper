package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://jkanime.net/super-cub/", g.Opt.ParseFunc)
		},
		ParseFunc: parser,
	}).Start()
}

func parser(g *geziyor.Geziyor, r *client.Response) {
	URL := r.Request.URL.String()
	log.Println("Enlace:", URL)
	if URL == "https://jkanime.net/super-cub/" {
		log.Println("Funciona")
		r.HTMLDoc.Find("div.anime__item").Each(func(_ int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			log.Println("Enlace:", link)
		})
	} else {
		log.Println("URL equivocada")
		main()
	}
}
