package main

import "fmt"

type koreanAge struct {
	koreanName string
	globalAge  int
}

func (k koreanAge) koreanNameAndAge() {
	koreanAge := k.globalAge + 2
	fmt.Printf("I'm Korean. My name is %s and I'm %d years old", k.koreanName, koreanAge)
}

func structTest() {
	korean := koreanAge{"hyuntae", 30}
	korean.koreanNameAndAge()
}