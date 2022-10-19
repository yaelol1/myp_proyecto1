<<<<<<< HEAD
package servidor
=======
package main
>>>>>>> feature_servidor

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
=======
	"log"
>>>>>>> feature_servidor
	"net"
)

// Servidor estructura que contiene los cuartos y los comandos para interactuar
// con el mismo.
type Servidor struct {
	cuartos map[string]*Cuarto
}


// NuevoServidor crea un servidor y devuelve su apuntador
func NuevoServidor() *Servidor {
<<<<<<< HEAD
	fmt.Print("Hola desde Nuevo ")
=======
>>>>>>> feature_servidor
	return &Servidor{
		cuartos: make(map[string]*Cuarto),
	}

}

func (s *Servidor) InicializaServidor() {
<<<<<<< HEAD
	fmt.Print("Hola desde inicializa")
	ln, err := net.Listen("tcp", ":8888")
=======
	fmt.Print("Servidor escuchando \n")
	ln, err := net.Listen("tcp", ":3306")
>>>>>>> feature_servidor
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
<<<<<<< HEAD
			// handle error
=======
			log.Printf("failed to accept connection: %s", err.Error())
			continue
>>>>>>> feature_servidor
		}
		go s.handleConnection(conn)
	}
}

// handleConnection acepta las conexiones y decide qué hacer con ellas
func (s *Servidor) handleConnection(conn net.Conn) {
<<<<<<< HEAD
	// Decodificador que lee directamente desde el socket
	d := json.NewDecoder(conn)

	var msg Mensaje

	err := d.Decode(&msg)
	fmt.Println(msg, err)

	if err != nil {
		// handle error
	}

	s.Response(msg, conn)
}

// Response acepta las respuestas de los clientes
func (s *Servidor) Response(msg Mensaje, conn net.Conn) {
	switch msg.tipo{
		case "ROOM_MESSAGE":
		r, ok := s.cuartos[msg.roomName]
		if !ok {
			r =NuevoCuarto(msg.roomName)
			s.cuartos[msg.roomName] = r
		}
	}
}
=======

	// Decodificador que lee directamente desde el socket
	decoder := json.NewDecoder(conn)

	// Interfaz, al no saber qué datos tendrá el JSON
	for{
		var jsonData interface{}
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

	d := json.NewEncoder(conn)

	if err := d.Encode(msg); err != nil {
		fmt.Println(err)
	}

	tipo, ok1 := msg["type"] // Checking for existing key and its value
	if !ok1 {
		panic("Type needed")
	}

	switch tipo {
	case "IDENTIFY":
		// Nombre
	case "STATUS":
		// publicar en el cuartos
	case "USERS":

	case "MESSAGE":
		// mensaje personal

	case "PUBLIC_MESSAGE":

	case "NEW_ROOM":
		// { "type": "CREATEROOM",
		// "roomname": "Sala 1" }
	case "INVITE":
		// "type": "INVITE",
		// "roomname": "Sala 1",
		// "users": [ "Luis", "Antonio", "Fernando" ]
	case "JOIN_ROOM":
		// { "type": "JOINROOM",
		// 	"roomname": "Sala 1" }
	case "ROOM_USERS":
		// { "type": "ROOMUSERS",
		// 	"roomname": "Sala 1" }
	case "ROOM_MESSAGE":
		nombreCuarto := msg["roomname"].(string)
		r, ok := s.cuartos[nombreCuarto]

		if !ok {
			r =NuevoCuarto(nombreCuarto)
			s.cuartos[nombreCuarto] = r
		}
		// "type": "ROOMESSAGE",
		// "roomname": "Sala 1",
		// "message": "¡Hola sala 1!" }
	case "LEAVE_ROOM":
		// "type": "LEAVEROOM",
		// "roomname": "Sala 1" }
	case "DISCONNECT":
		// { "type": "DISCONNECT" }
	default:
		fmt.Print("invalid", msg)

	}
}

>>>>>>> feature_servidor
