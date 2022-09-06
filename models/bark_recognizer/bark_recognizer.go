package bark_recognizer

import (
	"fmt"

	"github.com/Rindrics/dogs-door/models/dogdoor"
)

type BarkRecognizer struct {
	Door *dogdoor.DogDoor
}

func (br *BarkRecognizer) Recognize(bark string) {
	if br.Door.GetAllowedBark() == bark {
		fmt.Println("BarkRecognizer: detect ->'", bark, "'")
		br.Door.Open()
	} else {
		fmt.Println("the bark is not allowed")
	}
}
