package main

/*
 Ejercicio: Calcule e imprima el área y el perímetro del triángulo

 Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.

El programa debe:

Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.

Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.

Calcular el área del triángulo usando la fórmula base x altura / 2.

Calcular el perímetro del triángulo sumando los lados.

Imprimir el área y el perímetro del triángulo con dos decimales de precisión.

El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro. También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.

Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.

*/

import (
	"fmt"
	"math"
)

const precision = 2

func main() {

	var ladoA, ladoB float64

	fmt.Println("Lado A: ")
	fmt.Scanln(&ladoA)
	fmt.Println("Lado B: ")
	fmt.Scanln(&ladoB)

	//Variables
	hipotenusa := calculoHipotenusa(float32(ladoA), float32(ladoB))

	area := calcularArea(float32(ladoA), float32(ladoB))

	perimetro := calcularPerimetro(float32(ladoA), float32(ladoB), float32(hipotenusa))

	//Muestreo
	fmt.Printf("\nHipotenusa: %.*f \n", precision, hipotenusa)
	fmt.Printf("Area: %.*f\n", precision, area)
	fmt.Printf("Perimetro: %.*f\n", precision, perimetro)

}

func calculoHipotenusa(ladoA, ladoB float32) float64 {

	hipotenusa := math.Hypot(float64(ladoA), float64(ladoB))

	return hipotenusa
}

func calcularArea(base, altura float32) float32 {

	area := (base * altura) / 2

	return area
}

func calcularPerimetro(ladoA, ladoB, ladoC float32) float32 {

	perimetro := ladoA + ladoB + ladoC

	return perimetro
}
