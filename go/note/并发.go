package main

import "fmt"
import "time"
import "math/rand"
import "runtime"
import "sync"

//一共10张票
var total_tickets int32 = 10

//锁，可简写成：var mutex sync.Mutex
var mutex = &sync.Mutex{}

func sell_tickets(i int) {
	for {
		mutex.Lock()
		if total_tickets > 0 { //如果有票就卖
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets-- //卖一张票
			fmt.Println("id:", i, "  ticket:", total_tickets)
		}
		//加锁，防止卖超
		mutex.Unlock()
	}
}

func main() {
	//实现真正的并发，我的电脑是4核处理器，所以我设置了4
	runtime.GOMAXPROCS(4)

	//生成随机种子
	rand.Seed(time.Now().Unix())

	//多线程,并发5个goroutine来卖票
	for i := 0; i < 5; i++ {
		go sell_tickets(i)
	}

	//等待线程执行完
	var input string
	fmt.Scanln(&input)

	//退出时打印还有多少票
	fmt.Println(total_tickets, "done")
}
