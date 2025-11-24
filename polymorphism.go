// Golang program to illustrate the
// concept of interfaces

package main

// defining an interface
type Figure interface {
	Area() float64
}

// declaring a struct
type Rectangle struct {

	// declaring struct variables
	length float64
	width  float64
}

// declaring a struct
type Square struct {

	// declaring struct variable
	side float64
}

// function to calculate
// area of a rectangle
func (rect Rectangle) Area() float64 {

	// Area of rectangle = l * b
	area := rect.length * rect.width
	return area
}

// function to calculate
// area of a square
func (sq Square) Area() float64 {

	// Area of square = a * a
	area := sq.side * sq.side
	return area
}
