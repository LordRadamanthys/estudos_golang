package main

import "fmt"

func main() {
	funcionarios := map[string]float64{
		"Mateus":  1799.98,
		"Jorge":   12342.77,
		"Marcia":  12310.27,
		"Tatiane": 1222.77,
	}

	funcionarios["Romoaldo"] = 1350.00
	delete(funcionarios, "naoExiste")

	for nome, salario := range funcionarios {
		fmt.Printf("Nome: %s - Salario: %f\n", nome, salario)
	}

}
