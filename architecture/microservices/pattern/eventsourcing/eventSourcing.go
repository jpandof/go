package main

import (
	"fmt"
)

// En el patrón de Event Sourcing, en lugar de almacenar el estado actual de un objeto o servicio,
// almacenamos una secuencia de eventos que han ocurrido en el tiempo. Esto puede tener varias ventajas,
// como permitir un seguimiento de auditoría completo, la capacidad de retroceder a un estado anterior del sistema,
// y a veces puede simplificar la lógica de negocio.

func main() {
	account := NewAccount(100.0)
	account.Deposit(50.0)
	account.Withdraw(20.0)

	fmt.Println("Balance: ", account.Balance)
	for _, event := range account.Events {
		fmt.Printf("Evento: %s, Data: %v\n", event.Tipo, event.Data)
	}
}
