package main

import "time"

// Evento es la representación básica de un evento.
type Evento struct {
	Tipo      string
	Timestamp time.Time
	Data      map[string]interface{}
}

// La estructura de Account representa una cuenta bancaria.
type Account struct {
	Balance float64
	Events  []Evento
}

// NewAccount crea una nueva cuenta con un balance inicial.
func NewAccount(initialBalance float64) *Account {
	account := &Account{
		Balance: initialBalance,
		Events:  make([]Evento, 0),
	}
	account.addEvent("Creada", map[string]interface{}{"balance": initialBalance})
	return account
}

// addEvent registra un evento en la cuenta.
func (a *Account) addEvent(tipo string, data map[string]interface{}) {
	event := Evento{
		Tipo:      tipo,
		Timestamp: time.Now(),
		Data:      data,
	}
	a.Events = append(a.Events, event)
}

// Deposit agrega dinero a la cuenta y registra un evento.
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	a.addEvent("Deposito", map[string]interface{}{"monto": amount})
}

// Withdraw retira dinero de la cuenta y registra un evento.
func (a *Account) Withdraw(amount float64) {
	a.Balance -= amount
	a.addEvent("Retiro", map[string]interface{}{"monto": amount})
}
