package main

import (
	"encoding/json"
	"net"
	"fmt"
	//"bufio"
)

type Cliente struct {
	nombre   string
	cuartos map[string]*string
	conn net.Conn
}

// NuevoCliente crea el cliente y lo devuelve.
func NuevoCliente() *Cliente {
	return &Cliente{
		nombre: "Yael",
		cuartos: make(map[string]*string),
		conn:  nil,
	}
}

// Conectar conecta al cliente a un puerto.
func (c *Cliente) Conectar(){
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		// handle error
	}

	c.conn = conn
}


// Request manda peticiones a los clientes.
func (c *Cliente) Request(peticion map[string]interface{}){

	d := json.NewEncoder(c.conn)

	if err := d.Encode(peticion); err != nil {
		fmt.Println(err)
	}
}
