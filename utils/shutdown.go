package utils

import (
	"log"
	"net/http"
	"os"
	"syscall"
)

func Shutdown(w http.ResponseWriter, r *http.Request) {
	pidserver := os.Getpid()
	w.Write([]byte("Apagado"))
	log.Println("Apagado")
	syscall.Kill(pidserver, 9)
}
