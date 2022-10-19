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
	actionArr := strings.SplitN(action, " ",2)


	switch actionArr[0]{
		case "/msg":
		r := map[string]interface{}{"type": "PUBLIC_MESSAGE", "message": actionArr[1]}
		return r

		default:
		fmt.Printf(actionArr[0])
		fmt.Printf("Comando no váliodo \n")
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

	fmt.Printf("Primero esrcibre tu nombre:  \n")

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

		 translated := actionTranslator(action).(map[string]interface{})
		 c.Request(translated)
	}

}
