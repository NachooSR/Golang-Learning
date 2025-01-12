//Arreglos Dinamicos

package main

import (
	"fmt"
	"slices"
)

func main() {

	//Declaracion
	//var slice []int

	//Declaracion e inicializacion
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

	//Slice a partir de otro
	dias2 := diasSemana[0:5]

	//Longitud
	fmt.Println(len(diasSemana))
	//Con cap vemos la capacidad

	//Agregar Elementos
	dias2 = append(dias2, "Viernes")

	//Eliminar
	dias2 = slices.Delete(dias2, 0, 2)
	fmt.Println(dias2)

}
