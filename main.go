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
}