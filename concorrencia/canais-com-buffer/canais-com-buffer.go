package main

import (
	"fmt"
)

func main() {
	// 1. criando canal -> usado para enviar e receber dados
	canal := make(chan string, 2) // canal com buffer so bloquea quando atinge a capacidate de 2 do canal

	// enviando dados para o canal
	canal <- "ola mundo" // bloqueia (para fluxo e da deadlock) quando esta rodando na mesma thead -> por isso e criada em treads separadas

	//recebendo a mensagem que foi enviada no canal
	msg := <-canal
	fmt.Println(msg)

	fmt.Println("Fim do programa")
}
