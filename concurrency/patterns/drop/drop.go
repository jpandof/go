package main

import (
	"fmt"
	"time"
)

func main() {
	drop()
}

// fanOutDropPattern es una implementación del patrón Drop en Go.
// Este patrón es útil cuando necesitas manejar una cantidad alta de tareas y
// tienes un límite de cuántas puedes manejar a la vez. En este caso, una vez que
// alcanzamos la capacidad del canal, las tareas adicionales simplemente se "soltarán".
// Esto puede ser útil en situaciones donde se necesita evitar el backpressure,
// como en un servidor que está siendo atacado con muchas solicitudes.
func drop() {
	const capacity = 5
	ch := make(chan string, capacity)

	go func() {
		for p := range ch {
			fmt.Print("Employee : receive shutdown singal\n", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : sent signal :", w)
		default:
			fmt.Println("manager : dropped data :", w)

		}
		ch <- "paper"
		fmt.Println("manager : sent signal", w)
	}

	close(ch)
	fmt.Println("manager sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-----------------------------------------")
}
