/*Estructuras son similares a las clases en otros lenguajes*/
package main

//import "github.com/NachooSR/Golang-Learning/6-Interfaces/libro"

import "github.com/NachooSR/Golang-Learning/6-Interfaces/animal"

func main() {

	// var librito = libro.NuevoLibro("Luna de Pluton","Dross",345)
	// var libro2 = libro.NuevoTextoLibro("Ciencias Naturales","Isaac Newton","Santillana","Sexto",380)

	// libro.Print(librito)
	// libro.Print(libro2)

	var perrito = animal.Perro{"Simon"}
	var gatito = animal.Gato{"Tito"}

	animal.HacerSonido(perrito)
	animal.HacerSonido(gatito)

	animales := []animal.Animal{
		animal.Gato{"Tito"},
		animal.Perro{"Simon"},
		animal.Gato{"Milo"},
		animal.Perro{"Bachicha"},
		animal.Perro{"Lulu"},
		animal.Perro{"Cielo"},
		animal.Perro{"Santos"},
		animal.Perro{"Mushu"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
