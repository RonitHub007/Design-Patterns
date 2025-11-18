package main

func main() {
	ferrari := &SportsCar{
		brand: "Ferrari",
		model: "488 GTB",
	}

	var car Car = ferrari

	car.startEngine()
	car.shiftGear(1)
	car.accelerate()
	car.brake()
	car.stopEngine()
}
