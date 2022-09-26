package bark_recognizer

import (
	"fmt"

	"github.com/Rindrics/dogs-door/models/dogdoor"
)

type BarkRecognizer struct {
	Door *dogdoor.DogDoor
}

func (br *BarkRecognizer) Recognize(bark string, done chan bool) {
	fmt.Println("  BarkRecognizer: detect ->'", bark, "'")
	go br.Door.Open(done)
}
