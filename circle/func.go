package Circle

type Circle struct {
	radius float64
}

/*
实现函数,函数接受参数
输出两个参数的和
*/
func Calculate(a int, b int) int {
    return a + b
}

/*
实现circle类的方法
这个不会写对象的值 const
*/
func (c Circle) GetArea() float64 {
      return c.radius * c.radius * 3.14
}

/*
实现circle类的方法
这个会写对象的值 非const
*/
func (c* Circle) SetRadius(radius float64) {
	    c.radius = radius
}
