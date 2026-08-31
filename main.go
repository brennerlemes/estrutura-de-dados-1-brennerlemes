func (l *lista) removerInicio() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	removido := l.inicio
	l.inicio = removido.proximo
	return removido.valor, true
}

func (l *lista) removerFim() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}

	if l.inicio.proximo == nil {
		valor := l.inicio.valor
		l.inicio = nil
		return valor, true
	}

	atual := l.inicio
	for atual.proximo.proximo != nil {
		atual = atual.proximo
	}

	valor := atual.proximo.valor
	atual.proximo = nil
	return valor, true
}