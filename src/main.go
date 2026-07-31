package main

import (
	"fmt"
	"log"
	"mccomack/goback/src/idkwhatthisiscalledmaybemodule"
)

func main() {
	fmt.Println("Hello, World!")

	var a string
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(a)

	idkwhatthisiscalledmaybemodule.Goback()
}
