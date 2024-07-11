package main

import "fmt"

// Segregated interfaces
type Worker interface {
	Work()
}

type Eater interface {
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

func NewRobotWorker() Worker {
	return RobotWorker{}
}

func Work(worker Worker) {
	worker.Work()
}

func Eat(worker Eater) {
	worker.Eat()
}

func main() {
	human := HumanWorker{}
	robot := RobotWorker{}

	Work(human)
	Work(robot)

	Eat(human)

}
