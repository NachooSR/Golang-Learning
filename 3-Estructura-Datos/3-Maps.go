package main

import (
	"fmt"

	
)

func main() {

	//Conjunto de datos clave-valor
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
    
	//Agregar valor
	colors["negro"]="#00000"
	
    //Verificar existencia clave
    valor,ok := colors["negro"]

	//Eliminar elemento
	delete(colors,"negro")
  
	//Iterar sobre mapa
	for clave,value:=range colors{
		fmt.Printf("\nClave:%s ,Valor:%s",clave,value)
	}

	fmt.Println(colors)
    fmt.Println(valor,ok)

}