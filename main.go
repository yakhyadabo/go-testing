package main

import (
	"context"
	"fmt"
	"github.com/yakhyadabo/golang/testing/contextz"
	"github.com/yakhyadabo/golang/testing/inheritance"
	"github.com/yakhyadabo/golang/testing/interfacez"
	"log"
)

func main(){
	fmt.Println("")
	callInheritance()
	callInterface()
	callContext()
}

func callInheritance()  {
	fmt.Println("***")
	a := inheritance.MakeAnimal("animal", 7, 4)

	fmt.Printf("Initial Position x: %d y: %d\n", a.GetX(), a.GetY() )
	a.Move(11,19)
	fmt.Printf("Final Position x: %d y: %d", a.GetX(), a.GetY())
}

// main function
func callInterface() {

	// assigns value to struct members
	rect := interfacez.MakeRectangle(7, 4)
	circle := interfacez.MakeCircle( 4)

	// call calculate() with struct variable rect
	interfacez.Calculate(rect)
	interfacez.Calculate(circle)
}

func callContext() {
	ctx := context.Background()
	request := contextz.UserRequest{}
	val, err := request.FetchUserData(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
}