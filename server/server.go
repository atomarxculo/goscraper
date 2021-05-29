package server

import (
	"log"
	"net/http"
	"os"

	"github.com/atomarxculo/goscraper/getdata"
	"github.com/atomarxculo/goscraper/utils"
)

/* 	Inicia un servidor en un puerto que hayamos definido como variable de entorno
en caso de no tener una, por defecto lo inicia en el puerto 7171.
*/

func StartServer() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "7171"
	}

	http.HandleFunc("/", utils.Ping)
	http.HandleFunc("/test", getdata.GetData)
	http.HandleFunc("/shutdown", utils.Shutdown)
	http.HandleFunc("/api/search", getdata.Search)
	http.HandleFunc("/api/download", getdata.Download)
	http.HandleFunc("/api/view", getdata.Video)

	log.Println("listening on", ":"+PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
