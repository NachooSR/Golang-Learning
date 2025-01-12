package main

import "fmt"

func main (){


   /*****NUMEROS*****/
  var numerito int8 //8 Bits -128/127
  
  var numerito2 int //(Si no especificamos, la longitud depende del sist. operativo)

  var numero uint (solo numeros positivos)

  var decimal float32
  var decimal float64


  /*****BOOLEANS*****/
  var trueOfalse bool


  /*****STRINGS*****/
  var nombre string
  //Si quiero que las comillas aparezcan en mi mensaje \"algo\"

  /*****BYTE*****/
  var a byte = 'a'
  //Imprime el valor ascci de la a


  /*****RUNE*****/

  //A diferencia de byte, permite trabajar con un cacacter Individual. EJEMPLO

  
  s := "Golang 😃"

  //El string se transforma en una lista de runas
  runas := []rune(s)

  // Iteramos sobre cada rune
  for i, r := range runas {
	  fmt.Printf("Índice: %d, Rune: %c, Unicode: %U\n", i, r, r)
  }

  // Accedemos a un rune específico
  fmt.Println("Primer rune:", string(runas[0])) // Convierte rune a string
  
}