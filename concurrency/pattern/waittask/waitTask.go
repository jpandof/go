package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForTask()
}

// Este patrón se utiliza cuando es necesario realizar un trabajo en segundo plano,
// pero el programa necesita el resultado de ese trabajo antes de que pueda continuar.
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("Receive signal", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "paper"

	fmt.Println("manager : sent signal")

	time.Sleep(time.Second)

	fmt.Printf("------------------------------------------")

}
