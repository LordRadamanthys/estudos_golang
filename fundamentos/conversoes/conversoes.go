package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notafinal := int(nota)

	fmt.Println(notafinal)

	//cuidado ..
	fmt.Println("teste " + string(123)) //printando a tabela asci

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(100 - num)

	b, _ := strconv.ParseBool(strings.ToLower("TrUe"))
	if b {
		fmt.Println(b)
	}
}
