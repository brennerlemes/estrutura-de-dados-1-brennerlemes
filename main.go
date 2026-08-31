func (l *lista) removerPosicao(posicao int) (int, bool) {
	if posicao < 0 || l.inicio == nil {
		return 0, false
	}

	if posicao == 0 {
		valor := l.inicio.valor
		l.inicio = l.inicio.proximo
		return valor, true
	}

	anterior := l.inicio
	for i := 0; i < posicao-1; i++ {
		if anterior.proximo == nil {
			return 0, false
		}
		anterior = anterior.proximo
	}

	if anterior.proximo == nil {
		return 0, false
	}

	valor := anterior.proximo.valor
	anterior.proximo = anterior.proximo.proximo
	return valor, true
}