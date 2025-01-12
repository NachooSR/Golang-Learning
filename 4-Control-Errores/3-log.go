package main

import (
	"log"
	"os"
)

func main() {

	//Agregar prefijo al registro del error
	log.SetPrefix("main(): ")

    //Ambos detienen la ejecucion del programa
	// log.Fatal()
	// log.Panic()


	//Registrar errores
	file,err := os.OpenFile("info.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("SOY UN LOG")
   
   
}