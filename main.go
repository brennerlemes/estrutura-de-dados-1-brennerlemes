func (l *lista) tamanho() int {
	contador := 0
	atual := l.inicio
	for atual != nil {
		contador++
		atual = atual.proximo
	}
	return contador
}

func main() {
	l := &lista{}
	var opcao int

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Adicionar no início")
		fmt.Println("2. Adicionar no fim")
		fmt.Println("3. Adicionar em posição")
		fmt.Println("4. Remover do início")
		fmt.Println("5. Remover do fim")
		fmt.Println("6. Remover de posição")
		fmt.Println("7. Buscar posição de um valor")
		fmt.Println("8. Buscar valor em uma posição")
		fmt.Println("9. Ver tamanho")
		fmt.Println("10. Imprimir lista")
		fmt.Println("0. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			var valor int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			l.adicionarInicio(valor)

		case 2:
			var valor int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			l.adicionarFim(valor)

		case 3:
			var valor, posicao int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			ok := l.adicionarPosicao(valor, posicao)
			fmt.Println("Sucesso:", ok)

		case 4:
			valor, ok := l.removerInicio()
			fmt.Println("Removido:", valor, "| Sucesso:", ok)

		case 5:
			valor, ok := l.removerFim()
			fmt.Println("Removido:", valor, "| Sucesso:", ok)

		case 6:
			var posicao int
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			valor, ok := l.removerPosicao(posicao)
			fmt.Println("Removido:", valor, "| Sucesso:", ok)

		case 7:
			var valor int
			fmt.Print("Valor a buscar: ")
			fmt.Scan(&valor)
			pos, ok := l.posicao(valor)
			fmt.Println("Posição:", pos, "| Encontrado:", ok)

		case 8:
			var posicao int
			fmt.Print("Posição: ")
			fmt.Scan(&posicao)
			valor, ok := l.valorNaPosicao(posicao)
			fmt.Println("Valor:", valor, "| Encontrado:", ok)

		case 9:
			fmt.Println("Tamanho:", l.tamanho())

		case 10:
			l.imprimir()

		case 0:
			fmt.Println("Encerrando...")
			return

		default:
			fmt.Println("Opção inválida.")
		}
	}
}