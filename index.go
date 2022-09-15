package main

import (
	"fmt"
	moduloMath "math" // moduloMath é um Alias, opcional
)

func main() {
	const Pi = 3.1415
	var raio = 3.2
	area := Pi * moduloMath.Pow(raio, 2)
	fmt.Println("A área é", area)

	var n1, n2, n3, n4 = 7, 8, 9, 10
	media := (n1 + n2 + n3 + n4) / 4
	fmt.Println("A média Final é", media)
}