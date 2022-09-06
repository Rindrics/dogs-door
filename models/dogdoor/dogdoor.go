package dogdoor

import (
	"fmt"
	"time"
)

type DogDoor struct {
	open bool
}

func (dd *DogDoor) Open(done chan bool) {
	fmt.Println("  dog door opens")
	dd.open = true
	time.Sleep(3 * time.Second)
	dd.Close(done)
}

func (dd *DogDoor) Close(done chan bool) {
	fmt.Println("  dog door closes")
	dd.open = false
	done <- true
}

func (dd *DogDoor) IsOpen() bool {
	return dd.open
}
