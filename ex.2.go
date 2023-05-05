package main

import "fmt"

type livro struct {
	titulo string
	autor  string
	ano    int
}

func main() {
	p := livro{titulo: "Mim dê", autor: "Lobo pidão", ano: 1639}
	fmt.Println("Título: ", p.titulo)
	fmt.Println("Autor: ", p.autor)
	fmt.Println("Ano: ", p.ano)
}
