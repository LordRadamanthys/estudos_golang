package main

import (
	"fmt"
)

// NAO TEM TERNARIO
func nTemTernario(nota int) string {
	if nota > 6 {
		return "Passou"
	}
	return "Ih Rapaz...."
}

func main() {

	fmt.Println("Status:", nTemTernario(1))
}
