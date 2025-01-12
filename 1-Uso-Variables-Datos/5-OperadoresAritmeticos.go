package main

import (
    "fmt"
    "math"
)

func main(){

  //Funciones del paquete math

  //Constante Pi y demas
  fmt.Println(math.Pi)

  //Potencia
  cubo:= math.Pow(2,3) //num eleveado a ...
  fmt.Println(cubo)

  //Raiz cuadrada
  fmt.Println(math.Sqrt(16))

  //Raiz cubica
  fmt.Println(math.Cbrt(27))

  //Max y Min
  fmt.Println(math.Max(10, 20))
  fmt.Println(math.Min(10, 20))

  //Redondeos
  fmt.Println(math.Floor(3.7)) // Redondea hacia abajo
  fmt.Println(math.Ceil(3.2))  // Redondea hacia arriba
  fmt.Println(math.Round(3.5)) // Redondea al entero más cercano

  // Valor absoluto
  fmt.Println(math.Abs(-5))

  // Exponencial y logaritmos
  fmt.Println(math.Exp(1))  // e^x
  fmt.Println(math.Log(10)) // ln(x)
  fmt.Println(math.Log10(1000))
}