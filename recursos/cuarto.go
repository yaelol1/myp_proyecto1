package recursos

import (
	"github.com/yaelol1/myp_proyecto1/cliente"
	"net"
)
type Cuarto struct {
	name   string
	users  map[net.Adrr]*Cliente
}

// NuevoCuarto crea un cuarto y lo devuelve.
func NuevoCuarto(name string) *Cuarto{
	return &Cuarto{
		name: name,
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
