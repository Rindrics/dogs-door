package dogdoor

import (
	"time"
)

type DogDoor struct {
	open bool
}

func (dd *DogDoor) Open() {
	dd.open = true
	time.Sleep(3 * time.Second)
	dd.Close()
}

func (dd *DogDoor) Close() {
	dd.open = false
}

func (dd *DogDoor) IsOpen() bool {
	return dd.open
}
