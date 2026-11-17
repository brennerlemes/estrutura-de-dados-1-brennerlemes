package main

import "fmt"

const TamanhoTabela = 5

type No struct {
	Chave   int
	Valor   string
	Proximo *No
}

type TabelaHash struct {
	gavetas [TamanhoTabela]*No
}

func (t *TabelaHash) hash(chave int) int {
	return chave % TamanhoTabela
}

func (t *TabelaHash) Inserir(chave int, valor string) {
	indice := t.hash(chave)
	novo := &No{Chave: chave, Valor: valor, Proximo: t.gavetas[indice]}
	t.gavetas[indice] = novo
}

func (t *TabelaHash) Buscar(chave int) (string, bool) {
	atual := t.gavetas[t.hash(chave)]
	for atual != nil {
		if atual.Chave == chave {
			return atual.Valor, true
		}
		atual = atual.Proximo
	}
	return "", false
}

func (t *TabelaHash) Imprimir() {
	for i := 0; i < TamanhoTabela; i++ {
		fmt.Printf("Gaveta %d: ", i)
		if t.gavetas[i] == nil {
			fmt.Println("vazio")
			continue
		}
		for atual := t.gavetas[i]; atual != nil; atual = atual.Proximo {
			fmt.Printf("[%d: %s] -> ", atual.Chave, atual.Valor)
		}
		fmt.Println("nil")
	}
}

func main() {
	tabela := &TabelaHash{}

	// Insercoes de teste
	tabela.Inserir(10, "Borracha")
	tabela.Inserir(12, "Lapis")
	tabela.Inserir(25, "Caneta")
	tabela.Inserir(37, "Caderno")
	tabela.Inserir(42, "Mochila")

	fmt.Println("--- Estado da Tabela Hash ---")
	tabela.Imprimir()

	fmt.Println("\n--- Testes de Busca ---")
	chaves := []int{37, 10, 99}
	for _, c := range chaves {
		if val, ok := tabela.Buscar(c); ok {
			fmt.Printf("Chave %d encontrada: %s\n", c, val)
		} else {
			fmt.Printf("Chave %d NAO encontrada\n", c)
		}
	}
}
