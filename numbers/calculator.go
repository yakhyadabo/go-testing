package numbers

import "fmt"

func main()  {
	fmt.Println("Hello Calculator... ")
}

func add(a, b int) int {
	return a+b
}

func Sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum = add(sum, num)
	}
	return sum
}