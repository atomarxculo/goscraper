package main

import (
	"log"

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
	} else {
		log.Println("URL equivocada")
		main()
	}
}
