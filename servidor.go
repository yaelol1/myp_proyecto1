package myp_proyecto1

import(
	"fmt"
	"encoding/json"
	"net"
	"net/http"
)


// Servidor estructura que contiene los cuartos y los comandos para interactuar
// con el mismo.
type Servidor struct {
	cuartos  map[string]*cuarto
}

// TODO: net.Listen

// NuevoServidor crea un servidor y devuelve su apuntador
func NuevoServidor() *servidor {

}

func (s Servidor) InicializaServidor() {
	ln, err := net.Listen("tcp", ":8080")
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

// Response acepta las respuestas de los clientes
func (s Servidor) Response(){

}

// Request manda peticiones a los clientes
func (s Servidor) Request(){

}

/// Parser decodifica el Json entrante
func (s Servidor) parser(){

}
