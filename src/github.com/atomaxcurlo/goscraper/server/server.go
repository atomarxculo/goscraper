package server

import (
	"log"
	"net/http"
	"os"

	"github.com/atomarxculo/goscraper/getdata"
	"github.com/atomarxculo/goscraper/ping"
)

/* 	Inicia un servidor en un puerto que hayamos definido como variable de entorno
en caso de no tener una, por defecto lo inicia en el puerto 7171.
Tiene definida dos rutas, "ping", para comprobar si el servidor está funcionando y
"search", donde le pasaremos el anime a buscar.
*/

func StartServer() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "7171"
	}

	http.HandleFunc("/search", getdata.Anime)
	http.HandleFunc("/base", getdata.GetData)
	http.HandleFunc("/test", getdata.Download)
	http.HandleFunc("/", ping.Ping)

	log.Println("listening on", ":"+PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
