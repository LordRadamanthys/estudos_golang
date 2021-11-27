package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type nissan struct {
	carro
	turboLigado bool
}

func main() {
	n := nissan{}

	n.nome = "GTR"
	n.turboLigado = true
	n.velocidade = 50
	text := ""
	if n.turboLigado {
		text = "ligado"
	} else {
		text = "desligado"
	}
	fmt.Printf("O Nissan %s esta a %d km/h e com o turbo %s", n.nome, n.velocidade, text)
}
