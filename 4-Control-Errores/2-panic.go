package main

import "fmt"

func main() {

	division(10,2)
	division(10,5)
	division(20,0)

	//Antes sin recover la siguiente linea no se ejecutaba debido a que se corta el hilo de ejecucion
	division(20,4)
}

func division(dividendo, divisor int) {
   
	//Las funciones recover, recuperan los panic que ocurran
	defer func (){
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	validateZero(divisor)
	fmt.Println(dividendo/divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por 0")
	}
}