package myp_proyecto1

import (
	"encoding/json"
)

type cliente struct {
	nombre   string
	cuartos  map[string]*cuarto
}

// TODO: net.Dial -> connection -> Write
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
	// handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
// ...

