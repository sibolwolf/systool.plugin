package main

import (
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("./datahandle/datahandle.so")
	if err != nil {
		log.Println(err)
	}

	add, _ := p.Lookup("Add")
	substract, _ := p.Lookup("Substract")
	sum := add.(func(int, int) int)(3, 2)
	sub := substract.(func(int, int) int)(3, 2)
	log.Println(sum, sub)
}
