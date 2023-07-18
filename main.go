package main

import (
	"github.com/eplacenciatz/godesde0/middleware"
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
	// funciones.LlamarClosure()

	//Recursión
	// funciones.Exponencia(2)

	// Array & Slices
	// arreglos_slices.MuestroArreglos()
	// arreglos_slices.MuestraSlice()
	// arreglos_slices.Capacidad()

	// Mapas
	// mapas.MostrarMapas()

	// Estructuras POO
	// users.AltaUsuario()

	// Interfaces POO
	// Pedro := new(models.Hombre)
	// e.HumanosRespirando(Pedro)

	// Maria := new(models.Mujer)
	// e.HumanosRespirando(Maria)

	// Defer, Panic & Recover
	// defer_panic.VemosDefer()
	// defer_panic.EjemploPanic()

	// GORoutines (Async)
	// go goroutines.MiNombreLentooo("Pepe LePu")

	// fmt.Println("Estoy aqui")
	// var x string
	// fmt.Scanln(&x)

	// Channels (Async)
	// canal1 := make(chan bool)
	// go goroutines.MiNombreLentooo("Pepe LePu", canal1)
	// defer func() {
	// 	<-canal1
	// }()
	// fmt.Println("Estoy aqui")

	// <-canal1

	// WebServer
	// webserver.MiWebServer()

	// Middleware
	middleware.MiMiddleware()

}
