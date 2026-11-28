package main

import "fmt"

func main() {
	// 1. Criando canal com buffer 100 posicoes
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// 3. processando as funcoes - criando 45 theads para execultar a funcao
	for i := 0; i < 45; i++ {
		go worker(tarefas, resultados)
	}

	// 4. passando os dados para os canais
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	// 5. lendo os dados de dentro do canal
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

/*
 2. criando a funcao fora do main. Que vai receber os canais
    tarefas <-chan int: especifica que esse canal apenas recebe dados
    resultados chan<- int: especifica que esse canal apenas envia dados
*/
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
