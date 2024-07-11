package main

import "fmt"

// Worker Large interface
type Worker interface {
	Work()
	Eat()
}

type HumanWorker struct{}

func (hw HumanWorker) Work() {
	fmt.Println("Human is working")
}

func (hw HumanWorker) Eat() {
	fmt.Println("Human is eating")
}

func NewHumanWorker() Worker {
	return HumanWorker{}
}

type RobotWorker struct{}

func (rw RobotWorker) Work() {
	fmt.Println("Robot is working")
}

// RobotWorker must implement Eat() even though it doesn't eat
func (rw RobotWorker) Eat() {
	// Not needed for RobotWorker, but must be implemented
}

func NewRobotWorker() Worker {
	return RobotWorker{}
}

func main() {
	human := NewHumanWorker()
	robot := NewRobotWorker()

	human.Work()
	human.Eat()
	robot.Work()
	robot.Eat()
}
