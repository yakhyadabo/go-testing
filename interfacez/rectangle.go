package interfacez

type Rectangle struct { // struct to implement interface
	length, breadth float32
}

func (r Rectangle) area() float32 { // use struct to implement area() of interface
	return r.length * r.breadth
}

func (r Rectangle) perimeter() float32 { // use struct to implement area() of interface
	return 2* (r.length + r.breadth)
}

func MakeRectangle(length, breadth float32) Rectangle { // Return Value
	return Rectangle{
		length,
		breadth,
	}
}