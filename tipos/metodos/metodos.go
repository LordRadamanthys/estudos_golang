package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeSobrenome(n string) {
	partes := strings.Split(n, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	pessoa := pessoa{"Mateus", "Lima"}
	fmt.Println(pessoa.getNomeCompleto())

	pessoa.setNomeSobrenome("Jorginho trazADoze")
	fmt.Println(pessoa.getNomeCompleto())
}
