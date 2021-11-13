package main

import "fmt"

func main() {
	var aprovados = make(map[int]string)

	aprovados[12345] = "maria"
	aprovados[12346] = "jorge"
	aprovados[12347] = "miriam"

	fmt.Println(aprovados)

	for key, value := range aprovados {
		fmt.Printf("%d = %s\n", key, value)
	}

	fmt.Println(aprovados[12345])
	delete(aprovados, 12345)
	fmt.Println(aprovados)
}
