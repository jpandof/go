package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForResult()
}

// Este patrón se utiliza cuando necesitas ejecutar una tarea en segundo plano y esperar a que se complete antes de continuar.
// Mediante el uso de goroutines y canales, este patrón permite que la tarea se ejecute de manera concurrente con el resto del programa,
// pero garantiza que el resultado estará disponible cuando sea necesario.
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("Employee : sent signal")
	}()

	p := <-ch
	fmt.Println("manager : receive signal :", p)

	time.Sleep(time.Second)

	fmt.Println("-----------------------------------------")
}
