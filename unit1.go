package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name        string
	Age         int
	Address     string
	Email       string
	PhoneNumber string
}

func (h Human) getInfo() {
	fmt.Printf("Я %s и мне %d лет. Мой адрес: %s\n", h.Name, h.Age, h.Address)
}
func (h Human) isAdult() bool {
	return h.Age >= 18
}
func (h *Human) UpdateEmail(newEmail string) {
	h.Email = newEmail
}
func (h *Human) UpdatePhoneNumber(newPhoneNumber string) {
	h.PhoneNumber = newPhoneNumber
}

// Action Встраивание структуры Human. В таком случае, мы имеем свободный доступ к методам и полям структуры Human
type Action struct {
	Human
	ActionType string
}
