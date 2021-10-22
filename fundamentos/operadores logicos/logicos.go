package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {

	comprarTv50 := trab1 && trab2
	comprarTv23 := trab1 != trab2
	sorvete := trab1 || trab2

	return comprarTv50, comprarTv23, sorvete
}

func main() {
	tv50, tv23, sorvete := comprar(true, true)

	fmt.Printf("tv50: %t, tv23: %t, sorvete: %t, saudavel: %t", tv50, tv23, sorvete, !sorvete)
}
