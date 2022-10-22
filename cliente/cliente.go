// Package main crea un cliente que se conecta a un servidor
// envía y recibe mensajes con el protocolo
package main

import (
	"encoding/json"
	"net"
	"bufio"
	"fmt"
)

type Cliente struct {
	nombre   string
	conn net.Conn
}

// NuevoCliente crea el cliente y lo devuelve.
func NuevoCliente() *Cliente {
	return &Cliente{
		nombre: "Yael",
		conn:  nil,
	}
}

// Conectar conecta al cliente a un puerto.
func (c *Cliente) Conectar(){
	conn, err := net.Dial("tcp", ":3306")
	if err != nil {
		// handle error
		panic("Servidor no se pudo conectar")
	}

	c.conn = conn
}

// lee Responde a todos los mensajes entrantes del servidor
func (c *Cliente) lee(){
	// Decodificador que lee directamente desde el socket
	decoder := json.NewDecoder(bufio.NewReader(c.conn))

	// Interfaz, al no saber qué datos tendrá el JSON
	var jsonData interface{}
	for{
		err := decoder.Decode(&jsonData)

		if err != nil {
			fmt.Println(jsonData, err)
			continue
			// handle error
		}
		// Se convierte a un mapa
		msg := jsonData.(map[string]interface{})
		c.response(msg)
	}
}

// response Decide qué hacer con el mensaje del servidor
func (c *Cliente) response(msg map[string]interface{}){
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	// fmt.Print("DEBUG: ", msg, "\n")

	switch msg["type"] {
	case "ERROR":
		fallthrough
	case "WARNING":
		fallthrough
	case "INVITATION":
		fallthrough
	case "INFO":
		fmt.Println(msg["message"].(string))
	case "NEW_USER":
		fmt.Println("Nuevo usuario: ",msg["username"].(string))
	case "NEW_STATUS":
		fmt.Println("Usuario: ",msg["username"].(string),
			" cambió su estado a '", msg["status"].(string), "'")
	case "JOINED_ROOM":
		fmt.Println("(", msg["roomname"].(string), ") ", msg["username"].(string),
			" se unió al chat ")
	case "ROOM_USER_LIST":
		fmt.Print("(Del cuarto ) ")
		fallthrough
	case "USER_LIST":
		integrantesRaw := msg["usernames"]
		integrantesRaw2 := integrantesRaw.([]interface{})

		users := ""
		for _, user := range integrantesRaw2 {
			users += user.(string) + ", "
		}
		users = users[:len(users)-2]
		users += "."
		fmt.Println("todos los usuarios son: ", users)

	case "PUBLIC_MESSAGE_FROM":
		fmt.Println("(público) ", msg["username"].(string),
			": ", msg["message"].(string))
	case "ROOM_MESSAGE_FROM":
		fmt.Println("(", msg["roomname"].(string), ") ", msg["username"].(string),
			": ", msg["message"].(string))
	case "MESSAGE_FROM":
		fmt.Println("[susurro] ", msg["username"].(string),
			": ", msg["message"].(string))
	case "LEFT_ROOM":
		fmt.Println("(", msg["roomname"].(string), ") ", msg["username"].(string),
			" se salió del chat ")
	case "DISCONNECTED":
		fmt.Println(msg["username"].(string), " se desconectó")
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
