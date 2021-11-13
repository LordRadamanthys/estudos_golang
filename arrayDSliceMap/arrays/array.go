package main

import "fmt"

func main() {
	notas := []float64{7.3, 7.2, 9.9}

	// fmt.Println(notas)

	// notas[0], notas[1], notas[2] = 7.3, 7.2, 9.9

	fmt.Println(notas)

	fmt.Printf("Média %.2f", media(notas))

}

func media(arr []float64) float64 {
	var total float64 = 0.0

	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}

	result := total / float64(len(arr))
	return result
}
