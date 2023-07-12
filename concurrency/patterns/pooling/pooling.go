package main

import (
	"fmt"
	"time"
)

func main() {
	pooling()
}

func pooling() {
	ch := make(chan string)

	const emps = 2

	for i := 0; i < emps; i++ {
		go func(element int) {
			for p := range ch {
				fmt.Printf("Employee %d : receive signal : %s\n", element, p)
			}
			fmt.Printf("Employee %d : receive shutdown singal\n", element)
		}(i)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
		fmt.Println("manager : sent signal", w)
	}

	close(ch)
	fmt.Println("manager sent all")

	time.Sleep(time.Second)
	fmt.Println("-----------------------------------------")

}
