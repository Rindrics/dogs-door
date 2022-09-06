package bark_recognizer

import (
	"fmt"

	"github.com/Rindrics/dogs-door/models/dogdoor"
)

type BarkRecognizer struct {
	door dogdoor.DogDoor
}

func (br *BarkRecognizer) Recognize(bark string) {
	fmt.Println("BarkRecognizer: detect ->'", bark, "'")
	br.door.Open()
}
