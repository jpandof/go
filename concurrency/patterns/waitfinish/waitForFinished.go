package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForFinished()
}

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
