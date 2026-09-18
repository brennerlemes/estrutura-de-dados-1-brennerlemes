package main

import "fmt"

type No struct {
	valor int
	ant   *No
	prox  *No
}

type ListaDupla struct {
	head *No
	tail *No
}

func NovaLista(valores []int) *ListaDupla {
	l := &ListaDupla{}
	for _, v := range valores {
		l.inserirFimComTail(v)
	}
	return l
}

func (l *ListaDupla) inserirFimComTail(valor int) {
	novo := &No{valor: valor}
	if l.head == nil {
		l.head = novo
		l.tail = novo
		return
	}
	novo.ant = l.tail
	l.tail.prox = novo
	l.tail = novo
}

func (l *ListaDupla) InserirInicio(valor int) {
	novo := &No{valor: valor}
	if l.head == nil {
		l.head = novo
		l.tail = novo
		return
	}
	novo.prox = l.head
	l.head.ant = novo
	l.head = novo
}

func (l *ListaDupla) InserirFimSemTail(valor int) {
	novo := &No{valor: valor}
	if l.head == nil {
		l.head = novo
		l.tail = novo
		return
	}
	atual := l.head
	for atual.prox != nil {
		atual = atual.prox
	}
	atual.prox = novo
	novo.ant = atual
	l.tail = novo
}

func (l *ListaDupla) RemoverHead() {
	if l.head == nil {
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}
	l.head = l.head.prox
	l.head.ant = nil
}

func (l *ListaDupla) RemoverTail() {
	if l.tail == nil {
		return
	}
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return
	}
	l.tail = l.tail.ant
	l.tail.prox = nil
}

func (l *ListaDupla) Imprimir() {
	atual := l.head
	fmt.Print("[")
	for atual != nil {
		fmt.Print(atual.valor)
		if atual.prox != nil {
			fmt.Print(" <-> ")
		}
		atual = atual.prox
	}
	fmt.Println("]")
}

func main() {
	lista := NovaLista([]int{10, 20, 50, 60, 80})
	fmt.Println("Lista inicial:")
	lista.Imprimir()

	lista.InserirInicio(5)
	fmt.Println("\n1) Após inserir 5 no começo:")
	lista.Imprimir()

	lista.InserirFimSemTail(5)
	fmt.Println("\n2) Após inserir 5 no final (sem usar o tail):")
	lista.Imprimir()

	lista.RemoverHead()
	fmt.Println("\n3) Após remover o head:")
	lista.Imprimir()

	lista.RemoverTail()
	fmt.Println("\n4) Após remover o tail:")
	lista.Imprimir()
}