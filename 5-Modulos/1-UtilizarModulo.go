package main

import (
	"fmt"

	"github.com/NachooSR/greetings"
)


func main(){

	mensaje := greetings.Hello("Nacho")
    fmt.Println(mensaje)
}