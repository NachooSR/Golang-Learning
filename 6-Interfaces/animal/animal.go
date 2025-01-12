package animal

import "fmt"

//Una interfaz es un contrato, cada tipo de dato deben cumplir con esto, es decir, utilizar este metodo
type Animal interface{
	Sonido()
}

func HacerSonido (a Animal){
    a.Sonido()
}
///Estructuras
type Perro struct{
	Nombre string
}

type Gato struct{
	Nombre string
}

func (p Perro) Sonido(){
	fmt.Println(p.Nombre + " WOOF")
}

func (g Gato) Sonido(){
	fmt.Println(g.Nombre + " MIAU")
}