package main

import "fmt"

func imprimirAprovados(aprovados ...string) {

	fmt.Println("Lista de aporovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}

}

func main() {
	aprovados := []string{"Mateus", "Jorge", "Wilson"}

	imprimirAprovados(aprovados...)

	//teste slice com map
	alunos := []map[string]string{
		{"jorge": "sei la",
			"teste": "teste"},
	}

	fmt.Println(alunos)
}
