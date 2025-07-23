package main

import "fmt"

type Genitals interface {
	EngageInSex()
}

type Human struct {
	Legs string
	Arms string
	Face string
}

func (h *Human) Walk() {
	fmt.Println("Walking...")
}

type Dick struct{}

func (Dick) EngageInSex() {
	fmt.Println("fucking...")
}

type Pussy struct{}

func (Pussy) EngageInSex() {
	fmt.Println("getting fucked...")
}

type Male struct {
	*Human
	*Dick
}

func NewMale() *Male {
	return &Male{Human: &Human{}, Dick: &Dick{}}
}

type Female struct {
	*Human
	*Pussy
}

func NewFemale() *Female {
	return &Female{Human: &Human{}, Pussy: &Pussy{}}
}
