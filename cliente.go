package myp_proyecto1

import (
	"encoding/json"
	"net"
)

type Cliente struct {
	nombre   string
	cuartos  map[string]*cuarto
}

// TODO: net.Dial -> connection -> Write

// NuevoCliente crea el cliente y lo devuelve
func NuevoCliente() *Cliente {
}

// Conectar conecta al cliente a un puerto
func (c Cliente) Conectar(){
	conn, err := net.Dial("tcp", ":1252")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
}


// Request manda peticiones a los clientes
func (c Cliente) Request(){

}
