package recursos

import (
	"net"
	"github.com/yaelol1/myp_proyecto1/cliente"
)

type Cuarto struct {
	name   string
	users  map[net.Addr]*cliente.Cliente
}

// NuevoCuarto crea un cuarto y lo devuelve.
func NuevoCuarto(name string) *Cuarto{
	return &Cuarto{
		name: name,
		users: make(map[net.Addr]*Cliente),
	}
}

// RecibeMensaje recibe el PUBLICMESSAGEFROM.
func (c *Cuarto) RecibeMensaje(){

}

// agregaIntegrante agrega al integrante y manda un mensaje de bienvenida.
func (c *Cuarto) agregaIntegrante(){
}

// eliminaIntegrante elimina al integrante y manda un mensaje de despedida.
func (c *Cuarto) eliminaIntegrante(){
}
