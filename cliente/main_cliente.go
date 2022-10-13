package main

 import (
// 	"bufio"
// 	"log"
	"fmt"
// 	"net"
 )

func main(){
	s := NuevoCliente()
	s.Conectar();

	// Menu
	var action string
	for {
		fmt.Scan(&action)
	}

}
