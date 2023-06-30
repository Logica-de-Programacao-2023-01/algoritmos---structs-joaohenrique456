package main

import (
	"fmt"
	"time"
)

// q.4

type Musica struct {
	Titulo  string
	Artista string
	Duracao time.Duration
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func imprimirPlaylist(p Playlist) {
	fmt.Printf("Nome da Playlist: %s\n", p.Nome)

	var duracaoTotal time.Duration

	for _, musica := range p.Musicas {
		fmt.Printf("Título: %s, Artista: %s, Duração: %s\n", musica.Titulo, musica.Artista, musica.Duracao)
		duracaoTotal += musica.Duracao
	}

	fmt.Printf("Tempo total da Playlist: %s\n", duracaoTotal)
}

func main() {
	musicas := []Musica{
		Musica{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3 * time.Minute},
		Musica{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4 * time.Minute},
		Musica{Titulo: "Música 3", Artista: "Artista 3", Duracao: 5 * time.Minute},
	}

	playlist := Playlist{Nome: "Minha Playlist", Musicas: musicas}
	imprimirPlaylist(playlist)
}

// q.5

type MUS struct {
	Titulo  string
	Artista string
	Duracao string
}

type Playlist struct {
	Nome    string
	Musicas []MUS
}

func buscarPlaylistsPorTitulo(tituloMusica string, playlists []Playlist) []Playlist {
	resultado := []Playlist{}

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == tituloMusica {
				resultado = append(resultado, playlist)
				break
			}
		}
	}

	return resultado
}

func main() {
	musicas1 := []MUS{
		MUS{Titulo: "Música 1", Artista: "Artista 1", Duracao: "3:30"},
		MUS{Titulo: "Música 2", Artista: "Artista 2", Duracao: "4:15"},
		MUS{Titulo: "Música 3", Artista: "Artista 3", Duracao: "5:10"},
	}

	musicas2 := []MUS{
		MUS{Titulo: "Música 4", Artista: "Artista 4", Duracao: "2:45"},
		MUS{Titulo: "Música 5", Artista: "Artista 5", Duracao: "3:20"},
	}

	playlist1 := Playlist{Nome: "Playlist 1", Musicas: musicas1}
	playlist2 := Playlist{Nome: "Playlist 2", Musicas: musicas2}

	playlists := []Playlist{playlist1, playlist2}

	titulo := "Música 3"
	playlistsEncontradas := buscarPlaylistsPorTitulo(titulo, playlists)

	if len(playlistsEncontradas) > 0 {
		fmt.Printf("Playlists encontradas com a música '%s':\n", titulo)
		for _, playlist := range playlistsEncontradas {
			fmt.Println(playlist.Nome)
		}
	} else {
		fmt.Printf("Nenhuma playlist encontrada com a música '%s'\n", titulo)
	}
}

// q.6

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) AumentarSalario(porcentagem float64) {
	f.Salario *= (1 + porcentagem/100)
}

func (f *Funcionario) DiminuirSalario(porcentagem float64) {
	f.Salario *= (1 - porcentagem/100)
}

func (f *Funcionario) TempoServico() int {
	idadeInicioTrabalho := 18
	idadeAtual := time.Now().Year() - f.Idade
	tempoServico := idadeAtual - idadeInicioTrabalho
	return tempoServico
}

func main() {
	funcionario := Funcionario{
		Nome:    "Jonas",
		Salario: 3000,
		Idade:   25,
	}

	fmt.Println("Salário atual:", funcionario.Salario)
	funcionario.AumentarSalario(10)
	fmt.Println("Salário após aumento:", funcionario.Salario)
	funcionario.DiminuirSalario(5)
	fmt.Println("Salário após redução:", funcionario.Salario)

	tempoServico := funcionario.TempoServico()
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}

// q.7

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func (a *Animal) AlterarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Espécie: %s\n", a.Especie)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Som: %s\n", a.Som)
}

func main() {
	animal := Animal{
		Nome:    "Cachorro",
		Especie: "vira-lata",
		Idade:   5,
		Som:     "Latido",
	}

	animal.ImprimirInformacoes()

	animal.AlterarSom("Au au")

	animal.ImprimirInformacoes()
}

// q.8

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	var viagemMaisCara Viagem

	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		Viagem{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: "10/06/2023", Preco: 250.0},
		Viagem{Origem: "Rio de Janeiro", Destino: "Salvador", Data: "15/07/2023", Preco: 400.0},
		Viagem{Origem: "Salvador", Destino: "Fortaleza", Data: "20/08/2023", Preco: 300.0},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)

	fmt.Printf("Viagem mais cara:\nOrigem: %s\nDestino: %s\nData: %s\nPreço: R$ %.2f\n",
		viagemMaisCara.Origem, viagemMaisCara.Destino, viagemMaisCara.Data, viagemMaisCara.Preco)
}

// q.9

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(indice int) {
	if indice >= 0 && indice < len(a.Notas) {
		a.Notas = append(a.Notas[:indice], a.Notas[indice+1:]...)
	}
}

func (a *Aluno) CalcularMedia() float64 {
	total := 0.0

	for _, nota := range a.Notas {
		total += nota
	}

	media := total / float64(len(a.Notas))
	return media
}

func (a *Aluno) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Média: %.2f\n", a.CalcularMedia())
}

func main() {
	aluno := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{8.5, 7.9, 9.1},
	}

	aluno.ImprimirInformacoes()

	aluno.AdicionarNota(8.0)
	aluno.RemoverNota(1)

	aluno.ImprimirInformacoes()
}

// q.10

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) AdicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) RemoverAvaliacao(indice int) {
	if indice >= 0 && indice < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:indice], f.Avaliacoes[indice+1:]...)
	}
}

func (f *Filme) CalcularMediaAvaliacoes() float64 {
	total := 0

	for _, avaliacao := range f.Avaliacoes {
		total += avaliacao
	}

	media := float64(total) / float64(len(f.Avaliacoes))
	return media
}

func (f *Filme) ImprimirInformacoes() {
	fmt.Printf("Título: %s\n", f.Titulo)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Média de Avaliações: %.2f\n", f.CalcularMediaAvaliacoes())
}

func main() {
	filme := Filme{
		Titulo:     "Pulp Fiction",
		Diretor:    "Quentin Tarantino",
		Ano:        1994,
		Avaliacoes: []int{5, 4, 4, 5, 3},
	}

	filme.ImprimirInformacoes()

	filme.AdicionarAvaliacao(4)
	filme.RemoverAvaliacao(2)

	filme.ImprimirInformacoes()
}
