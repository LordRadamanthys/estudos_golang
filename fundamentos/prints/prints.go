package main

import "fmt"

func main() {

	fmt.Print("Teste")
	fmt.Print("Teste")

	println("Nova ")
	println("linha")

	x := 2.141516

	// fmt.Println("o valor de x é"+x)
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.0990999
	c := false
	d := "opa"

	fmt.Printf("\nVariavel a: %d \nVariavel b: %f\nVariavel c: %t\nVariavel d: %s\n", a, b, c, d)
	fmt.Printf("%v %v %v %v", a, b, c, d)
}
