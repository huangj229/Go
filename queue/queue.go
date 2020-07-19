/*通过一个非基本数据类型，展示Go语言使用方法*/
package queue

/*sync用来使用读写锁*/
import (
	"sync"
	"fmt"
)
 

/*大写开头的类型是可以被全局访问的*/
type Queue struct {
	items []interface{}  // 任意数据（空接口）的切片（vector）类型
	lock sync.RWMutex     //  读写锁，多线程安全
}

/*
Put接口，向队列后加入任意元素
*/
func (q *Queue) Put (items ...interface{}) { 
	if len(items) == 0 { 
		return 
	}

	q.lock.Lock() //  加锁保证线程安全
	q.items = append(q.items, items...) 
	q.lock.Unlock()

}

/*
Show接口，打印队列所有值
*/
func (q *Queue) Show () {
	for _,arg := range q.items {
		fmt.Println(arg)
	}
}
/*
Pop接口，推出队列头数据
*/
func (q *Queue) Pop() interface{} {
	if len(q.items) == 0 {
		return nil
	}

	q.lock.Lock()
	head := q.items[0]
	q.items = q.items[1:]
	q.lock.Unlock()
	return head
}

/*
Len接口，打印队列长度
*/
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.items)

}