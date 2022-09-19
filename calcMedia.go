package main

import (
	"fmt"
	"strconv"
)

func Media(n1 float64, n2 float64, n3 float64, n4 float64) {
	quantBimestre,_ := strconv.Atoi("4");
	media := (n1 + n2 + n3 + n4) / float64(quantBimestre);

	fmt.Println(ConvertForString(media))
}

func ConvertForString(value float64) string {
	str := strconv.Itoa(int(value))
	text := "A média final é: " + str
	return text
}