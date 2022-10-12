package main

 import (
// 	"bufio"
// 	"log"
	"fmt"
	"github.com/yaelol1/myp_proyecto1/servidor"
// 	"net"
 )

func main(){
	fmt.Print("Hola")
	var s servidor.Servidor
	s = servidor.NuevoServidor()
	s.InicializaServidor();
}
