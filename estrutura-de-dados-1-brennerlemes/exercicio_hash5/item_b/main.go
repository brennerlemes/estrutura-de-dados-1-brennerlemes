package main

import (
	"fmt"
	"unicode/utf8"
)

const Capacidade = 10

type TabelaNomes struct {
	posicoes [Capacidade][]string
}

func gerarIndice(nome string) int {
	return utf8.RuneCountInString(nome) % Capacidade
}

func (t *TabelaNomes) inserir(nome string) {
	i := gerarIndice(nome)
	t.posicoes[i] = append(t.posicoes[i], nome)
	fmt.Printf("\nInserido %q na posição %d\n", nome, i)
	t.mostrar()
}

func (t *TabelaNomes) mostrar() {
	for i := 0; i < Capacidade; i++ {
		if len(t.posicoes[i]) == 0 {
			fmt.Printf("[%d] vazio\n", i)
		} else {
			fmt.Printf("[%d] %v\n", i, t.posicoes[i])
		}
	}
}

func main() {
	t := &TabelaNomes{}
	nomes := []string{"Ana", "Carlos", "Maria", "Joao", "Beatriz", "Pedro", "Lucas"}
	for _, n := range nomes {
		t.inserir(n)
	}
}
