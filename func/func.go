package circle
/*
圆形类用来展示函数使用
*/
type Circle struct {
	radius float64
}

type interface{
	
}

/*
实现最普通的函数
函数接受参数原始类型的参数
返回两个参数的和
*/
func Calculate(a int, b int) int {
    return a + b
}

/*
实现circle类的方法
该实现不会写对象的值 const
*/
func (c Circle) GetArea() float64 {
      return c.radius * c.radius * 3.14
}

/*
实现circle类的方法
该实现会写对象的值 非const
*/
func (c* Circle) SetRadius(radius float64) {
	    c.radius = radius
}
