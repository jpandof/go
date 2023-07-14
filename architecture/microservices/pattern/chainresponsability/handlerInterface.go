package main

// Handler es la interfaz que manejará las solicitudes.
type Handler interface {
	SetNext(Handler)
	Handle(string)
}
