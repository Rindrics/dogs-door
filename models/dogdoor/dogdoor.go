package dogdoor

import (
	"fmt"
	"time"
)

type DogDoor struct {
	open bool
}

func (dd *DogDoor) Open() {
	fmt.Println("dog door opens")
	dd.open = true
	time.Sleep(3 * time.Second)
	dd.Close()
}

func (dd *DogDoor) Close() {
	fmt.Println("dog door closes")
	dd.open = false
}

func (dd *DogDoor) IsOpen() bool {
	return dd.open
}
