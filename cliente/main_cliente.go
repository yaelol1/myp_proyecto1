package main

 import (
// 	"bufio"
// 	"log"
	"strings"
	"fmt"
// 	"net"
 )

// instrucciones imprime las instrucciones de las funciones
// del chat.
func instrucciones(){
	fmt.Printf("Para mandar mensaje /msg mensaje \n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
}

// actionTranslator Toma el input del usuario y lo manda a
// request al servidor
func actionTranslator(action string) interface{} {
	actionArr := strings.SplitAfterN(action, " ", 2)

	for index, element := range actionArr {
		fmt.Println(index, element)
	}
	switch actionArr[0]{
		case "/msg":
		r := map[string]interface{}{"type": "PUBLICMESSAGE", "message": actionArr[0]}
		return r

		default:
		fmt.Printf("Comando no válido \n")
	}

	return nil
}

// main crea un cliente y lo conecta al servidor, también le
// da la bienvenida y abre el menú
func main(){
	var mensaje map[string]interface{}
	var action string
	s := NuevoCliente()
	s.Conectar();

	// Menu
	fmt.Printf("Bienvenido al chat \n")
	instrucciones();

	fmt.Printf("Primero esrcibre tu nombre:  \n")
	fmt.Scan(&action)

	// Envía el nombre
	mensaje = map[string]interface{}{"type": "IDENTIFY","username": action}
	s.Request(mensaje)

	// Envía el status como conectado
	mensaje = map[string]interface{}{"type": "STATUS","status": "CONNECTED"}
	s.Request(mensaje)

	for {
		fmt.Scan(&action)
		translated := actionTranslator(action).(map[string]interface{})
		s.Request(translated)
	}

}
