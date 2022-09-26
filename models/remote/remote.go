package remote

import (
	"fmt"

	"github.com/Rindrics/dogs-door/models/dogdoor"
)

type Remote struct {
	door dogdoor.DogDoor
}

func (r *Remote) PressButton(done chan bool) {
	fmt.Println("remote controller button is pressed")
	if r.door.IsOpen() {
		r.door.Close(done)
	} else {
		r.door.Open(done)
	}
}
