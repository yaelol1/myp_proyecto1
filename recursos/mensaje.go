package recursos

import (
)

type Mensaje struct {
	tipo string
	status string
	roomName string
	message string
	users  map[string]*Cliente
}
