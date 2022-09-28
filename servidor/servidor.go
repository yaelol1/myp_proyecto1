package servidor

import (
	"encoding/json"
	"fmt"
	"net"
	"github.com/yaelol1/myp_proyecto1/recursos"
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
	ln, err := net.Listen("tcp", ":1252")
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
			r = &Cuarto{
				name:    msg.roomName,
				members: make(map[net.Addr]*client),
			}
			s.cuartos[msg.roomName] = r
		}
	}
}
