package main

// El patrón de Chain of Responsibility, también conocido como cadena de responsabilidad,
// es un patrón de diseño de comportamiento que permite desacoplar a los emisores de solicitudes de sus receptores,
// dando a más de un objeto la oportunidad de manejar la solicitud.
// Los objetos de manejo se encadenan y pasan la solicitud a lo largo de la cadena hasta que un objeto la maneja
// Este patrón es útil cuando tienes varias maneras de manejar una solicitud y no sabes cuál debe manejar una solicitud específica hasta el momento de la ejecución.
// También permite agregar o eliminar fácilmente manejadores sin cambiar el código que emite la solicitud
func main() {
	zeroHandler := &ZeroHandler{}
	oneHandler := &OneHandler{}

	// Configuramos la cadena de responsabilidad.
	zeroHandler.SetNext(oneHandler)

	zeroHandler.Handle("zero") // ZeroHandler maneja la solicitud.
	zeroHandler.Handle("one")  // OneHandler maneja la solicitud.
	zeroHandler.Handle("two")  // La solicitud no se maneja.
}
