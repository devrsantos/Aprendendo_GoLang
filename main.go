package main

import (
	//C:\Program Files\Go\src\devrsantos\pacote_local
	"devrsantos/pacote_local"
	
	"github.com/devrsantos/golang_calcula_media"

	"fmt"
)

func main() {
	fmt.Println("Pacote Local: O Conceito final é de:", pacoteCalculaMedia.CalculaNota(5.0, 6.0, 3.0, 5.8))
	fmt.Println("Pacote Remote: O Conceito final é de:", media.CalculaNota(5.0, 6.0, 3.0, 5.8))
}

