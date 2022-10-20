package main

 import (
	"bufio"
	"os"
	"strings"
	"fmt"
 )

// instrucciones imprime las instrucciones de las funciones
// del chat.
func instrucciones(){
	fmt.Printf("Para mandar mensaje publico /msgAll mensaje \n")
	fmt.Printf("Para mandar mensaje privado /msg usuario mensaje \n")
	fmt.Printf("Para mandar mensaje a un cuarto /msg cuarto mensaje \n")
	fmt.Printf("Para mandar invitación a un cuarto /msg cuarto mensaje \n")
	fmt.Printf("Para crear un cuarto /room nombreCuarto usuario1 usuario2 ... (Los usuarios son opcionales)\n")
	fmt.Printf("Para invitar alguien a un cuarto /room nombreCuarto usuario1 usuario2 ... \n")
	fmt.Printf("Para acepar una invitación a un cuarto /accept nombreCuarto \n")
	fmt.Printf("Para pedir una lista de todos los usuarios /list \n")
	fmt.Printf("Para pedir una lista de todos los usuarios en un cuarto /list nombreCuarto\n")
	fmt.Printf("Para abandonar un cuarto /leave nombreCuarto \n")
	fmt.Printf("Para cerrar la aplicación /disconnect \n")
	fmt.Printf("Para imprimir las instrucciones otra vez /info \n")
}

// actionTranslator Toma el input del usuario y lo manda a
// request al servidor
func actionTranslator(action string) interface{} {
	actionArr := strings.SplitN(action, " ",2)


	switch actionArr[0]{
		case "/msgAll":
		r := map[string]interface{}{"type": "PUBLIC_MESSAGE", "message": actionArr[1]}
		return r

		case "/msg":

		recipAndMsg := strings.SplitN(actionArr[1], " ",2)
		r := map[string]interface{}{"type": "MESSAGE", "username": recipAndMsg[0],"message": recipAndMsg[1]}
		return r

		case "/info":
		instrucciones()

		case "/disconnect":
		os.Exit(0)

		default:
		fmt.Printf(actionArr[0])
		fmt.Printf("\n Comando no válido, para imprimir la lista de comandos escriba /info \n")
	}

	return nil
}

// main crea un cliente y lo conecta al servidor, también le
// da la bienvenida y abre el menú
func main(){
	var mensaje map[string]interface{}
	var action string
	c := NuevoCliente()
	c.Conectar()
	go c.lee()
	reader := bufio.NewReader(os.Stdin)

	// Menu
	fmt.Printf("Bienvenido al chat \n")
	instrucciones()

	fmt.Printf("\n \nPrimero esrcibre tu nombre:  \n")

	action, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}

	// remove the delimeter from the string
	action = strings.TrimSuffix(action, "\n")

	// Envía el nombre
	mensaje = map[string]interface{}{"type": "IDENTIFY","username": action}
	c.Request(mensaje)

	// Envía el status como conectado
	mensaje = map[string]interface{}{"type": "STATUS","status": "CONNECTED"}
	c.Request(mensaje)

	for {
		// Pide la acción nueva
		action, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
		}

		action = strings.TrimSuffix(action, "\n")

		translated := actionTranslator(action)
		if translated != nil{
			c.Request(translated.(map[string]interface{}))
		}
	}

}
