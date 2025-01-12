package main

import "fmt"

func main() {

	
	var x int = 10
	var p *int = &x

    fmt.Println(&x)//Direccion de memoria de la variable
	fmt.Println(p) //Aloja la dir. de memoria ya que es un puntero
	fmt.Println(*p) //Valor que almacena dicha dir. de memoria
    
	fmt.Println()
	//Cambio de valor mediante funcion y punteros
	var y = 20
	fmt.Println(y)
	editar(&y)
    fmt.Println(y)

}

//Trabajamos por referencia, no copia
func editar(x *int){

	*x=55
}