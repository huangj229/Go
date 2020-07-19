package test

import (
	"../queue"
	"../circle"
	"fmt"
)
/*
测试定义的基本类
*/
func TestFunc() {

	{
    fmt.Println("******************Circle Test*******************************")
	var c1 circle.Circle 
    var c2 circle.Circle 
	c1.SetRadius(1)
	c2.SetRadius(2)
	fmt.Println("c1.radius = ", c1.GetRadius(), ",c2.radius = ", c2.GetRadius(), ",AreaDiff = ", circle.CalculateAreaDiff(c1,c2))
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
