package main

import (
	"fmt"
	"unicode/utf8"
)

const (
	Capacidade     = 1000
	MaxTamanhoNome = 100
)

type Aluno struct {
	Matricula int
	Nome      string
}

type TabelaAlunos struct {
	posicoes [Capacidade][]Aluno
}

func calcularIndice(matricula, tamanhoNome int) int {
	idx := int(float64(matricula) / float64(tamanhoNome) * (1.0 / 10.0))
	return idx % Capacidade
}

func (t *TabelaAlunos) Inserir(a Aluno) error {
	if a.Matricula < 0 || a.Matricula > 1000 {
		return fmt.Errorf("matrícula %d fora do intervalo 0..1000", a.Matricula)
	}
	tam := utf8.RuneCountInString(a.Nome)
	if tam == 0 {
		return fmt.Errorf("nome vazio")
	}
	i := calcularIndice(a.Matricula, tam)
	t.posicoes[i] = append(t.posicoes[i], a)
	return nil
}

func (t *TabelaAlunos) Buscar(matricula int) (Aluno, bool) {
	for tam := 1; tam <= MaxTamanhoNome; tam++ {
		i := calcularIndice(matricula, tam)
		for _, a := range t.posicoes[i] {
			if a.Matricula == matricula {
				return a, true
			}
		}
	}
	return Aluno{}, false
}

func main() {
	t := &TabelaAlunos{}

	alunos := []Aluno{
		{100, "Ana"},
		{250, "Carlos"},
		{999, "Maria"},
		{500, "Joao"},
		{1000, "Beatriz"},
	}
	for _, a := range alunos {
		if err := t.Inserir(a); err != nil {
			fmt.Println("Erro:", err)
		}
	}

	for _, m := range []int{250, 1000, 777} {
		if a, ok := t.Buscar(m); ok {
			fmt.Printf("Matrícula %d encontrada: %s\n", m, a.Nome)
		} else {
			fmt.Printf("Matrícula %d NÃO encontrada\n", m)
		}
	}
}
