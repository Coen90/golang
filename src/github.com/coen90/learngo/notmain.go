package main

import (
	"fmt"

	"github.com/coen90/learngo/myDict"
)

func notmain() {
	dictionary := myDict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)

	}
}
