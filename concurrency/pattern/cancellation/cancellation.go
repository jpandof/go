package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cancellation()
}

// Se utiliza para manejar una situación donde una operación puede tardar demasiado tiempo
// y queremos poder cancelarla si no se completa en un plazo de tiempo determinado
func cancellation() {

	// el buffer tiene que ser 1 para que la señal que se envia desde la gorutina no se quede bloqueada ya que nadie la va a recibir despues del timeout
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : sent signal")
	}()

	tc := time.After(100 * time.Millisecond)

	select {
	case p := <-ch:
		fmt.Println("Manager receive signal :", p)
	case t := <-tc:
		fmt.Println("Manager : timeout :", t)
	}

	time.Sleep(time.Second)
	fmt.Println("-----------------------------------------")

}
