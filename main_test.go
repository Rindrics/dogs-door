package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/Rindrics/dogs-door/models/bark_recognizer"
	"github.com/Rindrics/dogs-door/models/dogdoor"
)

func TestMain(t *testing.T) {
	door := &dogdoor.DogDoor{}
	recognizer := bark_recognizer.BarkRecognizer{Door: door}

	done := make(chan bool)
	fmt.Println("Fido starts to bark:")
	newBark := "woof"
	recognizer.Recognize(newBark, done)
	time.Sleep(1 * time.Second)
	fmt.Println("Fido went out")
	<-done

	fmt.Println("Fido finished doing something")

	fmt.Println("Fido starts to bark:")
	recognizer.Recognize(newBark, done)
	time.Sleep(1 * time.Second)
	fmt.Println("Fido got back to the inside")
	<-done
}
