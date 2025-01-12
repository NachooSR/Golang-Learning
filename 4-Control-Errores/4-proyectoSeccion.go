package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	
)

type Contacto struct {
	Id          int    `json: "id"`
	Nombre      string `json: "name"`
	Apellido    string `json: "apellido"`
	Mail        string `json: "mail"`
	NumTelefono string `json: "numTelefono"`
}

// Guardar contactos en un archivo json

// ****************FUNCIONES DEL ARCHIVO************//
func saveContacts(contacts []Contacto) error {
	file, err := os.Create("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(contacts)

	if err != nil {
		return err
	}
	return nil
}

func loadContacts(contactos *[]Contacto) error {

	file, err := os.Open("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(contactos)
	if err != nil {
		return err
	}
	return nil
}

var idAuto int = 1

func main() {

	

	/*Aunque la ejecucion del programa no se corte, al iniciar la app el archivo esta vacio
	  por lo que se va arrojar el error EOF (end of file)
	  La solucion, una variable global que empieza en 0 por lo que no vamos a preguntar si hay error. Una vez
	  que agregamos un contacto, aumenta el contador en 1, y ahora si verificamos el error     */

	/*Nueva solucion funcion que verifica si hay contenido en el archivo*/

	var newContacs []Contacto
	booleancito,err := archivoVacio("contacts.json")
	if booleancito == false  && err != nil {
		fmt.Println("Error en el archivo: ",err)
	}else{
		err = loadContacts(&newContacs) //Con esto evito repetir codigo y queda cargado el primer arreglo
		if err !=nil {
			fmt.Println("Error al cargar los contactos: ",err)
		}
	}

	var option int


	for option != 6 {

		menu()
		fmt.Scanln(&option)

		switch option {

		case 1:
			listarContactos(newContacs)

		case 2:
			var temporal Contacto
			var opcion int
			for {
				temporal = addContact()

				// contactosAgregados++ //Ahora que hay minimo un contacto agregado, el error va ser verificado en el inicio

				idAuto++

				newContacs = append(newContacs, temporal)
				fmt.Println("Desea agregar otro contacto: 1-Si 2-No")
				fmt.Scanln(&opcion)

				if opcion == 2 {
					break
				}

			}

			err := saveContacts(newContacs)
			if err != nil {
				fmt.Println("Error al cargar usuarios: ", err)
			} else {
				fmt.Println("Contacto agregado con exito")
			}

		case 3:
			
			var indice int
			listarContactos(newContacs)

			fmt.Println("Ingrese el numero del contacto a borrar: ")
			fmt.Scanln(&indice)
			eliminarContacto(&newContacs, indice-1)
			saveContacts(newContacs)
			listarContactos(newContacs)

		case 4:

		case 5:

		case 6:

		default:
			fmt.Println("Opcion incorrecta, intente nuevamente")
		}

	}

}

func addContact() Contacto {

	var aux Contacto

	reader := bufio.NewReader(os.Stdin)

	aux.Id = idAuto

	fmt.Print("Nombre:")
	aux.Nombre, _ = reader.ReadString('\n')
	aux.Nombre = strings.TrimSpace(aux.Nombre)

	fmt.Print("Apellido:")
	aux.Apellido, _ = reader.ReadString('\n')
	aux.Apellido = strings.TrimSpace(aux.Apellido)

	fmt.Print("Mail:")
	aux.Mail, _ = reader.ReadString('\n')
	aux.Mail = strings.TrimSpace(aux.Mail)

	fmt.Print("Numero:")
	aux.NumTelefono, _ = reader.ReadString('\n')
	aux.NumTelefono = strings.TrimSpace(aux.NumTelefono)

	return aux
}

func menu() {
	fmt.Println(" _____________________")
	fmt.Println("|1-Listar contactos   |")
	fmt.Println("|2-Agregar  contacto  |")
	fmt.Println("|3-Eliminar contacto  |")
	fmt.Println("|4-Modificar contacto |")
	fmt.Println("|5-Buscar contacto    |")
	fmt.Println("|6-Salir              |")
}

func listarContactos(conctactos []Contacto) {

	for i := 0; i < len(conctactos); i++ {

		mostrarContacto(conctactos[i])
		fmt.Println("=============")
	}

}

func mostrarContacto(contacto Contacto) {
	fmt.Printf("\n%d)Nombre:%s %s \n", contacto.Id, contacto.Nombre, contacto.Apellido)
	fmt.Println("Numero: ", contacto.NumTelefono)
	fmt.Println("Mail: ", contacto.Mail)
}

func eliminarContacto(contactos *[]Contacto, indice int) {
	*contactos = slices.Delete(*contactos, indice, indice+1)
}

// FUNCION PARA VERIFICAR SI EL ARCHIVO TIENE CONTENIDO
func archivoVacio(filename string) (bool, error) {

	info, err := os.Stat(filename)

	if err != nil {
		return false, err
	}

	return info.Size() != 0, nil

}
