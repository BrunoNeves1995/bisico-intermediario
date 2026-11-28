package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. criando canal -> usado para enviar e receber dados
	canal := make(chan string)

	// 2. usando canal
	go escrever("Ola mundo", canal) // sem go da erro de deadlok

	/*
		// 4. pegando os dados enviado dentro do canal
		for {
			// 6. verificamos se o canal esta fechado e saimos do loop infinito for
			mensagem, aberto := <-canal
			if !aberto {
				break
			}
			fmt.Println(mensagem)
		}
	*/

	// 4. pegando os dados enviado dentro do canal
	for msg := range canal {
		fmt.Println(msg)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i <= 5; i++ {
		// 3. enviado dados para dentro do canal
		canal <- texto
		time.Sleep(time.Second)
	}

	// 5. fechando o canal para evitar deadlock
	close(canal)
}
