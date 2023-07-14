package main

import "fmt"

// ZeroHandler maneja las solicitudes que contienen la palabra "zero".
type ZeroHandler struct {
	next Handler
}

func (h *ZeroHandler) SetNext(next Handler) {
	h.next = next
}

func (h *ZeroHandler) Handle(request string) {
	if request == "zero" {
		fmt.Println("ZeroHandler maneja la solicitud")
	} else if h.next != nil {
		h.next.Handle(request)
	}
}
