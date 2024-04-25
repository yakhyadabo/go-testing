package calculator

import "fmt"

func main()  {
	fmt.Println("Hello Calculator... ")
}

func Add(a, b int) int {
	return a+b
}

func Sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum = Add(sum, num)
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}