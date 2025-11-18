package main

import "fmt"

type Car interface {
	startEngine()
	shiftGear(gear int64)
	accelerate()
	brake()
	stopEngine()
}

type SportsCar struct {
	brand        string
	model        string
	isEngineOn   bool
	currentSpeed int64
	currentGear  int64
}

func (sc *SportsCar) startEngine() {
	sc.isEngineOn = true
	fmt.Printf("Engine starts with a roar, brand: %s", sc.brand)
}
func (sc *SportsCar) shiftGear(gear int64) {
	if !sc.isEngineOn {
		fmt.Println("Gear Shift not happening as Engine is off")
	}
	sc.currentGear = gear
	fmt.Printf("Gear is shifted by %d", gear)
	fmt.Printf("gear is now %d", sc.currentGear)
}
func (sc *SportsCar) accelerate() {
	if !sc.isEngineOn {
		fmt.Println("Accelaretion not happening as Engine is off")
	}
	sc.currentSpeed += 20
	fmt.Printf("spped is now %d", sc.currentSpeed)
}
func (sc *SportsCar) brake() {
	if !sc.isEngineOn {
		fmt.Println("Break not happening as Engine is off")
	}
	sc.currentSpeed -= 20
	fmt.Printf("spped is now %d", sc.currentSpeed)
}
func (sc *SportsCar) stopEngine() {
	sc.isEngineOn = false
	fmt.Printf("Engine stops with a sound, brand: %s", sc.brand)
}
