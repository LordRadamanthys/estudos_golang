package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("numero invalido %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

// uint garante inteiros positivos
func fatorial1(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * fatorial1(n-1)

}

func main() {
	resultado := fatorial1(5)

	fmt.Println(resultado)
}
