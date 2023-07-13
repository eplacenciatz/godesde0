package arreglos_slices

import "fmt"

var tablaS []int = []int{1, 2, 3}
var arreglo [10]int = [10]int{2, 4, 5, 7, 8, 9, 11, 14, 17, 20}

func MuestraSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   // Slice desde la 3era posicion
	porcion2 := arreglo[:5]  // Slice desde 0 hasta 5ta posicion
	porcion3 := arreglo[6:7] // Slice desde 6ta hasta 7ma posicion

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
