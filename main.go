package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	fmt.Println("Hello")
	rpio.Open()
	defer rpio.Close()
	pin := rpio.Pin(6)
	//4 22 6 26
	pin.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
	for ok := true; ok; ok = true {
		//Maybe something here to respond to keystrokes?
	}
}
