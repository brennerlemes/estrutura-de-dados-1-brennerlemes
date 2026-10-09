package main

import "fmt"

const Capacidade = 5

// PilhaArray é uma pilha (LIFO) de capacidade fixa; vazia quando topo == -1
type PilhaArray struct {
	dados [Capacidade]int
	topo  int
}

func NovaPilhaArray() *PilhaArray {
	return &PilhaArray{topo: -1}
}

func (p *PilhaArray) Vazia() bool {
	return p.topo == -1
}

func (p *PilhaArray) Cheia() bool {
	return p.topo == Capacidade-1
}

// Empilhar devolve false em caso de overflow
func (p *PilhaArray) Empilhar(v int) bool {
	if p.Cheia() {
		return false
	}
	p.topo++
	p.dados[p.topo] = v
	return true
}

// Desempilhar devolve ok == false em caso de underflow
func (p *PilhaArray) Desempilhar() (int, bool) {
	if p.Vazia() {
		return 0, false
	}
	v := p.dados[p.topo]
	p.topo--
	return v, true
}

func (p *PilhaArray) Topo() (int, bool) {
	if p.Vazia() {
		return 0, false
	}
	return p.dados[p.topo], true
}

func main() {
	p := NovaPilhaArray()

	// 1. Empilhar 10..60 (a sexta chamada deve dar overflow)
	for _, v := range []int{10, 20, 30, 40, 50, 60} {
		ok := p.Empilhar(v)
		fmt.Printf("Empilhar(%d) -> %v\n", v, ok)
	}

	// 2. Desempilhar tudo
	fmt.Print("Desempilhados: ")
	for !p.Vazia() {
		v, _ := p.Desempilhar()
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 3. Underflow
	_, ok := p.Desempilhar()
	fmt.Println("Desempilhar de pilha vazia -> ok =", ok)
}
