package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com a nota", nota)
	} else {
		fmt.Println("Reprovado com a nota", nota)
	}
}

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"

	} else if nota >= 7 && nota < 8 {
		return "B+"

	} else if nota >= 5 && nota < 7 {
		return "C"

	} else {
		return "E"

	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(4)

	fmt.Println(notaParaConceito(9.1))
	fmt.Println(notaParaConceito(1))
	fmt.Println(notaParaConceito(7.1))
}
