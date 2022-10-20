package main

import (
	"encoding/json"
	"net"
	"bufio"
	"fmt"
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
	conn, err := net.Dial("tcp", ":3306")
	if err != nil {
		// handle error

	}

	c.conn = conn
}

func (c *Cliente) lee(){
	// Decodificador que lee directamente desde el socket
	decoder := json.NewDecoder(bufio.NewReader(c.conn))

	// Interfaz, al no saber qué datos tendrá el JSON
	var jsonData interface{}
	for{
		err := decoder.Decode(&jsonData)

		if err != nil {
			fmt.Println(jsonData, err)
			return
			// handle error
		}
		// Se convierte a un mapa
		msg := jsonData.(map[string]interface{})
		c.response(msg)
	}
}

func (c *Cliente) response(msg map[string]interface{}){

	fmt.Println(msg)

	tipo, ok1 := msg["type"] // Checking for existing key and its value
	if !ok1 {
		fmt.Println("Se necesita un tipo")
		return
	}


	switch tipo {
	case "NEW_USER":
	case "NEW_STATUS":
	case "USER_LIST":
	case "MESSAGE_FROM":
	case "PUBLIC_MESSAGE_FROM":
	case "JOINED_ROOM":
	case "ROOM_USER_LIST":
	case "ROOM_MESSAGE_FROM":
	case "LEFT_ROOM":
	case "DISCONECTED":
	default:
		fmt.Print("DEBUG: invalid", msg)
	}
}

// Request manda peticiones a los clientes.
func (c *Cliente) Request(peticion map[string]interface{}){

	d := json.NewEncoder(c.conn)

	if err := d.Encode(peticion); err != nil {
		fmt.Println(err)
	}
}
