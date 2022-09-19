package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	const Pi = 3.1415
	var raio = 3.2
	area := Pi * math.Pow(raio, 2)
	fmt.Println("A área é", area)

	FuncBasicas()
	FuncMath()

	nome := `Renan Augusto`
	fmt.Println("A variável nome é do tipo", reflect.TypeOf(nome), "e tem", len(nome), "caracteres")
}