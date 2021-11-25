package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	// passando o endereço de memoria da variavel para a func
	// assim podendo incrementar a variavel fora do escopo
	inc2(&n)
	fmt.Println(n)
}
