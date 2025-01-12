package main

import (
	"fmt"
	"time"
)

func main() {

	//Arroja fecha y hora actual
	t := time.Now()

	hora:= t.Hour()
	
    // if condicion
	//Anidados

	if hora < 12  {
		fmt.Println("Es de manana")
	}else if hora < 17{
		fmt.Println("Es tarde")
	}else{
		fmt.Println("Noche")
	}


}