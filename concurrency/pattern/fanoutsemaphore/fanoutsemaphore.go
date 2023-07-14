package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fanOutSemaphore()
}

// Este patrón es útil cuando quieres limitar la cantidad de gorutinas que se ejecutan al mismo tiempo.
// En este caso, aunque se lanzan 20 gorutinas, solo 5 pueden ejecutarse al mismo tiempo gracias al semáforo.
// Esto puede ser útil si las gorutinas están haciendo algo que es limitado por recursos (como hacer solicitudes a un servicio web que tiene una tasa límite),
// y quieres evitar sobrecargar esos recursos.
func fanOutSemaphore() {

	emps := 20
	ch := make(chan string, emps)

	const capacity = 5
	sem := make(chan bool, capacity)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("Employee : sent signal :", emp)
			}
			<-sem
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
