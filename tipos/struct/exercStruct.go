package main

import "fmt"

type aluno struct {
	nome  string
	idade int
	nota  float64
}

type sala struct {
	numSala int
	alunos  []aluno
}

func newAluno() aluno {
	return aluno{}
}
func (al aluno) charge() []aluno {
	aluno1 := aluno{
		nome:  "Jorge",
		idade: 21,
		nota:  10,
	}

	aluno2 := aluno{
		nome:  "Marcia",
		idade: 18,
		nota:  7.9,
	}

	aluno3 := aluno{
		nome:  "Aline",
		idade: 19,
		nota:  9.3,
	}

	return []aluno{aluno1, aluno2, aluno3}
}

func main() {
	a1 := newAluno()
	alunosArr := a1.charge()
	sala1 := sala{
		numSala: 1,
		alunos:  alunosArr,
	}

	a2 := aluno{nome: "Maria do bairro", idade: 18, nota: 2.9}
	a2.nome = "Lidia"
	sala2 := append(sala1.alunos, a2)
	fmt.Println(sala2)
}
