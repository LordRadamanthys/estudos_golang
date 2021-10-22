package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma=> ", a+b)
	fmt.Println("Subtracao=> ", a-b)
	fmt.Println("Divisao=> ", a/b)
	fmt.Println("Multiplicacao=> ", a*b)
	fmt.Println("Modulo/resto=> ", a%b)

	//bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)

	// operacoes com a aclasse math

	c := 3.0
	d := 2.0

	fmt.Println("Maior=> ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor=> ", math.Min(c, d))
	fmt.Println("Exponenciacao=> ", math.Pow(c, d))

}
