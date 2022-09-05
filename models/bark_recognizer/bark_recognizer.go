package bark_recognizer

import "github.com/Rindrics/dogs-door/models/dogdoor"

type BarkRecognizer struct {
	door dogdoor.DogDoor
}

func (br *BarkRecognizer) Recognize(bark string) {
	// recognize given bark
}
