package main

import "fmt"

//Declaracion de estructura
type Persona struct{
   nombre string
   edad int
   correo string
}

func main(){


	
	p:=Persona{
		nombre: "Carlos",
		edad: 21,
		correo: "carlos@gmail.com",
	}
	fmt.Println(p)

}