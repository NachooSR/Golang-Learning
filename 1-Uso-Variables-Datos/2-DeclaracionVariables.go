package main

func main() {

	/* Declaracion Variable*/
	// var nombre string

	//Muchas variables
	//1-
	var nombre, dni string

	//2-
	var (
		nombre, apellido string
		edad             int
	)

	//3-
	var nombre, apellido, edad = "Juan", "Perez", 21

	/* Asignacion*/
	var nombre string = "Juan"

	/* Declaracion dinamica  (Unicamente dentro de funciones) */
	nombre := "Juan"

}
