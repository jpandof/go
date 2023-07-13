package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForFinished()
}

// Este patrón es comúnmente utilizado en situaciones donde es necesario esperar la finalización de una o más tareas
// antes de que el programa pueda proceder, garantizando que todas las tareas previas han finalizado su ejecución.
func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("Employee : sent signal")
	}()

	_, wd := <-ch

	fmt.Println("manager : receive signal :", wd)
	time.Sleep(time.Second)
	fmt.Println("-----------------------------------------")

}
