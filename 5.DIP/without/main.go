package main

import "fmt"

// Low-level module
type LightBulb struct{}

func (lb LightBulb) TurnOn() {
	fmt.Println("LightBulb is turned on")
}

func (lb LightBulb) TurnOff() {
	fmt.Println("LightBulb is turned off")
}

type Door struct{}

func (d Door) Open() {
	fmt.Println("Door is open")
}

func (d Door) Close() {
	fmt.Println("Door is closed")
}

// High-level module
type Switch struct {
	lightBulb LightBulb
	Door      Door
}

func (s Switch) Operate(on bool) {
	if on {
		s.lightBulb.TurnOn()
	} else {
		s.lightBulb.TurnOff()
	}
}

func main() {
	lightBulb := LightBulb{}
	sw := Switch{lightBulb: lightBulb}
	sw.Operate(true)
	sw.Operate(false)
}
