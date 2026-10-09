package main

import "fmt"

type No struct {
	valor   int
	proximo *No
}

// PilhaEncadeada é uma pilha (LIFO) sem limite de capacidade
type PilhaEncadeada struct {
	topo    *No
	tamanho int
}

func (p *PilhaEncadeada) Empilhar(v int) {
	novo := &No{valor: v, proximo: p.topo}
	p.topo = novo
	p.tamanho++
}

func (p *PilhaEncadeada) Desempilhar() (int, bool) {
	if p.topo == nil {
		return 0, false
	}
	v := p.topo.valor
	p.topo = p.topo.proximo
	p.tamanho--
	return v, true
}

func (p *PilhaEncadeada) Topo() (int, bool) {
	if p.topo == nil {
		return 0, false
	}
	return p.topo.valor, true
}

func (p *PilhaEncadeada) Vazia() bool {
	return p.topo == nil
}

func (p *PilhaEncadeada) Tamanho() int {
	return p.tamanho
}

func main() {
	p := &PilhaEncadeada{}

	// 1. Empilhar 1, 2 e 3; imprimir Tamanho e Topo (esperado: 3 e 3)
	p.Empilhar(1)
	p.Empilhar(2)
	p.Empilhar(3)
	topo, _ := p.Topo()
	fmt.Println("Tamanho:", p.Tamanho(), "| Topo:", topo)

	// 2. Desempilhar até esvaziar e conferir que Tamanho volta a 0
	fmt.Print("Desempilhados: ")
	for !p.Vazia() {
		v, _ := p.Desempilhar()
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println("Tamanho após esvaziar:", p.Tamanho())

	// 3. Mesma sequência de testes do Exercício 1 (sem overflow)
	fmt.Println("\n--- Sequência do Exercício 1 (sem overflow) ---")
	for _, v := range []int{10, 20, 30, 40, 50, 60} {
		p.Empilhar(v)
		fmt.Printf("Empilhar(%d) | tamanho: %d\n", v, p.Tamanho())
	}

	fmt.Print("Desempilhados: ")
	for !p.Vazia() {
		v, _ := p.Desempilhar()
		fmt.Print(v, " ")
	}
	fmt.Println()

	_, ok := p.Desempilhar()
	fmt.Println("Desempilhar de pilha vazia -> ok =", ok)
}
