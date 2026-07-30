package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, World!")

	var a string
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(a)
}
