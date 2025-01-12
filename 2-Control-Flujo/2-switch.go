package main

import (
	"fmt"
	"runtime"
)

func main() {

	os := runtime.GOOS

    //switch (variable){
	//casos            }


	switch os {

	case "windows":
		fmt.Println("Se esta ejecutando en windows")
		break //No es necesario el break

	case "linux":
		fmt.Println("Se esta ejecutando en Linux")
		break

	case "darwin":
		fmt.Println("Se esta ejecutando en MacOs")
		break

	default: fmt.Println("Otro sistema operativo")
	}

	//En los cases tambien se pueden poner condiciones
	

}