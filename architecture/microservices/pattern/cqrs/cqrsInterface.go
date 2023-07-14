package main

type Command interface {
	Execute()
}

type Query interface {
	Execute() interface{}
}
