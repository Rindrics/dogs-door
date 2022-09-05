package dogdoor

type DogDoor struct {
	open bool
}

func (dd *DogDoor) Open() {
	dd.open = true
}

func (dd *DogDoor) Close() {
	dd.open = false
}

func (dd *DogDoor) IsOpen() bool {
	return dd.open
}
