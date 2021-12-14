package main

import (
	"io"
	"net/http"
)

func get() chan interface{} {
	c := make(chan interface{})
	go func() {
		resp, _ := http.Get("https://servicodados.ibge.gov.br/api/v1/localidades/distritos")
		body, _ := io.ReadAll(resp.Body)
		c <- body
	}()

	return c
}

// func main() {
// 	v := get()
// 	fmt.Println(<-v)
// }
