package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// Servidor estructura que contiene los cuartos y los comandos para interactuar
// con el mismo.
type Servidor struct {
	cuartos map[string]*Cuarto
}


// NuevoServidor crea un servidor y devuelve su apuntador
func NuevoServidor() *Servidor {
	return &Servidor{
		cuartos: make(map[string]*Cuarto),
	}

}

func (s *Servidor) InicializaServidor() {
	fmt.Print("Servidor escuchando \n")
	ln, err := net.Listen("tcp", ":3306")
	if err != nil {
		// handle error
	}

	general := NuevoCuarto("general")
	s.cuartos["General"] = general
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		// add accepted connection to the general room
		general.agregaIntegrante(conn, "")

		// handle input from connection
		go s.handleConnection(conn)
	}
}

// handleConnection acepta las conexiones y decide qué hacer con ellas
func (s *Servidor) handleConnection(conn net.Conn) {

	// Decodificador que lee directamente desde el socket
	decoder := json.NewDecoder(conn)

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
		s.Response(msg, conn)
	}
}

// Response acepta las respuestas de los clientes
func (s *Servidor) Response(msg map[string]interface{} , conn net.Conn) {
	fmt.Print("\n Response: ", msg)


	tipo, ok1 := msg["type"] // Checking for existing key and its value
	if !ok1 {
		panic("Type needed")
	}

	switch tipo {
	case "IDENTIFY":
	case "STATUS":
	case "USERS":
	case "MESSAGE":
	case "PUBLIC_MESSAGE":
		general := s.cuartos["General"]
		general.Broadcast(conn, msg["message"].(string))

	case "NEW_ROOM":
	case "INVITE":
	case "JOIN_ROOM":
	case "ROOM_USERS":
	case "ROOM_MESSAGE":
		nombreCuarto := msg["roomname"].(string)
		r, ok := s.cuartos[nombreCuarto]

		if !ok {
			break
		}
		fmt.Println(r)
	case "LEAVE_ROOM":
	case "DISCONNECT":
	default:
		fmt.Print("invalid", msg)

	}
}

