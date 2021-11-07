package main

import (
	"fmt"
	"time"
)

func channel() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isItSexy(person, c)
	}
	fmt.Println("Waiting for messages")
	resultOne := <-c
	resultTwo := <-c
	fmt.Println("Received this message:", resultOne)
	fmt.Println("Received this message:", resultTwo)
}

func isItSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
