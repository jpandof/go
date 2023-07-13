package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fanOut()
}

// Este patrón se utiliza cuando tienes una cantidad de tareas que pueden realizarse en paralelo
// y quieres distribuir esas tareas entre varias goroutines para aprovechar
// la concurrencia y potencialmente acelerar el tiempo total de procesamiento.
func fanOut() {
	emps := 20
	ch := make(chan string, emps)
	for e := 0; e < emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("Employee : sent signal :", emp)
		}(e)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("Manager receive signal", emps)
		emps--
	}

	time.Sleep(time.Second)
	fmt.Println("-----------------------------------------")

}
