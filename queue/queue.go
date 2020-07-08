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
func (q *Queue) Put (items ...interface{}) { // 这里的参数不明白为什么这么写！！！
	if len(items) == 0 { // 切片无数据
		return 
	}

	q.lock.Lock()
	q.items = append(q.items, items...) // 将切片合在一起，第二个切片用三个点
	q.lock.Unlock()

}