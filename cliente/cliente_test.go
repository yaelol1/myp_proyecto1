package cliente

import (
	// "context"
	// "errors"
	// "fmt"
	// "math/rand"
	// "strconv"
	 "testing"
	// "time"
	// "encoding/json"
)

func TestConectaCliente(t *testing.T){
	c := NuevoCliente()
	c.Conectar()
	go c.lee()
}
