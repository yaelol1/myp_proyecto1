package main

import (
	"encoding/json"
	"net"
	"fmt"
	"bufio"
)

type Cliente struct {
	nombre   string
	cuartos map[net.Addr]*Cuarto
	conn net.Conn
}

// TODO: net.Dial -> connection -> Write

// NuevoCliente crea el cliente y lo devuelve
func NuevoCliente() *Cliente {
	return &Cliente{
		nombre: "Yael",
		cuartos: make(map[net.Addr]*Cuarto),
	}
}

// Conectar conecta al cliente a un puerto
func (c *Cliente) Conectar(){
	conn, err := net.Dial("tcp", ":1252")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status, err)
}


// Request manda peticiones a los clientes
func (c *Cliente) Request(){

	d := json.NewEncoder(c.conn)

	var msg Mensaje

	err := d.Encode(&msg)
	fmt.Println(msg, err)

	if err != nil {
		// handle error
	}
}
