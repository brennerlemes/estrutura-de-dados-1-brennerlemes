package main

import "fmt"

// ---- Fila encadeada (do Exercício 5) ----

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

func (f *FilaEncadeada) Tamanho() int {
	return f.tamanho
}

// ---- Atendimento por senha ----

type Atendimento struct {
	proxSenha int
	fila      *FilaEncadeada
}

func NovoAtendimento() *Atendimento {
	return &Atendimento{proxSenha: 1, fila: &FilaEncadeada{}}
}

// Chegar gera a próxima senha, enfileira e imprime o estado
func (a *Atendimento) Chegar() int {
	senha := a.proxSenha
	a.proxSenha++
	a.fila.Enfileirar(senha)
	fmt.Printf("Chegou senha %d | aguardando: %d\n", senha, a.fila.Tamanho())
	return senha
}

// Chamar desenfileira a senha mais antiga e imprime o estado
func (a *Atendimento) Chamar() (int, bool) {
	senha, ok := a.fila.Desenfileirar()
	if !ok {
		fmt.Println("Nenhuma senha aguardando")
		return 0, false
	}
	fmt.Printf("Chamando senha %d | aguardando: %d\n", senha, a.fila.Tamanho())
	return senha, true
}

func main() {
	a := NovoAtendimento()

	a.Chegar()
	a.Chegar()
	a.Chamar()
	a.Chegar()
	a.Chamar()
	a.Chamar()
	a.Chamar()
}
