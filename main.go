package main

import "fmt"

func main() {
	slice := []int{1,2,3}
	fmt.Println("slice:", slice)
	slice = append(slice, 4)
	fmt.Println("slice:", slice)

    for i := 0; i < len(slice); i++ {
		fmt.Println("index = ", i," value = ", slice[i])
	}

	for i,v := range slice {
		fmt.Println("index = ", i," value = ", v)
	}

    var a int = 1
	var b int = 2

	fmt.Println(calculate(a, b))
}

func calculate(a int, b int) int {
	fmt.Println("a = ", a)
	defer fmt.Println("b = ", b)
	fmt.Println("a2 = ", a)
    return a + b
}