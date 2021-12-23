package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")

	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)

}

func horaCertaJorge(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")

	fmt.Fprintf(w, "<h1>AOOOOOBA JORGE - Hora certa: %s</h1>", s)

}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	http.HandleFunc("/jorge", horaCertaJorge)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":4040", nil))
}
