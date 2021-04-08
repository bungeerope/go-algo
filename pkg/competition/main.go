package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * 此示例代码是错误使用方式
 * 如果data++,在 if 之前运行，将不会输出任何日志
 * 如果data++,在 print 之后运行，将会打印 data = 0
 * 如果data++,在 if 之后、 print之前运行，将会打印 data = 1
 */
func wrongMethod() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

/**
 * 在代码中加入阻塞方式防止数据竞争，并不是一个解决方案
 * 竞争条件是最难以发现的并发bug类型之一
 */
func noGracefulMethod() {
	var data int
	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

/**
 * 借助临界区内存同步防止数据竞争，简而言之：加锁
 */
func syncDataMethod() {
	var data int
	var memoryAccess sync.Mutex
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is 0.\n")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func deadLockMethod() {

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum = %v \n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

func main() {
	// 错误示例
	wrongMethod()

	// 不优雅方式保证程序运行
	noGracefulMethod()

	// 通过临界区内存同步防止数据竞争
	syncDataMethod()

	// 触发死锁竞争
	deadLockMethod()
}
