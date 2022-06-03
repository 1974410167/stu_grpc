package main

import (
	"fmt"
	"sync"
)

func main(){

	// 互斥锁
	mutex := &sync.Mutex{}
	// 加锁
	mutex.Lock()

	// Update共享变量 (比如切片，结构体指针等)

	// 解锁
	mutex.Unlock()

	// 共享锁
	rWMutex := &sync.RWMutex{}
	// 对于修改互斥
	rWMutex.Lock()
	// Update 共享变量
	rWMutex.Unlock()

	// 对于读取共享
	rWMutex.RLock()
	// Read 共享变量
	rWMutex.RUnlock()

	// 优雅的等待多个协程结束，然后再继续向下执行
	wg := &sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			// Do something
			wg.Done()
		}()
	}
	wg.Wait()
	// continue

	// 读写安全的Map
	// 使用Store(interface {}，interface {})添加元素。
	// 使用Load(interface {}) interface {}检索元素。
	// 使用Delete(interface {})删除元素。
	// 使用LoadOrStore(interface {}，interface {}) (interface {}，bool)检索或添加之前不存在的元素。如果键之前在map中存在，则返回的布尔值为true。
	// 使用Range遍历元素。
	m := &sync.Map{}

	// 添加元素
	m.Store(1, "one")
	m.Store(2, "two")

	// 获取元素1
	value, contains := m.Load(1)
	if contains {
		fmt.Printf("%s\n", value.(string))
	}

	// 返回已存value，否则把指定的键值存储到map中
	value, loaded := m.LoadOrStore(3, "three")
	if !loaded {
		fmt.Printf("%s\n", value.(string))
	}

	m.Delete(3)

	// 迭代所有元素
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%d: %s\n", key.(int), value.(string))
		return true
	})

	// 并发池，可以并发安全的get或者put元素
	pool := &sync.Pool{}

	type Connection struct {
		// 假如这是一个连接
		id int
	}

	pool.Put(&Connection{id: 1})
	pool.Put(&Connection{id: 2})
	pool.Put(&Connection{id: 3})
	connection := pool.Get().(*Connection)
	fmt.Printf("%d\n", connection.id)
	connection = pool.Get().(*Connection)
	fmt.Printf("%d\n", connection.id)
	connection = pool.Get().(*Connection)
	fmt.Printf("%d\n", connection.id)

	// 确保一个函数仅仅执行一次，如下，开启私服
	once := &sync.Once{}

	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("xxxx")
			})
			wg.Done()
		}()
	}
	wg.Wait()

}
