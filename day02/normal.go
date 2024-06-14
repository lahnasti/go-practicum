package main

import "fmt"

type Vihicle interface {
	Start()
	Stop()
}

type Car struct {
	Model string
}

type Motocycle struct {
	Model string
}

func (c Car) Start() string {
	return fmt.Sprintf("%s IS STARTING", c.Model)
}

func (c Car) Stop() string {
	return fmt.Sprintf("%s IS STOPPING", c.Model)
}

func (m Motocycle) Start() string {
	return fmt.Sprintf("%s IS STARTING", m.Model)
}

func (m Motocycle) Stop() string {
	return fmt.Sprintf("%s IS STOPPING", m.Model)
}

func main () {
	firstCar := Car {
		"Porshe",
	}
	firstMotocycle := Motocycle {
		"Yamaha",
	}
	fmt.Println(firstCar.Start())
	fmt.Println(firstCar.Stop())
	fmt.Println(firstMotocycle.Start())
	fmt.Println(firstMotocycle.Stop())
}