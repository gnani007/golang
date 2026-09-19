package main

import "fmt"

type name string

func (n name) log() {
	fmt.Println(n)
}

func main() {
	var name name = "Gnan"
	name.log()
}