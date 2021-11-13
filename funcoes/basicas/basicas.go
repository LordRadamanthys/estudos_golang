package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("%s e %s", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "parametro1", "parametro2"
}

func main() {
	f1()
	f2("teste", "teste")

	r3, r4 := f3(), f4("teste", "teste1")
	fmt.Printf("\nresultado fun3: %s\nresultado fun4: %s\n", r3, r4)

	r5p1, r5p2 := f5()

	fmt.Println(r5p1)
	fmt.Println(r5p2)
}
