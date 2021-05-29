package utils

import (
	"log"
	"net/http"
)

/* Escribe "ping" por consola cuando consultamos el servidor */

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping")
	w.Write([]byte("ping"))
}
