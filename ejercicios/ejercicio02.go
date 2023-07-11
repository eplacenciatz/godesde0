package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese número")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El número ingresado es incorrecto " + err.Error())
				continue
			} else {
				for j := 1; j <= 10; j++ {
					fmt.Printf("%d x %d = %d \n", j, numero, j*numero)
				}
				break
			}
		}
	}

}
