package main

import "fmt"

type Humano struct {
	nombre string
	edad   int
	correo string
}

// Para que se vuelva un metodo la funcion debe tener un receptor ()--> antes del nombre de la misma

/*Si no debemos modificar algun campo de la estructura
el metodo no necesita si o si utilizar un puntero y puede trabajar con la copia
*/
func (h Humano) diHola() {
	fmt.Println("Hola mi nombre es, ",h.nombre)
}

func main() {
   
	h:=Humano{"Carlos",20,"carlos@gmail.com"}

	h.diHola()
}
