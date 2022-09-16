package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	const Pi = 3.1415
	var raio = 3.2
	area := Pi * math.Pow(raio, 2)
	fmt.Println("A área é", area)

	mediaFinal := strconv.Itoa(int(media(7, 8, 7.9, 6.9)))

	fmt.Println("A média final é de: "+ mediaFinal);

	nome := `Renan Augusto`
	fmt.Println("A variável nome é do tipo", reflect.TypeOf(nome), "e tem", len(nome), "caracteres")
}