func (l *lista) valorNaPosicao(posicaoProcurada int) (int, bool) {
	if posicaoProcurada < 0 {
		return 0, false
	}

	atual := l.inicio
	indice := 0

	for atual != nil {
		if indice == posicaoProcurada {
			return atual.valor, true
		}
		atual = atual.proximo
		indice++
	}

	return 0, false
}