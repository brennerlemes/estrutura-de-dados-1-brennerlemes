func (l *lista) posicao(valorProcurado int) (int, bool) {
	atual := l.inicio
	indice := 0

	for atual != nil {
		if atual.valor == valorProcurado {
			return indice, true
		}
		atual = atual.proximo
		indice++
	}

	return 0, false
}