package test

import (
    "fmt"
	"../queue"
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

	var c1 circle
	c1.radius = 10;
	fmt.Println(c1.getArea())
	c1.setRadius(20)
	fmt.Println(c1.getArea())

	var myint queue.Queue // 声明一个任意类型的队列
	myint.Put(3)
	myint.Put(4)
	myint.Show()
	tmp := myint.Pop()
	myint.Show()
	fmt.Println("tmp = ",tmp)
}

type circle struct {
	radius float64
}

/*实现函数*/
func calculate(a int, b int) int {
	fmt.Println("a = ", a)
	defer fmt.Println("b = ", b)
	fmt.Println("a2 = ", a)
    return a + b
}

/*实现circle类的方法*/
func (c circle) getArea() float64 {
      return c.radius * c.radius * 3.14
}

/*实现circle类的方法*/
func (c* circle) setRadius(radius float64) {
	    c.radius = radius
}

