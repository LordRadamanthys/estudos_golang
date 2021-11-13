package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"gabriela silva": 868.89,
			"Gomes silva":    868.89,
		},
		"M": {
			"Mateus": 15000.00,
			"Marcos": 15000.00,
		},
		"N": {
			"Nathalia": 15000.00,
			"Nilce":    15000.00,
		},
	}

	// delete(funcsPorLetra, "N")

	for _, funcionario := range funcsPorLetra {
		for nome, valor := range funcionario {
			fmt.Printf("Nome: %s - Valor: %f\n", nome, valor)
		}
	}
	// fmt.Println(funcsPorLetra)
}
