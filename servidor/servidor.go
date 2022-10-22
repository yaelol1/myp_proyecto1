// Package main contiene todo el servidor que recibirá conexiones,
// las cuales serán tratados como usuarios de un chat de acuerdo al protocolo.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// Un User contiene todos los atributos de una conexión
type User struct {
	userName string
	status string
	conn net.Conn
}

// Un Servidor contiene los cuartos y los usuarios en él.
type Servidor struct {
	cuartos map[string]*Cuarto
	general *Cuarto
	users   map[string]*User
}

// NuevoServidor crea un servidor y devuelve su apuntador
func NuevoServidor() *Servidor {
	return &Servidor{
		cuartos: make(map[string]*Cuarto),
		users:   make(map[string]*User),
		general: NuevoCuarto("General"),
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

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		// handle input from connection
		go s.handleConnection(conn)
	}
}


// handleConnection acepta las conexiones y decide qué hacer con ellas.
func (s *Servidor) handleConnection(conn net.Conn) {
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
	fmt.Println("Request: ", msg)

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
	fmt.Println("Sending: ", msg)

	d := json.NewEncoder(conn)
	if err := d.Encode(msg); err != nil {
		fmt.Println(err)
	}
}



// user devuelve al usuario a partir de su conexión o nombre.
func (s *Servidor) user(userAsked interface{}) *User{

	conn, isConn := userAsked.(net.Conn)
	if isConn {
		return s.general.User(conn)
	}

	name, isName := userAsked.(string)
	if isName {
		return s.users[name]
	}

	panic("Not a name or connection was passed")
}

// idetify identifica a un usuario en todo el servidor, además de cambiar su
// nombre en general.
func (s *Servidor) identify(conn net.Conn, msg map[string]interface{}) {

	userName := msg["username"].(string)
	_, nameTaken := s.users[userName]
	if !nameTaken {
		newUser := User{
			userName: userName,
			status: "away",
			conn: conn,
		}
		s.users[userName] = &newUser
		s.general.invitaIntegrante(userName)
		s.general.agregaIntegrante(conn, &newUser)

		response := map[string]interface{}{"type": "NEW_USER",
			"username": userName}

		s.general.Broadcast(conn, response)
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
	user := s.user(conn)
	userName := user.userName
	newStatus := msg["status"].(string)
	if user.status == newStatus {

		selfResponse := map[string]interface{}{"type": "WARNING",
			"message": "El estado ya es '"+newStatus+"'",
			"operation": "STATUS",
			"status": newStatus }

		s.send(conn, selfResponse)
		return
	}

	// respuesta a usuarios
	response := map[string]interface{}{"type": "NEW_STATUS",
		"username": userName,
		"status": newStatus}
	s.general.Broadcast(conn, response)

	// respuesta info a conexión
	response = map[string]interface{}{"type": "INFO",
		"message": "success",
		"operation": "STATUS"}

	s.send(conn, response)
}

// usersList envía una lista de todos los usuarios en el servidor.
func (s *Servidor) usersList(conn net.Conn, msg map[string]interface{}) {
	var users []string
	users = s.general.userList()
	selfResponse := map[string]interface{}{"type": "USER_LIST",
		"usernames": users}
	s.send(conn, selfResponse)
}

// message envía un mensaje de usuario a usuario.
func (s *Servidor) message(conn net.Conn, msg map[string]interface{}) {
	user := msg["username"] // Checking for existing key and its value
	connRecipient, existe := s.users[user.(string)]
	if !existe {
		selfResponse := map[string]interface{}{"type": "WARNING",
			"message":   "El usuario '" + user.(string) + "' no existe",
			"operation": "MESSAGE",
			"username":  user.(string)}
		s.send(conn, selfResponse)
		return
	}

	userName := s.user(conn).userName
	Response := map[string]interface{}{"type": "MESSAGE_FROM",
		"username": userName,
		"message":   msg["message"]}
	s.send(connRecipient.conn, Response)

}

// publicMessage envía un mensaje a todos los usuarios conectados.
func (s *Servidor) publicMessage(conn net.Conn, msg map[string]interface{}) {
	userName := s.user(conn).userName
	response := map[string]interface{}{"type": "PUBLIC_MESSAGE_FROM",
		"username": userName,
		"message":  msg["message"]}
	s.general.Broadcast(conn, response)

	selfResponse := map[string]interface{}{"type": "INFO",
		"message":   "success",
		"operation": "IDENTIFY"}
	s.send(conn, selfResponse)
}

// newRoom crea un nuevo cuarto en el servidor.
func (s *Servidor) newRoom(conn net.Conn, msg map[string]interface{}) {
	user := s.user(conn)
	roomname := msg["roomname"].(string)
	_, ok := s.cuartos[roomname]
	if !ok {
		s.cuartos[roomname] = NuevoCuarto(roomname)
		s.cuartos[roomname].agregaIntegrante(conn, user)

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
	roomname := msg["roomname"].(string)
	integrantesRaw := msg["usernames"]
	integrantesRaw2 := integrantesRaw.([]interface{})

	room, existsRoom := s.cuartos[roomname]
	if !existsRoom {
		selfResponse := map[string]interface{}{"type": "WARNING",
			"message": "La sala '"+roomname+"' no existe",
			"roomname": roomname,
			"operation": "INVITE"}
		s.send(conn, selfResponse)
		return
	}

	// la sala existe, entonces se mandan las invitaciones
	sender := s.user(conn).userName
	toSend := map[string]interface{}{"type": "INVITATION",
		"message":  sender + " te invitó al cuarto '" + roomname + "'",
		"username": sender,
		"roomname": roomname}

	for _, user := range integrantesRaw2 {
		fmt.Print("usernames: " + user.(string))
		userToInvite, existsUserToI := s.users[user.(string)]
		if !existsUserToI {
			selfResponse := map[string]interface{}{"type": "WARNING",
				"message": "El usuario '"+user.(string)+"' no existe",
				"username": user.(string),
				"operation": "INVITE"}
			s.send(conn, selfResponse)
			continue
		}
		s.send(userToInvite.conn, toSend)
		room.invitaIntegrante(userToInvite.userName)
	}

	// info success, se mandaron las invitaciones
	selfResponse := map[string]interface{}{"type": "INFO",
		"message": "success",
		"roomname": roomname,
		"operation": "INVITE"}
	s.send(conn, selfResponse)
}

// joinRoom une al usuario a un servidor si es que lo invitaron al cuarto.
func (s *Servidor) joinRoom(conn net.Conn, msg map[string]interface{}) {
	user := s.user(conn)
	username := user.userName
	roomname := msg["roomname"].(string)

	msgErr, esValido := s.validaCuarto(conn, msg)
	if !esValido {
		msgErr["operation"]= "JOIN_ROOM"
		s.send(conn, msgErr)
		return
	}

	roomToJoin := s.cuartos[roomname]
	// verifica que el usuario tenga una invitación
	if !(roomToJoin.fueInvitado(username)){
		selfResponse := map[string]interface{}{"type": "WARNING",
			"message": "El usuario no ha sido invitado al cuarto '"+roomname+"'",
			"operation": "JOIN_ROOM",
			"roomname": roomname }
		s.send(conn, selfResponse)
		return
	}

	roomToJoin.agregaIntegrante(conn, user)

	// mensaje a los demás usuarios
	response := map[string]interface{}{"type": "JOINED_ROOM",
		"username": username,
		"roomname": roomname }
	roomToJoin.Broadcast(conn, response)

	// mensaje de éxito
	selfResponse := map[string]interface{}{"type": "INFO",
		"message": "success",
		"operation": "JOIN_ROOM",
		"roomname": roomname }
	s.send(conn, selfResponse)
}

// roomUsers manda una lista de los usuarios dentro de un cuarto.
func (s *Servidor) roomUsers(conn net.Conn, msg map[string]interface{}) {
	roomname := msg["roomname"].(string)

	msgErr, esValido := s.validaCuarto(conn, msg)
	if !esValido {
		msgErr["operation"]= "ROOM_USERS"
		s.send(conn, msgErr)
		return
	}

	cuarto := s.cuartos[roomname]

	userList := cuarto.userList()
	selfResponse := map[string]interface{}{"type": "ROOM_USER_LIST",
		"usernames": userList}
	s.send(conn, selfResponse)
}

// roomMessage envía un mensaje dentro de un cuarto.
func (s *Servidor) roomMessage(conn net.Conn, msg map[string]interface{}) {
	roomname := msg["roomname"].(string)

	msgErr, esValido := s.validaCuarto(conn, msg)
	if !esValido {
		msgErr["operation"]= "ROOM_MESSAGE"
		s.send(conn, msgErr)
		return
	}

	r := s.cuartos[roomname]
	userName := s.user(conn).userName
	response := map[string]interface{}{"type": "ROOM_MESSAGE_FROM",
		"roomname": roomname,
		"username": userName,
		"message":  msg["message"]}
	r.Broadcast(conn, response)

	selfResponse := map[string]interface{}{"type": "INFO",
		"message":   "success",
		"operation": "ROOM_MESSAGE"}
	s.send(conn, selfResponse)
}

// leaveRoom sale de un cuarto.
func (s *Servidor) leaveRoom(conn net.Conn, msg map[string]interface{}) {
	roomname := msg["roomname"].(string)
	msgErr, esValido := s.validaCuarto(conn, msg)
	if !esValido {
		msgErr["operation"]= "LEAVE_ROOM"
		s.send(conn, msgErr)
		return
	}

	room  := s.cuartos[roomname]
	userName := s.user(conn).userName
	response := map[string]interface{}{"type": "LEAVE_ROOM",
		"roomname": roomname,
		"username": userName }
	room.Broadcast(conn, response)

	selfResponse := map[string]interface{}{"type": "INFO",
		"message":   "success",
		"roomname": roomname,
		"operation": "LEAVE_ROOM"}
	s.send(conn, selfResponse)

	room.eliminaIntegrante(conn)
}

// disconnect desconecta al usuario de todos los cuartos, incluido el principal.
func (s *Servidor) disconnect(conn net.Conn, msg map[string]interface{}) {
	userName := s.user(conn).userName

	// respuesta a usuarios
	response := map[string]interface{}{"type": "DISCONNECTED",
		"username": userName }

	s.general.Broadcast(conn, response)

	for roomname, room := range s.cuartos {
		if room.esIntegrante(conn) {
			room.eliminaIntegrante(conn)

			response := map[string]interface{}{"type": "LEFT_ROOM",
				"roomname": roomname,
				"username": userName }
			room.Broadcast(conn, response)
		}
	}
}

// validaCuarto valida que el cuarto exista y el usuario esté en el mismo.
func (s *Servidor) validaCuarto(conn net.Conn ,msg map[string]interface{}) (map[string]interface{}, bool) {
	roomname := msg["roomname"].(string)
	room, ok := s.cuartos[roomname]
	// se asegura que el cuarto exista
	if !ok {
		selfResponse := map[string]interface{}{"type": "WARNING",
			"message": "La sala '"+roomname+"' no existe",
			"roomname": roomname }
		return selfResponse, false
	}

	if !(room.esIntegrante(conn)){
		selfResponse := map[string]interface{}{"type": "WARNING",
			"message": "El usuario no se ha unido a '"+
				roomname+"'",
			"roomname": roomname }
		return selfResponse, false
	}
	return make(map[string]interface{}), true
}

// info
func (s *Servidor) info(conn net.Conn, msg map[string]interface{}) {
}

// warning
func (s *Servidor) warning(conn net.Conn, msg map[string]interface{}) {
}
