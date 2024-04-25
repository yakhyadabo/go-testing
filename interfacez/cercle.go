package interfacez

import "math"

type Circle struct { // struct to implement interface
	radius float32 // 2πr
}

func (c Circle) area() float32 { // use struct to implement area() of interface
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float32 { // use struct to implement area() of interface
	return 2 * math.Pi * c.radius
}

func MakeCircle(radius float32) Circle { // Return Value
	return Circle{
		radius,
	}
}