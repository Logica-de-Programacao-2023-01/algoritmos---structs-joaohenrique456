package main

import "fmt"

type Aluno struct {
	notas []float64
	nome  string
	media float64
	idade int
}

func nAluno(nome string, idade int, notas []float64) Aluno {
	var M float64
	for _, x := range notas {
		M += x
	}
	media := M / float64(len(notas))
	return Aluno{nome: nome, idade: idade, notas: notas, media: media}
}
func main() {
	notas := []float64{8.5, 1.3, 3.4, 5.9}
	aluno := nAluno("Carlos", 18, notas)
	fmt.Println(aluno)
}
