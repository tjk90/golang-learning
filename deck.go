package main

import "fmt"

type deck []string

func (d deck) printThis() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
