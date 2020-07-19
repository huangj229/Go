/*
通过一个圆形的类
展示Go语言的使用方法
和隐含的设计思想
*/

/*包名等同于外层文件夹名，不等于本文件名，一般使用小写，代表一个模块*/
package circle

/*引入包，Go语言要求包在文件被用到，否则会报错*/
import "math"

/*
定义类型，类型首字母需要是大写，使用大写的类型可以被其他文件导入并访问，小写的为本文件可见
按照GO的设计，类型只需要声明成员变量，私有的变量小写即可，例如圆形的半径是小写的
*/
type Circle struct {
	radius float64
}

/*
实现circle类的方法，实现访问半径的接口
这个不会写对象的值 类似于c++的const
注意：函数名大写开头表示public
*/
func (c Circle) GetRadius() float64 {
	return c.radius
}

/*
实现circle类的方法，实现设置半径的接口
这个会写对象的值 类似于c++的非const函数
*/
func (c* Circle) SetRadius(radius float64) {
	    c.radius = radius
}

/*
实现circle类的方法，实现计算圆的面积的接口
*/

func (c Circle) GetArea() float64 {
	return c.radius * c.radius * 3.14
}

/*
实现非类成员函数，该函数接收两个圆的实例，返回圆形面积的差值
*/
func CalculateAreaDiff(a Circle, b Circle) float64 {
    return math.Abs(a.GetArea()-b.GetArea())
}
