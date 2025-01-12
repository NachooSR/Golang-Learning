package main

import "fmt"

func main() {

	hello("Nacho")
	
    saludo:=hello2("Nachito")
	fmt.Println(saludo)
}

// nombre  //Parametros
func hello(name string) {

	//Entre llaves el bloque de codigo
	fmt.Println("Hola",name,":)")
}

//Funcion con retorno
func hello2(name string) string{
  return "Hola, " + name
}