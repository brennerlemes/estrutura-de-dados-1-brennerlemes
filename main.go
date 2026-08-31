package main

import "fmt"

type no struct {
	valor   int
	proximo *no
}

type lista struct {
	inicio *no
}

func (l *lista) adicionarInicio(valor int) {
	novo := &no{valor: valor}
	novo.proximo = l.inicio
	l.inicio = novo
}

func (l *lista) adicionarFim(valor int) {
	novo := &no{valor: valor}

	if l.inicio == nil {
		l.inicio = novo
		return
	}

	atual := l.inicio
	for atual.proximo != nil {
		atual = atual.proximo
	}
	atual.proximo = novo
}

func (l *lista) imprimir() {
	atual := l.inicio
	for atual != nil {
		fmt.Printf("%d -> ", atual.valor)
		atual = atual.proximo
	}
	fmt.Println("nil")
}

func main() {
	l := &lista{}

	l.adicionarInicio(10)
	l.adicionarFim(20)
	l.adicionarInicio(5)
	l.adicionarFim(30)
	l.adicionarInicio(1)

	l.imprimir()
}