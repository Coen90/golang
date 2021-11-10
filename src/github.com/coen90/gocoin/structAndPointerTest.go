package main

import (
	"fmt"

	"github.com/coen90/gocoin/person"
)

func structAndPointerTest() {
	nico := person.Person{}
	nico.SetDetails("nico", 12)
	fmt.Println("Main's 'nico'", nico)
	fmt.Println(nico.Name())
}