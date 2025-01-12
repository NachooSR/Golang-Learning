package main

import (
	"fmt"
	"math/rand"

	
)

func main() {

	jugar()
	
}

func jugar(){
  
	numAleatorio:= rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos= 10

	for intentos < maxIntentos {
		fmt.Printf("\nIngrese un numero(%d intentos restantes): ",maxIntentos-intentos)
		fmt.Scanln(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Println("FELICITACIONES ACERTO!!!")
			JugarAgain()
			return //Si ponemos break el ultimo mensaje se ejecutara
		}else if numIngresado < numAleatorio{
			fmt.Println("El numero a adivinar es mayor")
		}else{
			fmt.Println("El numero a adivinar es menor")
		}
		
		
		intentos++
	}

	fmt.Println("Lo siento no acerto, el numero era: ",numAleatorio)
	JugarAgain()
}

func JugarAgain(){
	var eleccion string
	fmt.Println("Desea jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch (eleccion){
	case "s": jugar()
	JugarAgain()
	break
	case "n":
		fmt.Println("Gracias por jugar!")
		return
	}
}