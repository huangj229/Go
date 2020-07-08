package main

import "fmt"

func main() {
	slice := []int{1,2,3}
	fmt.Println("slice:", slice)
	slice = append(slice, 4)
	fmt.Println("slice:", slice)
}