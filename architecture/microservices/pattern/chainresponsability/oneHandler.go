package main

import "fmt"

// OneHandler maneja las solicitudes que contienen la palabra "one".
type OneHandler struct {
	next Handler
}

func (h *OneHandler) SetNext(next Handler) {
	h.next = next
}

func (h *OneHandler) Handle(request string) {
	if request == "one" {
		fmt.Println("OneHandler maneja la solicitud")
	} else if h.next != nil {
		h.next.Handle(request)
	}
}
