package main

import "fmt"

const Capacidade = 5

// FilaArray é uma fila (FIFO) em array circular de capacidade fixa
type FilaArray struct {
	dados  [Capacidade]int
	cabeca int // posição do próximo a sair
	cauda  int // posição onde o próximo será inserido
	tam    int
}

func (f *FilaArray) Vazia() bool {
	return f.tam == 0
}

func (f *FilaArray) Cheia() bool {
	return f.tam == Capacidade
}

// Enfileirar devolve false em caso de overflow
func (f *FilaArray) Enfileirar(v int) bool {
	if f.Cheia() {
		return false
	}
	f.dados[f.cauda] = v
	f.cauda = (f.cauda + 1) % Capacidade
	f.tam++
	return true
}

// Desenfileirar devolve ok == false em caso de underflow
func (f *FilaArray) Desenfileirar() (int, bool) {
	if f.Vazia() {
		return 0, false
	}
	v := f.dados[f.cabeca]
	f.cabeca = (f.cabeca + 1) % Capacidade
	f.tam--
	return v, true
}

func (f *FilaArray) Frente() (int, bool) {
	if f.Vazia() {
		return 0, false
	}
	return f.dados[f.cabeca], true
}

func main() {
	f := &FilaArray{}

	// 1. Enfileirar 10..50; 60 deve dar overflow
	for _, v := range []int{10, 20, 30, 40, 50, 60} {
		ok := f.Enfileirar(v)
		fmt.Printf("Enfileirar(%d) -> %v\n", v, ok)
	}

	// 2. Desenfileirar duas vezes (saem 10 e 20)
	for i := 0; i < 2; i++ {
		v, _ := f.Desenfileirar()
		fmt.Println("Saiu:", v)
	}

	// 3. Enfileirar 60 e 70 (ocupam as posições 0 e 1 do array: wrap-around)
	fmt.Println("Enfileirar(60) ->", f.Enfileirar(60))
	fmt.Println("Enfileirar(70) ->", f.Enfileirar(70))

	// 4. Desenfileirar tudo (esperado: 30 40 50 60 70)
	fmt.Print("Desenfileirados: ")
	for !f.Vazia() {
		v, _ := f.Desenfileirar()
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 5. Underflow
	_, ok := f.Desenfileirar()
	fmt.Println("Desenfileirar de fila vazia -> ok =", ok)
}
