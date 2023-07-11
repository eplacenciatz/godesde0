package ejercicios

import (
	"strconv"
)

func EsMayorACien(numero string) (int, string) {
	// numeroConvert, _ := strconv.Atoi(numero)

	numeroConvert, err := strconv.Atoi(numero)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}

	var msg string

	if numeroConvert > 100 {
		msg = "Es mayor a 100"
	} else {
		msg = "Es menor a 100"
	}

	return numeroConvert, msg
}
