package main

import "fmt"

const TamanhoTabela = 7

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
		if t.gavetas[i] == nil {
			fmt.Printf("Gaveta %d: vazio\n", i)
			continue
		}
		fmt.Printf("Gaveta %d: ", i)
		for atual := t.gavetas[i]; atual != nil; atual = atual.Proximo {
			fmt.Printf("[%d: %s] -> ", atual.Chave, atual.Valor)
		}
		fmt.Println("nil")
	}
}

func (t *TabelaHash) Remover(chave int) bool {
	indice := t.hash(chave)
	atual := t.gavetas[indice]
	var anterior *No

	for atual != nil {
		if atual.Chave == chave {
			if anterior == nil {
				// remoção na cabeça da lista
				t.gavetas[indice] = atual.Proximo
			} else {
				anterior.Proximo = atual.Proximo
			}
			return true
		}
		anterior = atual
		atual = atual.Proximo
	}
	return false
}

func main() {
	tabela := &TabelaHash{}

	tabela.Inserir(14, "Algoritmos")
	tabela.Inserir(23, "Estruturas de Dados")
	tabela.Inserir(7, "Banco de Dados")
	tabela.Inserir(30, "Redes")
	tabela.Inserir(44, "Compiladores")
	tabela.Inserir(19, "Sistemas Operacionais")

	fmt.Println("--- Estado da Tabela Hash ---")
	tabela.Imprimir()

	fmt.Println("\n--- Testes de Busca ---")
	for _, codigo := range []int{30, 14, 99} {
		if titulo, ok := tabela.Buscar(codigo); ok {
			fmt.Printf("Chave %d encontrada: %s\n", codigo, titulo)
		} else {
			fmt.Printf("Chave %d NAO encontrada\n", codigo)
		}
	}
}
