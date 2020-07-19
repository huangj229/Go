package test

import (
	"../queue"
	"../shape"
	"fmt"
)
/*
测试定义的基本类
*/
func TestFunc() {

	{
    fmt.Println("******************Circle Test*******************************")
	var c1 shape.Circle 
    var c2 shape.Circle 
	c1.SetRadius(1)
	c2.SetRadius(2)
	fmt.Println("c1.radius =", c1.GetRadius(), ",c2.radius =", c2.GetRadius(), ",AreaDiff =", shape.CalculateAreaDiff(c1,c2))
    fmt.Println("*********************End************************************")
    }
	{
    fmt.Println("*******************Rectangle Test***************************")
	var r1 shape.Rectangle
	r1.SetRectangle(2,3)
	w,h := r1.GetRectangle()
    fmt.Println("w =" , w, "h =", h,"area =", r1.GetArea())
    fmt.Println("*********************End************************************")
    }
    {
    fmt.Println("******************Queue Test********************************")
	var myQ queue.Queue
	myQ.Put([]float64{1.23, 2.11, 3.77}, 3)
	myQ.Put("你好")
	fmt.Println("myQ len=", myQ.Len())
    myQ.Show()
	fmt.Println("*********************End************************************")
	}
}
