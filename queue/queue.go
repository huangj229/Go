package queue

import (
	"sync"
	"fmt"
)
 

/*大写开头的类型是可以被全局访问的*/
type Queue struct {
	items []interface{}  // 任意数据的切片（vector）类型
	lock sync.RWMutex     //  读写锁，多线程安全
}

//加入数据
/* 传入参数为什么不能写 items []interface{} 
而要写 items ...interface{}
因为Go不会讲传入的任意类型的切片自动转型成为 []interface{}类型
这是因为 interface{} 会占用两个字长的存储空间，
一个是自身的 methods 数据，一个是指向其存储值的指针，
也就是 interface 变量存储的值，因而 slice []interface{} 
其长度是固定的N*2，但是 []T 的长度是N*sizeof(T)

https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/
*/
func (q *Queue) Put (items ...interface{}) { // ...在前面，语法糖，可变参数 http://c.biancheng.net/view/60.html
	if len(items) == 0 { // 切片无数据
		return 
	}

	q.lock.Lock()
	q.items = append(q.items, items...) // ...在后面，append需要接受列表展开的作为参数  http://c.biancheng.net/view/28.html
	q.lock.Unlock()

}

func (q *Queue) Show () {
	for _,arg := range q.items {
		fmt.Println(arg)
	}
}