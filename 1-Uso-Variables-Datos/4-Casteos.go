package main

import (
	"fmt"
	"strconv"
)

func main() {

	var integer16 int16 = 50
	var integer32 int32 = 100

	// ARROJA ERROR != TIPOS fmt.Println(integer16 + integer32)

	fmt.Println(int32(integer16) + integer32) //CASTEO


	//CASTEO DE STRING A ENTERO
    s:= "100"
    numerito,_ := strconv.Atoi(s)
   
	fmt.Println(numerito)

    ///Al reves
	numerito = 40
	s= strconv.Itoa(numerito)

   //%v en fmt cuando no sabemos que tipo de dato estamos mostrando
}