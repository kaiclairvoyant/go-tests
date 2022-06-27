package main

import "fmt"

func main() {
	p, err := NewPerson("Kai", "Clairvoyant", 28)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Made by %v %v, a %v years old engineer learning Go testing\n", p.name, p.lastName, p.age)
	fmt.Println("Type \"go test\" to try it out")
}
