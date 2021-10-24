package main

import "fmt"

func main() {

	i := 1

	var p *int = nil

	p = &i //pegando o endereço da variavel i
	*p = 55
	fmt.Printf("ponteiro %v valor %v", p, *p)

}
