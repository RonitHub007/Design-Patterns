package main

import "fmt"

type car struct {
	brand        string
	model        string
	isEngineon   bool
	currentSpeed float64
}
type carMethods interface {
	startEngine()
	stopEngine()
	accelarate()
	breakApply()
}

type manualCar struct {
	*car
}
type electricCar struct {
	*car
}

type manualCarMethods interface {
	carMethods
	shiftGear()
}
type electricCarMethods interface {
	carMethods
	chargebattery()
}

func (c *car) startEngine() {
	if c.isEngineon {
		fmt.Println("Engine is on , start The car")
	} else {
		fmt.Println("Engine is off, Start is not posiible")
	}
}
func (ev *electricCar) chargebattery() {
	fmt.Println("Charge your battery")
}

func (ev *manualCar) shiftGear() {
	fmt.Println("Gear shift activated")
}
