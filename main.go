package main

import (
	"devrsantos/media"

	"fmt"
)

func main() {
	fmt.Println(media.Oi())
	fmt.Println("O Conceito final é de:", media.CalculaNota(5.0, 6.0, 3.0, 5.8))
}

