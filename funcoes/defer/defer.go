package main

import "fmt"

func obterValorAprovado(v int) int {

	defer fmt.Println("Fim")

	if v >= 5000 {
		fmt.Println("Valor alto")
		return v
	} else {
		fmt.Println("valor abaixo")
		return v
	}
}

func main() {
	v := obterValorAprovado(5000)
	fmt.Println(v)
}
