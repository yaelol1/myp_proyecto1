package main

import (
	"encoding/json"
	"net"
	"fmt"
)

type Cuarto struct {
	name   string
	users  map[net.Conn]string
}

// NuevoCuarto crea un cuarto y lo devuelve.
func NuevoCuarto(name string) *Cuarto{
	return &Cuarto{
		name: name,
		users: make(map[net.Conn]string),
	}
}

// RecibeMensaje recibe el PUBLICMESSAGEFROM.
func (c *Cuarto) RecibeMensaje(){

}

// agregaIntegrante agrega al integrante y manda un mensaje de bienvenida.
func (c *Cuarto) agregaIntegrante(dir net.Conn, name string){
	c.users[dir] = name
}

// eliminaIntegrante elimina al integrante y manda un mensaje de despedida.
func (c *Cuarto) eliminaIntegrante(){
}

func (c *Cuarto) Broadcast(conn net.Conn, mensaje string){

	msg := map[string]string{"type": "PUBLIC_MESSAGE_FROM", "username": c.users[conn], "message": mensaje}

	for addr, _ := range c.users {
		d := json.NewEncoder(addr)
		if err := d.Encode(msg); err != nil {
			fmt.Println(err)
		}
	}

}
