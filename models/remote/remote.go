package remote

import "github.com/Rindrics/dogs-door/models/dogdoor"

type Remote struct {
	door dogdoor.DogDoor
}

func (r *Remote) PressButton() {
	if r.door.IsOpen() {
		r.door.Close()
	} else {
		r.door.Open()
	}
}
