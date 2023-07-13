package main

import (
	"github.com/eplacenciatz/godesde0/funciones"
)

func main() {
	// estado, texto := variables.ConviertoaTexto(123)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// CONDICIONALES
	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("Esto no es Windows, es", os)
	// } else {
	// 	fmt.Println("Esto es", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// EJERCICIO 01
	// numero, texto := ejercicios.EsMayorACien("10")
	// fmt.Println(numero)
	// fmt.Println(texto)

	// Captura de datos por pantalla
	// teclado.IngresoNumeros()

	// Iteraciones
	// iteraciones.Iterar()

	// EJERCICIO 02 Tabla multiplicar
	// fmt.Println(ejercicios.TablaMultiplicar())

	// Manejo de archivos
	// files.GrabaTabla()

	// files.SumaTabla()

	// files.LeoArchivo()

	// files.LeoArchivo2()

	// Funciones Anónimas y Closures
	// funciones.Calculos()
	funciones.LlamarClosure()
}
