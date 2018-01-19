package main

import "fmt"
import "time"
import "math/rand"
import "runtime"
import "sync"

var total_tickets int32 = 10

//可简写成：var mutex sync.Mutex
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
	//多线程
	runtime.GOMAXPROCS(4)        //我的电脑是4核处理器，所以我设置了4
	rand.Seed(time.Now().Unix()) //生成随机种子

	for i := 0; i < 5; i++ { //并发5个goroutine来卖票
		go sell_tickets(i)
	}
	//等待线程执行完
	var input string
	fmt.Scanln(&input)
	fmt.Println(total_tickets, "done") //退出时打印还有多少票
}


/**
package main

import "fmt"
import "time"
import "math/rand"

var total_tickets int32 = 10

func sell_tickets(i int) {
	for {
		if total_tickets > 0 { //如果有票就卖
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets-- //卖一张票
			fmt.Println("id:", i, "  ticket:", total_tickets)
		} else {
			break
		}
	}
}

func main() {

	//下面这种会卖超
	//for i := 0; i < 5; i++ { //并发5个goroutine来卖票
	 //    go sell_tickets(i)
	//}

	//并发50个goroutine来卖票
	go func() {
		for i := 0; i < 50; i++ {
			sell_tickets(i)
		}
	}()

	//等待线程执行完
	var input string
	fmt.Scanln(&input)

	//退出时打印还有多少票
	fmt.Println(total_tickets, "done")
}

**/
