package myp_proyecto1

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

}

func (s Servidor) InicializaServidor() {
	ln, err := net.Listen("tcp", ":1252")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

// handleConnection acepta las conexiones y decide qué hacer con ellas
func (s Servidor) handleConnection(conn net.Conn) {
	// Decodificador que lee directamente desde el socket
	d := json.NewDecoder(c)

	var msg Mensaje

	err := d.Decode(&msg)
	fmt.Println(msg, err)

	if err != nil {
		// handle error
	}

	Response(msg)
}

// Response acepta las respuestas de los clientes
func (s Servidor) Response(msg Mensaje) {

}
