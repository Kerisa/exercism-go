package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot stores data of robot
type Robot struct {
	name string
}

var usedName = map[string]bool{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Name allocates a new name to robot if it's name is empty or return current name 
func (r *Robot)Name() string {
	if len(r.name) != 0 {
		return r.name
	}

	for {
		a1 := 'A' + rand.Int() % 26
		a2 := 'A' + rand.Int() % 26
		n := rand.Int() % 1000
		name := fmt.Sprintf("%c%c%03d", a1, a2, n)
		if !usedName[name] {
			usedName[name] = true
			r.name = name
			break
		}
	}

	return r.name
}

// Reset clear robot's name
func (r *Robot)Reset() {
	r.name = ""
}