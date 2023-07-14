package main

import (
	"fmt"
	"time"
)

// El patrón de compartimentos estancos o "Bulkhead" es un patrón de diseño de software que se utiliza para mejorar
// la resistencia y la tolerancia a fallos en un sistema distribuido.
// Se inspira en el diseño de los compartimentos estancos de los barcos.
// Si un compartimento se inunda, los demás permanecen secos y el barco sigue flotando.
//
// Aplicado al desarrollo de software, este patrón implica aislar elementos de una aplicación en compartimentos independientes para que si uno falla,
// los demás puedan seguir funcionando. De esta manera, el patrón de compartimentos estancos puede ayudar a prevenir
// que un problema en una parte del sistema se propague a otras partes del sistema.
//
// Por ejemplo, si tienes un sistema con varios microservicios y uno de ellos se cae o empieza a funcionar de manera incorrecta,
// si se implementa correctamente el patrón de compartimentos estancos,
// los demás microservicios pueden seguir funcionando sin verse afectados por el que está fallando.

// Caso Práctico
// Un caso práctico podría ser una aplicación de comercio electrónico con varios microservicios, como servicio de catálogo,
//  servicio de pedidos y servicio de autenticación de usuarios.
//
//Supongamos que se produce un error en el servicio de autenticación debido a un pico de tráfico inusual.
// Sin el patrón de compartimentos estancos, este problema podría afectar a todo el sistema,
//  causando que los usuarios no solo no puedan iniciar sesión, sino que tampoco puedan navegar por el catálogo ni realizar pedidos.
//
//Si aplicamos el patrón de compartimentos estancos, se podrían aislar las instancias del servicio de autenticación en su propio "compartimento".
// Si se produce un fallo en este servicio, no afectará a los demás microservicios,
// y los usuarios aún podrán navegar por el catálogo y realizar pedidos como invitados mientras se resuelve el problema del servicio de autenticación.
//
//Para implementarlo en Go, podrías diseñar tu sistema para que cada microservicio se ejecute en su propia goroutine
// y se comunique con los demás a través de canales. Esto permitiría que cada servicio se ejecute de forma
// independiente y aísle los errores para evitar que se propaguen al resto del sistema

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Inicializando 3 workers, cada uno ejecutándose en su propia goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Enviando 9 trabajos
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Recibiendo los resultados de los trabajos
	for a := 1; a <= 9; a++ {
		<-results
	}
}
