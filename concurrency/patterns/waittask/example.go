package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForTask()
}

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
