package main

import (
	"github.com/gorilla/mux" // Paquete para manejar rutas
	"log"
	"net/http"
)

// Este patrón se utiliza para encapsular la complejidad interna de la aplicación de microservicios de los clientes externos.
// En otras palabras, en lugar de tener los clientes interactuando directamente con cada microservicio,
// los clientes interactúan con la Gateway API, que a su vez interactúa con los microservicios correspondientes
func main() {
	r := mux.NewRouter()

	// Registrar rutas para cada microservicio
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Aquí se implementaría la lógica para manejar las solicitudes al microservicio de usuarios.
	})

	r.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Aquí se implementaría la lógica para manejar las solicitudes al microservicio de pedidos.
	})

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":8080", r))
}
