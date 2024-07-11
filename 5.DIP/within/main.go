package main

import "fmt"

// Switchable Abstraction
type Switchable interface {
	TurnOn()
	TurnOff()
}

// LightBulb Low-level module
type LightBulb struct{}

func (lb LightBulb) TurnOn() {
	fmt.Println("LightBulb is turned on")
}

func (lb LightBulb) TurnOff() {
	fmt.Println("LightBulb is turned off")
}

// Switch High-level module
type Switch struct {
	device Switchable
}

func (s Switch) Operate(on bool) {
	if on {
		s.device.TurnOn()
	} else {
		s.device.TurnOff()
	}
}

func main() {
	lightBulb := LightBulb{}
	sw := Switch{device: lightBulb}
	sw.Operate(true)
	sw.Operate(false)
}
