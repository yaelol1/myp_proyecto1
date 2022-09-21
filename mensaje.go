package myp_proyecto1

type Mensaje struct {
	tipo string
	status string
	roomname string
	message string
	users  map[string]*Cliente
}
