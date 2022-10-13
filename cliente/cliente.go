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

// TODO: net.Dial -> connection -> Write

// NuevoCliente crea el cliente y lo devuelve
func NuevoCliente() *Cliente {
	fmt.Printf("holaaaa nuevo \n")
	return &Cliente{
		nombre: "Yael",
		cuartos: make(map[string]*string),
		conn:  nil,
	}
}

// Conectar conecta al cliente a un puerto
func (c *Cliente) Conectar(){
	fmt.Printf("HOlaaaa conectar \n")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		// handle error
	}

	// status, err := bufio.NewReader(conn).ReadString('\n')
	// if false {
	// 	fmt.Println(status, err)
	// }
	c.conn = conn
}


// Request manda peticiones a los clientes
func (c *Cliente) Request(){

	d := json.NewEncoder(c.conn)
	fmt.Printf("HOlaaaa request\n")


	 var msg string
	 msg = " {\"type\": \"ROOM_MESSAGE\", \"roomname\": \"Sala 1\", \"message\": \"¡Hola sala 1!\" }"
	 c.conn.Write([]byte("holaaa me mandé desde cliente"))

	 c.conn.Write([]byte(msg))
	 err := d.Encode("Holaaaa me mandé desde cliente")
	 fmt.Println(msg, err)

	 if err != nil {
	  	// handle error
	 }
}