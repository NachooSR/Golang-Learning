package main

import "fmt"

func main() {

	// For unico ciclo repetitivo

	// Bucle infinito --> for { }

	//1-
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//2- Declaracion de variable en el bucle
    for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	//Break
	for i<=10 {
		if i==5 {
			break //Si se cumple la condicion el ciclo corta
		}
	}

	//Continue
	for i<=0 {
		if i==5 {
			continue //Saltea la iteracion actual
		}
	}

}