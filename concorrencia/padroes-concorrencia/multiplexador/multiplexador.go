package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	canal := multiplexar(escrever("Ola mundo"), escrever("nova mensagem"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			// vai verificar qual canal tem mensagem e jogar para canal de saida
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			numero, _ := rand.Int(rand.Reader, big.NewInt(2000))
			time.Sleep(time.Millisecond * time.Duration(numero.BitLen()))
		}
	}()

	return canal
}
