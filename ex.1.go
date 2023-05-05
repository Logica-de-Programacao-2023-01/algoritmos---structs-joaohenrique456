package main

import "fmt"

type Retangulo struct {
	largura int
	altura  int
}

func area(r Retangulo) int {
	return r.largura * r.altura
}

func main() {
	var ret Retangulo
	fmt.Print("Informe a altura: ")
	fmt.Scan(&ret.altura)
	fmt.Print("Informe a largura: ")
	fmt.Scan(&ret.largura)
	a := area(ret)
	fmt.Println("Área: ", a)
}
