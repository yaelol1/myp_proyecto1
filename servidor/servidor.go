package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)


// Un User contiene todos los atributos de una conexión
type User struct {
	rooms map[string]*Cuarto
	usersName string
	conn net.Conn
}

// Un Servidor contiene los cuartos y los usuarios en él.
type Servidor struct {
	cuartos map[string]*Cuarto
	users   map[string]*User
}

// NuevoServidor crea un servidor y devuelve su apuntador
func NuevoServidor() *Servidor {
	return &Servidor{
		cuartos: make(map[string]*Cuarto),
		users:   make(map[string]net.Conn),
	}

}

// IncicializaServidor inicia el servidor y crea el cuarto general,
// mete a cada conexión al cuarto principal.
func (s *Servidor) InicializaServidor() {
	fmt.Print("Servidor escuchando \n")
	ln, err := net.Listen("tcp", ":3306")
	if err != nil {
		// handle error
	}

	general := NuevoCuarto("general")
	s.cuartos["General"] = general
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		// add accepted connection to the general room
		general.agregaIntegrante(conn, "")

		// handle input from connection
		go s.handleConnection(conn)
	}
}

// handleConnection acepta las conexiones y decide qué hacer con ellas.
func (s *Servidor) handleConnection(conn net.Conn) {
	// prueba raw
	// b:=make([]byte, 100)
	// bs, errb := conn.Read(b)
	// fmt.Println("Mensaje:", string(b[:bs]), errb)

	// Si hay en error solo cierra la conexión
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	// Decodificador que lee directamente desde el socket
	decoder := json.NewDecoder(conn)

	// Interfaz, al no saber qué datos tendrá el JSON
	var jsonData interface{}
	for {

		err := decoder.Decode(&jsonData)

		if err != nil {
			fmt.Println(jsonData, err)
			return
			// handle error
		}
		// Se convierte a un mapa
		msg := jsonData.(map[string]interface{})
		s.Response(msg, conn)
	}
}

// Response acepta las respuestas de los clientes.
func (s *Servidor) Response(msg map[string]interface{}, conn net.Conn) {
	fmt.Print("\n Request: ", msg)

	// en caso de error solo termina la función
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	tipo := msg["type"] // Checking for existing key and its value

	switch tipo {

	case "ERROR":
	case "WARNING":
	case "INFO":
	case "IDENTIFY":
		s.identify(conn, msg)
	case "STATUS":
		s.status(conn, msg)
	case "USERS":
		s.usersList(conn, msg)
	case "MESSAGE":
		s.message(conn, msg)
	case "PUBLIC_MESSAGE":
		s.publicMessage(conn, msg)
	case "NEW_ROOM":
		s.newRoom(conn, msg)
	case "INVITE":
		s.invite(conn, msg)
	case "JOIN_ROOM":
		s.joinRoom(conn, msg)
	case "ROOM_USERS":
		s.roomUsers(conn, msg)
	case "ROOM_MESSAGE":
		s.roomMessage(conn, msg)
	case "LEAVE_ROOM":
		s.leaveRoom(conn, msg)
	case "DISCONNECT":
		s.disconnect(conn, msg)
	default:
		fmt.Print("invalid", msg)

	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
}

// send envía un mensaje a una sola conexión.
func (s *Servidor) send(conn net.Conn, msg map[string]interface{}) {
	d := json.NewEncoder(conn)
	if err := d.Encode(msg); err != nil {
		fmt.Println(err)
	}
}

// validaEntrada valida la entrada de datos.
// TODO: Necesario?
func (s *Servidor) validaEntrada(entrada string, msg net.Conn) string {
	panic("")
}

// userConn devuelve la conexión de un cliente a partir de su nombre.
func (s *Servidor) userConn(name string) net.Conn {
	return s.users[name]
}

// userName devuelve el nombre de un usuario a partir de su conexión.
func (s *Servidor) userName(conn net.Conn) string {
	return s.cuartos["General"].obtenNombre(conn)
}

// idetify identifica a un usuario en todo el servidor, además de cambiar su
// nombre en general.
func (s *Servidor) identify(conn net.Conn, msg map[string]interface{}) {
	username, ok1 := msg["username"] // Checking for existing key and its value
	if !ok1 {
		return
	}

	userName := username.(string)
	_, nameTaken := s.users[userName]
	if !nameTaken {
		s.users[userName] = conn
		s.cuartos["General"].agregaIntegrante(conn, userName)

		response := map[string]interface{}{"type": "NEW_USER",
			"username": userName}

		s.cuartos["General"].Broadcast(conn, response)
		return
	}

	selfResponse := map[string]interface{}{"type": "WARNING",
		"message": "El usuario '"+ userName+"' ya existe",
		"operation": "IDENTIFY",
		"username": userName}
	s.send(conn, selfResponse)

}

// status cambia el status del usuario.
func (s *Servidor) status(conn net.Conn, msg map[string]interface{}) {
	userName := s.userName(conn)
	newStatus := msg["status"]
	response := map[string]interface{}{"type": "NEW_STATUS",
		"username": userName,
		"status": newStatus}

	s.cuartos["General"].Broadcast(conn, response)
}

// usersList envía una lista de todos los usuarios en el servidor.
func (s *Servidor) usersList(conn net.Conn, msg map[string]interface{}) {
	var users []string
	users = s.cuartos["General"].userList()
	selfResponse := map[string]interface{}{"type": "USER_LIST",
		"usernames": users}
	s.send(conn, selfResponse)
}

// message envía un mensaje de usuario a usuario.
func (s *Servidor) message(conn net.Conn, msg map[string]interface{}) {
	user, ok := msg["username"] // Checking for existing key and its value
	if !ok {
		return
	}
	connRecipient, existe := s.users[user.(string)]
	if !existe {
		selfResponse := map[string]interface{}{"type": "WARNING",
			"message":   "El usuario '" + user.(string) + "' no existe",
			"operation": "MESSAGE",
			"username":  user.(string)}
		s.send(conn, selfResponse)
		return
	}

	userName := s.userName(conn)
	Response := map[string]interface{}{"type": "MESSAGE",
		"usernames": userName,
		"message":   msg["message"]}
	s.send(connRecipient, Response)

}

// publicMessage envía un mensaje a todos los usuarios conectados.
func (s *Servidor) publicMessage(conn net.Conn, msg map[string]interface{}) {
	userName := s.userName(conn)
	response := map[string]interface{}{"type": "PUBLIC_MESSAGE_FROM",
		"username": userName,
		"message":  msg["message"]}
	s.cuartos["General"].Broadcast(conn, response)

	selfResponse := map[string]interface{}{"type": "INFO",
		"message":   "success",
		"operation": "IDENTIFY"}
	s.send(conn, selfResponse)
}

// newRoom crea un nuevo cuarto en el servidor.
func (s *Servidor) newRoom(conn net.Conn, msg map[string]interface{}) {
	userName := s.userName(conn)
	roomname := msg["roomname"].(string)
	_, ok := s.cuartos[roomname]
	if !ok {
		s.cuartos[roomname] = NuevoCuarto(roomname)
		s.cuartos[roomname].agregaIntegrante(conn, userName)

		selfResponse := map[string]interface{}{"type": "INFO",
			"message":   "success",
			"operation": "NEW_ROOM"}
		s.send(conn, selfResponse)
		return
	}

	selfResponse := map[string]interface{}{"type": "WARNING",
		"message":   "El cuarto '" + roomname + "' ya existe",
		"operation": "NEW_ROOM"}
	s.send(conn, selfResponse)
}

// invite invita uno o más usuarios a un cuarto del servidor.
func (s *Servidor) invite(conn net.Conn, msg map[string]interface{}) {
	roomname, okR := msg["roomname"].(string)
	if !okR {
		fmt.Print("invalid roomname")
		return
	}
	integrantesRaw, okI := msg["usernames"]
	if !okI {
		fmt.Print("invalid usernames")
		return
	}
	integrantesRaw2 := integrantesRaw.([]interface{})
	room := s.cuartos[roomname]
	sender := s.userName(conn)
	toSend := map[string]interface{}{"type": "INVITATION",
		"message":  sender + " te invitó al cuarto '" + roomname + "'",
		"username": sender,
		"roomname": roomname}

	for _, user := range integrantesRaw2 {
		fmt.Print("usernames: " + user.(string))
		conn := s.users[user.(string)]
		s.send(conn, toSend)
		room.agregaIntegrante(conn, conn.RemoteAddr().String()+
			" "+user.(string))
	}
}

// joinRoom une al usuario a un servidor si es que lo invitaron al cuarto.
func (s *Servidor) joinRoom(conn net.Conn, msg map[string]interface{}) {
	roomName := msg["roomname"].(string)
	roomToJoin := s.cuartos[roomName]
	str := roomToJoin.obtenNombre(conn)
	if str == "" {
		return
	}
	if strings.HasPrefix(str, conn.RemoteAddr().String()+" ") {
		roomToJoin.agregaIntegrante(conn,
			strings.TrimPrefix(str, conn.RemoteAddr().String()+" "))
		return
	}
}

// roomUsers manda una lista de los usuarios dentro de un cuarto.
func (s *Servidor) roomUsers(conn net.Conn, msg map[string]interface{}) {
	roomname := msg["roomname"].(string)
	cuarto := s.cuartos[roomname]
	userList := cuarto.userList
	selfResponse := map[string]interface{}{"type": "ROOM_USER_LIST",
		"usernames": userList}
	s.send(conn, selfResponse)
}

// roomMessage envía un mensaje dentro de un cuarto.
func (s *Servidor) roomMessage(conn net.Conn, msg map[string]interface{}) {
	nombreCuarto := msg["roomname"].(string)
	r, ok := s.cuartos[nombreCuarto]
	if !ok {
		return
	}

	userName := s.userName(conn)
	response := map[string]interface{}{"type": "ROOM_MESSAGE_FROM",
		"roomname": nombreCuarto,
		"username": userName,
		"message":  msg["message"]}
	r.Broadcast(conn, response)

	selfResponse := map[string]interface{}{"type": "INFO",
		"message":   "success",
		"operation": "ROOM_MESSAGE"}
	s.send(conn, selfResponse)
}

// leaveRoom
func (s *Servidor) leaveRoom(conn net.Conn, msg map[string]interface{}) {
	roomname := msg["roomname"].(string)
	room := s.cuartos[roomname]

	room.eliminaIntegrante(conn)
}

// disconnect
func (s *Servidor) disconnect(conn net.Conn, msg map[string]interface{}) {

	for _, room := range s.cuartos {
		room.eliminaIntegrante(conn)
	}
}

// info
func (s *Servidor) info(conn net.Conn, msg map[string]interface{}) {
}

// warning
func (s *Servidor) warning(conn net.Conn, msg map[string]interface{}) {
}
