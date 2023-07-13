package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 69}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 20
	tabla[2] = 10

	tabla2 := [10]string{"Pepe", "Lucas"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
