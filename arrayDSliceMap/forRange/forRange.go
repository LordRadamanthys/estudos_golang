package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, elemento := range numeros {
		fmt.Printf("%d - %d\n", i, elemento)
	}

	// ignorando indice e pegando só o elemento
	for _, elemento := range numeros {
		fmt.Println(elemento)
	}
}
