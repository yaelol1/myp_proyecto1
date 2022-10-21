package main

import (
	"encoding/json"
	"net"
	"fmt"
)

type Cuarto struct {
	name   string
	users  map[net.Conn]*User
	invites map[string]bool
}

// NuevoCuarto crea un cuarto y lo devuelve.
func NuevoCuarto(name string) *Cuarto{
	return &Cuarto{
		name: name,
		users: make(map[net.Conn]*User),
		invites: make(map[string]bool),
	}
}

// userList manda una lista con todos los usuarios.
func (c *Cuarto) userList() []string {
	users := make([]string, 0)
	for _, n := range c.users {
		users = append(users, n.userName)
	}
	return users
}

// RecibeMensaje recibe una conexión y devuelve el nombre de la persona.
func (c *Cuarto) obtenNombre(conn net.Conn) string{
	return c.users[conn].userName
}

// agregaIntegrante agrega al integrante al mapa de usuarios.
func (c *Cuarto) agregaIntegrante(conn net.Conn, user *User){
	delete(c.invites, user.userName)
	c.users[conn] = user
}

// invitaIntegrante agrega al integrante a la lista de invitados.
func (c *Cuarto) invitaIntegrante(userName string){
	c.invites[userName] = true
}


// fueInvitado verifica si un usuario fue previamente invitado.
func (c *Cuarto) fueInvitado(userName string) bool{
	_, invitado := c.invites[userName]
	return invitado
}

// esIntegrante
func (c *Cuarto) esIntegrante(conn net.Conn) bool{
	_, esIntegrante := c.users[conn]
	return esIntegrante
}

// eliminaIntegrante elimina al integrante y manda un mensaje de despedida.
func (c *Cuarto) eliminaIntegrante(conn net.Conn){
	delete(c.users, conn)
}

// Broadcast envía un mensaje a todas las personas en el cuarto menos al que envía el mensaje.
func (c *Cuarto) Broadcast(conn net.Conn, mensaje map[string]interface{}){

	for addr, _ := range c.users {
		if addr != conn {
			d := json.NewEncoder(addr)
			if err := d.Encode(mensaje); err != nil {
				fmt.Println(err)
			}
		}
	}

}
