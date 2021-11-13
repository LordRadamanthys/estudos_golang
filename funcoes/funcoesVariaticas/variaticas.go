package main

import "fmt"

func media(numero ...float64) float64 {

	total := 0.0

	for _, valor := range numero {
		total += valor
	}

	return total / float64(len(numero))
}

func main() {

	fmt.Printf("Média: %.2f", media(3.3, 4.4, 6.6, 9.9, 9.9, 10))
}
