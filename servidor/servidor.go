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
	users  map[string]net.Conn
}


// NuevoServidor crea un servidor y devuelve su apuntador
func NuevoServidor() *Servidor {
	return &Servidor{
		cuartos: make(map[string]*Cuarto),
		users:  make(map[string]net.Conn),
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

	case "ERROR":
	case "WARNING":
	case "INFO":
	case "IDENTIFY":
		usuario, ok1 := msg["username"] // Checking for existing key and its value
		if !ok1 {
			break
		}
		usuarioS := usuario.(string)
		s.users[usuarioS] = conn
		s.cuartos["General"].agregaIntegrante(conn, usuarioS)

	case "STATUS":
		general := s.cuartos["General"]
		userName:=general.obtenNombre(conn)

		response:= map[string]interface{} {"type": "NEW_USER",
			"username": userName}

		general.Broadcast(conn, response)

	case "USERS":

		users := "[ "
		for k, _ := range s.users {
			users += "\" " + k + " \","
		}
		users = users[:len(users)-1]
		users += " ]"
		selfResponse:= map[string]interface{} {"type": "USER_LIST",
			"usernames": users}
		s.send(conn, selfResponse)

		fmt.Print("\n Response: ", users)

	case "MESSAGE":
		user, ok := msg["username"] // Checking for existing key and its value
		if !ok {
			return
		}
		connRecipient, existe := s.users[user.(string)]
		if !existe {
			selfResponse:= map[string]interface{} {"type": "WARNING",
				"message": "El usuario '"+user.(string)+"' no existe",
				"operation":"MESSAGE",
				"username": user.(string)}
			s.send(conn, selfResponse)
			return
		}

		userName:=  s.cuartos["General"].obtenNombre(conn)
		Response:= map[string]interface{} {"type": "MESSAGE",
			"usernames": userName,
			"message": msg["message"]}
		s.send(connRecipient, Response)



	case "PUBLIC_MESSAGE":
		general := s.cuartos["General"]
		userName:=general.obtenNombre(conn)

		response:= map[string]interface{} {"type": "PUBLIC_MESSAGE_FROM",
			"username": userName,
			"message": msg["message"]}
		general.Broadcast(conn, response)

		selfResponse:= map[string]interface{} {"type": "INFO",
			"message": "success",
			"operation":"IDENTIFY"}
		s.send(conn, selfResponse)

	case "NEW_ROOM":
		roomname := msg["rooomname"].(string)
		_, ok := s.cuartos[roomname]
		if !ok {
			s.cuartos[msg["roomnane"].(string)] = NuevoCuarto(roomname)

			selfResponse:= map[string]interface{} {"type": "INFO",
				"message": "success",
				"operation":"NEW_ROOM"}
			s.send(conn, selfResponse)
			return
		}

		selfResponse:= map[string]interface{} {"type": "WARNING",
			"message": "El cuarto '"+roomname+"' ya existe",
			"operation":"NEW_ROOM"}
		s.send(conn, selfResponse)

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

func (s *Servidor) send(conn net.Conn, msg map[string]interface{}){
	d := json.NewEncoder(conn)
	if err := d.Encode(msg); err != nil {
		fmt.Println(err)
	}
}

func toList (usersToConvert map[string]interface{}) string{

	users := "[ "
	for k, _ := range usersToConvert {
		users += "\" " + k + " \","
	}
	users = users[:len(users)-1]
	users += " ]"

	return users
}
