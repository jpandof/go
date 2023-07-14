package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	settings := gobreaker.Settings{
		Interval:    time.Duration(10) * time.Second, // período de reseteo del circuit breaker
		MaxRequests: 5,                               // número de solicitudes permitidas para probar si el circuito debe resetearse
		Timeout:     time.Duration(30) * time.Second, // tiempo que el circuit breaker debe estar abierto antes de permitir las solicitudes
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
	}

	cb = gobreaker.NewCircuitBreaker(settings)
}

func UnreliableService() error {
	// Supongamos que este es un servicio que puede fallar.
	// Para simular esto, simplemente devolvemos un error.
	return errors.New("service failed")
}

func CallService() error {
	result, err := cb.Execute(func() (interface{}, error) {
		return nil, UnreliableService()
	})

	if err != nil {
		fmt.Println("Service call failed:", err)
		return err
	}

	fmt.Println("Service call successful:", result)
	return nil
}

func main() {
	for i := 0; i < 10; i++ {
		CallService()
		time.Sleep(2 * time.Second)
	}
}
