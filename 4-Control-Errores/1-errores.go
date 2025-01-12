package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//Errores/Defer/Panic

func main() {

	str := "123"
	numerito, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Numerito: ", numerito)

	resultado, errorcito := divide(10, 10)
	if errorcito != nil {
		fmt.Println(errorcito)
		return
	}
	fmt.Println(resultado)


	///Defer 
	//Las funciones con defer se agregan a una especie de pila

	file,erro2 := os.Create("hola.txt")
	if erro2 != nil {
		fmt.Println(erro2)
		return
	}

	defer file.Close()

	_,errorazo := file.Write([]byte("Hola escrito en archivo"))
	 
	if  errorazo != nil {
		fmt.Println(errorazo)
		return
	}
	// Si hubiese un error, la siguiente sentencia no se ejecutaria
	//file.Close()

}

// Errores personalizados
func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por 0")
	} else {
		return dividendo / divisor, nil
	}
}
