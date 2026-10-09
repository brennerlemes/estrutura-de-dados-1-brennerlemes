package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ---- Pilha encadeada (do Exercício 2) ----

type No struct {
	valor   int
	proximo *No
}

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

func (p *PilhaEncadeada) Tamanho() int {
	return p.tamanho
}

// ---- Calculadora RPN ----

// AvaliarRPN avalia uma expressão em notação polonesa reversa.
// Tokens separados por espaço; operadores + - * / e inteiros.
func AvaliarRPN(expr string) (int, error) {
	pilha := &PilhaEncadeada{}

	for _, tok := range strings.Fields(expr) {
		switch tok {
		case "+", "-", "*", "/":
			// desempilha b primeiro, depois a; empilha a op b
			b, ok := pilha.Desempilhar()
			if !ok {
				return 0, errors.New("operandos insuficientes")
			}
			a, ok := pilha.Desempilhar()
			if !ok {
				return 0, errors.New("operandos insuficientes")
			}

			var r int
			switch tok {
			case "+":
				r = a + b
			case "-":
				r = a - b
			case "*":
				r = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("divisão por zero")
				}
				r = a / b
			}
			pilha.Empilhar(r)

		default:
			n, err := strconv.Atoi(tok)
			if err != nil {
				return 0, fmt.Errorf("token inválido: %q", tok)
			}
			pilha.Empilhar(n)
		}
	}

	if pilha.Tamanho() == 0 {
		return 0, errors.New("pilha vazia ao final")
	}
	if pilha.Tamanho() > 1 {
		return 0, errors.New("expressão termina com mais de um valor na pilha")
	}

	resultado, _ := pilha.Desempilhar()
	return resultado, nil
}

func main() {
	expressoes := []string{
		"3 4 + 2 *",         // 14
		"5 1 2 + 4 * + 3 -", // 14
		"3 +",               // erro
		"1 2",               // erro
		"4 0 /",             // erro
		"2 x +",             // erro
	}

	for _, e := range expressoes {
		r, err := AvaliarRPN(e)
		if err != nil {
			fmt.Printf("%-20s -> erro: %v\n", e, err)
		} else {
			fmt.Printf("%-20s -> %d\n", e, r)
		}
	}
}
