package main

import (
	"log"
	"net/http"
)

// El patrón Sidecar es un patrón de diseño de microservicios que permite desacoplar la infraestructura
// y la lógica de negocio de una aplicación, permitiendo a los desarrolladores centrarse
// en los aspectos importantes de su aplicación sin tener que preocuparse por las cuestiones de infraestructura.
//
//Por ejemplo, podrías tener un servicio de base de datos que necesita realizar copias de seguridad regulares.
// En lugar de codificar esta lógica directamente en el servicio de la base de datos (lo que puede llevar a problemas de acoplamiento),
// puedes utilizar un servicio sidecar para manejar las copias de seguridad.
// El servicio sidecar se ejecuta en el mismo entorno que el servicio principal y puede comunicarse directamente con él,
// pero es una entidad separada y no está acoplada al servicio principal.
//
// Aquí hay un ejemplo simplificado de cómo podrías estructurar esto en Go.
// En este caso, el servicio principal es un simple servidor web,
// y el sidecar es un servicio de logging que registra todas las solicitudes que llegan al servidor.

// MainService maneja las solicitudes entrantes
func MainService(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola, mundo!"))
}

// LoggingSidecar registra información sobre la solicitud
func LoggingSidecar(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recibida solicitud %s %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/", LoggingSidecar(MainService))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// En este ejemplo, el servicio principal (MainService) es un simple servidor web que responde con "Hola, mundo!".
// El sidecar (LoggingSidecar) es una función que toma un http.HandlerFunc como parámetro y devuelve otro http.HandlerFunc.
// El sidecar envuelve al manejador original, añadiendo funcionalidad adicional antes de llamar al manejador original.
// En este caso, simplemente registra la información sobre las solicitudes entrantes.
//
// Este es un ejemplo muy básico y simplificado, y en la práctica, un sidecar probablemente se ejecutaría
// en un proceso separado o incluso en un contenedor de Docker separado.
// Sin embargo, espero que te ayude a entender el concepto básico de los servicios sidecar
// y cómo pueden ser utilizados para desacoplar la funcionalidad de infraestructura de tu aplicación de la lógica principal del negocio
