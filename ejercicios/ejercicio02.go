package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TablaMultiplicar() string {
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
					texto += fmt.Sprintf("%d x %d = %d \n", j, numero, j*numero)
				}
				texto += fmt.Sprintln()
				break
			}
		}
	}
	return texto

}
