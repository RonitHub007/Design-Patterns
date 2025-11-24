package main

import "fmt"

func main() {
	// ferrari := &SportsCar{
	// 	brand: "Ferrari",
	// 	model: "488 GTB",
	// }

	// var car Car = ferrari

	// car.startEngine()
	// car.shiftGear(1)
	// car.accelerate()
	// car.brake()
	// car.stopEngine()

	/*

		Inheritence

	*/
	car1 := &car{
		brand:        "Volkswagen",
		model:        "Sedan",
		isEngineon:   true,
		currentSpeed: 20,
	}
	manualCar1 := &manualCar{
		car: car1,
	}
	manualCar1.shiftGear()
	car2 := &car{
		brand:        "nexa",
		model:        "Creata",
		isEngineon:   false,
		currentSpeed: 70,
	}

	ev1 := &electricCar{
		car: car2,
	}

	ev1.chargebattery()
	ev1.startEngine()
	manualCar1.startEngine()

	/*

		Polymorphism

	*/
	// declaring a rectangle instance
	rectangle := Rectangle{

		length: 10.5,
		width:  12.25,
	}

	// declaring a square instance
	square := Square{

		side: 15.0,
	}

	// The Figure interface can hold rectangle and square type as they both implements the interface
	var f1 Figure = rectangle
	var f2 Figure = square

	// printing the calculated result
	fmt.Printf("Area of rectangle: %.3f unit sq.\n", f1.Area())
	fmt.Printf("Area of square: %.3f unit sq.\n", f2.Area())
}
