package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Type   string
	Result string
}

// El patrón de Saga es un diseño de microservicio que proporciona una forma de resolver problemas de transacciones distribuidas.
// Este patrón sugiere dividir una transacción grande y duradera en una serie de transacciones locales más pequeñas.
// Cada una de estas transacciones locales actualiza los datos dentro de un solo servicio y luego publica un evento.
// Otros servicios escuchan estos eventos y realizan sus propias transacciones locales en respuesta.
// Si en algún momento una transacción falla, se ejecutan transacciones compensatorias para deshacer los pasos anteriores y mantener la coherencia de los datos.
func main() {
	var wg sync.WaitGroup

	events := make(chan Event)

	// Iniciar la transacción
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Iniciando transacción...")
		events <- Event{Type: "inicio", Result: "éxito"}

		// Cierra el canal aquí, después de enviar todos los eventos.
		close(events)
	}()

	// Procesar eventos
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range events {
			fmt.Println("Procesando evento:", event.Type)
			time.Sleep(500 * time.Millisecond) // Simulando un procesamiento
			fmt.Println("Evento procesado:", event.Type, "- Resultado:", event.Result)
		}
	}()

	wg.Wait()

	fmt.Println("Transacción completada")
}
