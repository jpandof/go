package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForResult()
}

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
