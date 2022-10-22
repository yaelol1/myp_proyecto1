package main

 import (
	"bufio"
	"os"
	"time"
	"os/exec"
	"strings"
	"fmt"
 )

// instrucciones imprime las instrucciones de las funciones
// del chat.
func instrucciones(){
	fmt.Printf("/msgAll mensaje Para mandar mensaje publico  \n")
	fmt.Printf("/msg cuarto mensaje Para mandar mensaje a un cuarto \n")
	fmt.Printf("/whisper usuario mensaje  Para mandar mensaje privado \n")
	fmt.Printf("/room nombreCuarto ... Para crear un cuarto \n")
	fmt.Printf("/invite nombreCuarto usuario1 usuario2 ... Para invitar usuarios a un cuarto \n")
	fmt.Printf("/accept nombreCuarto Para acepar una invitación a un cuarto  \n")
	fmt.Printf("/list nombreCuarto Para pedir una lista de todos los usuarios, si no hay nombre de cuarto se listan todos los usuarios \n")
	fmt.Printf("/leave nombreCuarto Para abandonar un cuarto \n")
	fmt.Printf("/disconnect Para cerrar la aplicación \n")
	fmt.Printf("/clear Para limpiar la pantalla \n")
	fmt.Printf("/help Para imprimir las instrucciones otra vez \n")
}

// actionTranslator Toma el input del usuario y lo manda a
// request al servidor
func actionTranslator(action string) interface{} {
	// Si hay un error continua
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("\n Comando o formato no válido, para imprimir la lista de comandos escriba /help \n")
		}
	}()

	actionArr := strings.SplitN(action, " ",2)


	switch actionArr[0]{
		case "/msgAll":
		r := map[string]interface{}{"type": "PUBLIC_MESSAGE", "message": actionArr[1]}
		return r

		case "/whisper":
		recipAndMsg := strings.SplitN(actionArr[1], " ",2)
		r := map[string]interface{}{"type": "MESSAGE", "username": recipAndMsg[0],"message": recipAndMsg[1]}
		return r

		case "/msg":
		recipAndMsg := strings.SplitN(actionArr[1], " ",2)
		r := map[string]interface{}{"type": "ROOM_MESSAGE", "roomname": recipAndMsg[0],"message": recipAndMsg[1]}
		return r

		case "/room":
		r := map[string]interface{}{"type": "NEW_ROOM", "roomname": actionArr[1]}
		return r

		case "/invite":
		recipAndMsg := strings.SplitN(actionArr[1], " ",2)
		users := strings.Fields(recipAndMsg[1])
		r := map[string]interface{}{"type": "INVITE", "roomname": recipAndMsg[0], "usernames": users}
		return r

		case "/accept":
		r := map[string]interface{}{"type": "JOIN_ROOM", "roomname": actionArr[1]}
		return r

		case "/list":
		if len(actionArr) < 2 {
			r := map[string]interface{}{"type": "USERS"}
			return r
		}
		r := map[string]interface{}{"type": "ROOM_USERS", "roomname": actionArr[1]}
		return r

		case "/leave":
		r := map[string]interface{}{"type": "LEAVE_ROOM", "roomname": actionArr[1]}
		return r

		case "/help":
		instrucciones()

		case "/clear":
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()

		case "/disconnect":
		r := map[string]interface{}{"type": "DISCONNECT"}
		go exit()
		return r

		default:
		fmt.Printf(actionArr[0])
		fmt.Printf("\n Comando no válido, para imprimir la lista de comandos escriba /info \n")
	}

	return nil
}

func exit(){

	DurationOfTime := time.Duration(1) * time.Second
	f := func() {
		os.Exit(0)
	}
	Timer1 := time.AfterFunc(DurationOfTime, f)
	defer Timer1.Stop()
	time.Sleep(1 * time.Second)
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
	c.nombre = action
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
