package recursos

import (
	"github.com/yaelol1/myp_proyecto1/cliente"
)

type Mensaje struct {
	tipo string
	status string
	roomName string
	message string
	users  map[string]*Cliente
}
