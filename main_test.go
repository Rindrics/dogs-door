package main

import (
	"fmt"
	"testing"

	"github.com/Rindrics/dogs-door/models/bark_recognizer"
	"github.com/Rindrics/dogs-door/models/dogdoor"
)

func TestMain(t *testing.T) {
	door := &dogdoor.DogDoor{}
	recognizer := bark_recognizer.BarkRecognizer{Door: door}

	fmt.Println("Fido starts to bark:")
	newBark := "woof"
	recognizer.Recognize(newBark)

	fmt.Println("Fido went out")

	if door.IsOpen() {
		t.Error("door still opened")
	}

	fmt.Println("Fido finished doing something")

	fmt.Println("Fido starts to bark:")
	recognizer.Recognize(newBark)

	fmt.Println("Fido got back to the inside")
}
