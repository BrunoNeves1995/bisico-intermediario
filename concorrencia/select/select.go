package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. criando canal -> usado para enviar e receber dados
	cana1, canal2 := make(chan string), make(chan string) // canal com buffer so bloquea quando atinge a capacidate de 2 do canal

	// Criando uma nova Thead de execucao
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			// 2. enviando valor para o canal
			cana1 <- "Canal 1"
		}
	}()
	// Criando uma nova Thead de execucao
	go func() {
		for {
			time.Sleep(time.Second * 2)
			// 2. enviando valor para o canal
			canal2 <- "Canal 2"
		}
	}()

	/* Criando uma nova thead para receber as mensagem dos canais
	   ->  Nao e eficiente porque a minha mesgCanal1 espera mesgCanal2 que leva 2 segundos para execultar, com isso a mesgCanal1 perde 3 fluxo de execucao por conta disso
	       sendo mesgCanal1 roda a cada 500 Milesegundos.
		for {
			mesgCanal1 := <-cana1
			fmt.Println(mesgCanal1)

			mesgCanal2 := <-canal2
			fmt.Println(mesgCanal2)
		}
	*/

	for {
		// Nao espera o canal 2 para execultar, sendo que canal 2 demora mais que canal 1
		select {
		case mesgCanal1 := <-cana1:
			fmt.Println(mesgCanal1)
		case mesgCanal2 := <-canal2:
			fmt.Println(mesgCanal2)
		}
	}
}
