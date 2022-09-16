package main

import "strconv"

func media(n1 float64, n2 float64, n3 float64, n4 float64) float64 {
	quantBimestre,_ := strconv.Atoi("4");
	media := (n1 + n2 + n3 + n4) / float64(quantBimestre);
	return media;
}