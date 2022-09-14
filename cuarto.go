package myp_proyecto1

//import (
//	"net"
//)

type Cuarto struct {
	name    string
	users  map[string]*cliente
}

// NuevoCuarto crea un cuarto y lo devuelve.
func NuevoCuarto() *Cuarto{

}

// RecibeMensaje recibe el PUBLICMESSAGEFROM.
func (c Cuarto) RecibeMensaje(){

}

// agregaIntegrante agrega al integrante y manda un mensaje de bienvenida.
func (c Cuarto) agregaIntegrante(){
}

// eliminaIntegrante elimina al integrante y manda un mensaje de despedida.
func (c Cuarto) eliminaIntegrante(){
}
