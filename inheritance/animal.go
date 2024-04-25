package inheritance

type Position struct {
	x,y int
}

type Animal struct {
	name string
	*Position
}

func MakeAnimal(name string, x,y int) Animal { // Return Value
	return Animal{
		name,
		&Position{x, y},
	}
}

func (p *Position) Move (x,y int)  {
	p.x = x
	p.y = y
}

func (a *Animal) GetX () int {
	return a.x
}

func (a *Animal) GetY () int {
	return a.y
}

// main function
//func main() {
//
//	a := Animal{"animal", &Position{7, 4}}
//
//	fmt.Printf("Initial Position x: %d y: %d\n", a.x, a.y)
//
//	a.Move(11,19)
//	fmt.Printf("Final Position x: %d y: %d", a.x, a.y)
//}