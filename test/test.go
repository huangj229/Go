package test

import (
    "fmt"
	"../queue"
	"../func"
)

func Test() {
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

	var c1 func.Circle
	c1.radius = 10;
	fmt.Println(c1.GetArea())
	c1.SetRadius(20)
	fmt.Println(c1.GetArea())

	var myint queue.Queue // 声明一个任意类型的队列
	myint.Put(3)
	myint.Put(4)
	myint.Show()
	tmp := myint.Pop()
	myint.Show()
	fmt.Println("tmp = ",tmp)
}
