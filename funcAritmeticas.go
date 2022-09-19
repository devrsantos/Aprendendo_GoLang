package main;

import (
	"fmt"
	"math"
)

func FuncBasicas() {
	a := 2
	b := 3

	fmt.Println("Soma => ", a + b)
	fmt.Println("Subtração => ", a - b)
	fmt.Println("Multiplicação => ", a * b)
	fmt.Println("Divisão => ", a / b)
	fmt.Println("Modulo => ", a & b)
}

func FuncMath() {
	c := 4.0
	d := 5.0

	fmt.Println("Maior => ", math.Max(float64(c), float64(d)))
	fmt.Println("Menor => ", math.Min(c, d))
	fmt.Println("Exponenciação => ", math.Pow(c, d))
}