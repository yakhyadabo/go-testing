package interfacez

import "fmt"


// access method of the interface
func Calculate(shape Shape) {
	fmt.Println("Area:", shape.area())
	fmt.Println("Perimeter:", shape.perimeter())
}

//// main function
//func main() {
//
//	// assigns value to struct members
//	rect := Rectangle{7, 4}
//	circle := Circle{radius: 4}
//
//	// call calculate() with struct variable rect
//	calculate(rect)
//	calculate(circle)
//}