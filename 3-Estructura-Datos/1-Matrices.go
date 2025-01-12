package main

import "fmt"

func main() {

	//Declaracion 1
	//var arreglo [5]int

	//Declaracion e inicializacion
	var arreglito = [5]int{1, 2, 3, 4, 5}

	//Iterar con Range
	for i, value := range arreglito {
		fmt.Printf("\nIndice: %d,Valor: %d\n",i,value)
	}

    //Matrices
	var matriz= [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(matriz)
}