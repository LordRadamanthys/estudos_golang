package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	//numeros inteiros
	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	// só positivos= uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("o byte é", reflect.TypeOf(b))

	// valores com sinal int8,int16,int32, int64
	i1 := math.MaxInt64

	fmt.Println("ovalor maximo de i1 é", i1)
	fmt.Println("o tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode (int32)
	fmt.Println(i2)

	// numeros flutuantes
	var x float32 = 49.99
	// var x  = float32(49.99)
	fmt.Println("o tipo de x é", reflect.TypeOf(x))

	//boolean
	bo := true

	fmt.Println("O tipo da variave bo é", reflect.TypeOf(bo))

	//string
	s1 := "Ola string"

	fmt.Println(s1 + "!!")
	fmt.Println("O tamanho da variavel s1 é", len(s1))

	s2 := `ola
	meu
	nome
	é mateus`

	fmt.Println(s2)

	char := 'a'
	fmt.Println("o tipo char é", reflect.TypeOf(char))
}
