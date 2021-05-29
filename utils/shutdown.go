package utils

import (
	"log"
	"net/http"
	"os"
	"syscall"
	"time"
)

func Shutdown(w http.ResponseWriter, r *http.Request) {
	pidserver := os.Getpid()
	log.Println("Apagado")
	w.Write([]byte("Apagado"))
	time.Sleep(5 * time.Second)
	syscall.Kill(pidserver, 9)
}
