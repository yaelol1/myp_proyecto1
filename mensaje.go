package myp_proyecto1

type Message struct {
	tipo string
	status string
	roomname string
	message string
	users  map[string]*cliente
}
