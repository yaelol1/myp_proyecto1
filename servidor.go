package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// Servidor estructura que contiene los cuartos y los comandos para interactuar
// con el mismo.
type Servidor struct {
	cuartos map[string]*Cuarto
}


// NuevoServidor crea un servidor y devuelve su apuntador
func NuevoServidor() *Servidor {
	fmt.Print("Hola desde Nuevo ")
	return &Servidor{
		cuartos: make(map[string]*Cuarto),
	}

}

func (s *Servidor) InicializaServidor() {
	fmt.Print("Hola desde inicializa")
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go s.handleConnection(conn)
	}
}

// handleConnection acepta las conexiones y decide qué hacer con ellas
func (s *Servidor) handleConnection(conn net.Conn) {
	// Decodificador que lee directamente desde el socket
	decoder := json.NewDecoder(conn)

	// Interfaz, al no saber qué datos tendrá el JSON
	var jsonData interface{}
	err := decoder.Decode(&jsonData)

	if err != nil {
		fmt.Println(jsonData, err)
		// handle error
	}

	// Se convierte a un mapa
	msg := jsonData.(map[string]interface{})
	s.Response(msg, conn)
}

// Response acepta las respuestas de los clientes
func (s *Servidor) Response(msg map[string]interface{} , conn net.Conn) {
	fmt.Print("Response")

	tipo, ok1 := msg["type"] // Checking for existing key and its value
	if !ok1 {
		panic("Type needed")
	}

	switch tipo {
	case "ROOM_MESSAGE":
		nombreCuarto := msg["roomname"].(string)
		r, ok := s.cuartos[nombreCuarto]

		if !ok {
			r =NuevoCuarto(nombreCuarto)
			s.cuartos[nombreCuarto] = r
		}

	case "STATUS":
		// publicar en el cuartos
	case "IDENTIFY":
		// Nombre
	case "MESSAGE":
		// mensaje personal
	case "CREATEROOM":
		// { "type": "CREATEROOM",
		// "roomname": "Sala 1" }
	case "INVITE":
		// "type": "INVITE",
		// "roomname": "Sala 1",
		// "users": [ "Luis", "Antonio", "Fernando" ]
	case "JOINROOM":
		// { "type": "JOINROOM",
		// 	"roomname": "Sala 1" }
	case "ROOMUSERS":
		// { "type": "ROOMUSERS",
		// 	"roomname": "Sala 1" }
	case "ROOMESSAGE":
		// "type": "ROOMESSAGE",
		// "roomname": "Sala 1",
		// "message": "¡Hola sala 1!" }
	case "LEAVEROOM":
		// "type": "LEAVEROOM",
		// "roomname": "Sala 1" }
	case "DISCONNECT":
		// { "type": "DISCONNECT" }
	case default: 

	}
}

