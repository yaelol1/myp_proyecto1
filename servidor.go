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
	d := json.NewDecoder(conn)

	// Interfaz, al no saber qué datos tendrá el JSON
	var f interface{}
	err := d.Decode(&f)

	if err != nil {
		fmt.Println(f, err)
		// handle error
	}

	// Se convierte a un mapa
	m := f.(map[string]interface{})
	s.Response(m, conn)
}

// Response acepta las respuestas de los clientes
func (s *Servidor) Response(msg map[string]interface{} , conn net.Conn) {
	fmt.Print("Response")

	// val1, ok1 := msg["user"] // Checking for existing key and its value
	// fmt.Println(val1, ok1)

	// switch msg[tipo]{
	// 	case "ROOM_MESSAGE":
	// 	fmt.Print("ROOM_MESSAGE")

	// 	fmt.Print(msg.message)
	// 	r, ok := s.cuartos[msg.roomName]
	// 	if !ok {
	// 		r =NuevoCuarto(msg.roomName)
	// 		s.cuartos[msg.roomName] = r
	// 	}
	// }
}
