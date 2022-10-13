package main

 import (
// 	"bufio"
// 	"log"
	// "fmt"
// 	"net"
 )

func main(){
	s := NuevoCliente()
	s.Conectar();
	s.Request();
}
