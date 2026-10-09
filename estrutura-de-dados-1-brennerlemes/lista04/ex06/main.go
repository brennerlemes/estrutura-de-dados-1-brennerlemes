package main

import "fmt"

type No struct {
	valor   int
	proximo *No
}

type FilaEncadeada struct {
	cabeca  *No
	cauda   *No
	tamanho int
}

func (f *FilaEncadeada) Enfileirar(v int) {
	novo := &No{valor: v}
	if f.cauda == nil { 
		f.cabeca = novo
		f.cauda = novo
	} else {
		f.cauda.proximo = novo
		f.cauda = novo
	}
	f.tamanho++
}

func (f *FilaEncadeada) Desenfileirar() (int, bool) {
	if f.cabeca == nil {
		return 0, false
	}
	v := f.cabeca.valor
	f.cabeca = f.cabeca.proximo
	if f.cabeca == nil { 
		f.cauda = nil
	}
	f.tamanho--
	return v, true
}

func (f *FilaEncadeada) Frente() (int, bool) {
	if f.cabeca == nil {
		return 0, false
	}
	return f.cabeca.valor, true
}

func (f *FilaEncadeada) Vazia() bool {
	return f.cabeca == nil
}

func (f *FilaEncadeada) Tamanho() int {
	return f.tamanho
}

func main() {
	f := &FilaEncadeada{}

	f.Enfileirar(1)
	f.Enfileirar(2)
	v1, _ := f.Desenfileirar()
	v2, _ := f.Desenfileirar()
	fmt.Println("Saíram:", v1, v2, "| fila vazia:", f.Vazia())

	f.Enfileirar(3)
	frente, _ := f.Frente()
	fmt.Println("Frente:", frente, "| Tamanho:", f.Tamanho())
}
