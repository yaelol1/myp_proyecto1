package myp_proyecto1

import(
	"fmt"
	"encoding/json"
)

type servidor struct {
	rooms    map[string]*room
	commands chan command
}
