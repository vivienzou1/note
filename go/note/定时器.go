/**
Go语言中可以使用time.NewTimer或time.NewTicker来设置一个定时器，这个定时器会绑定在你的当前channel中，通过channel的阻塞通知机器来通知你的程序。

下面是一个timer的示例。
*/
package main
 
import "time"
import "fmt"
 
func main() {
    timer := time.NewTimer(2*time.Second)
 
    <- timer.C
    fmt.Println("timer expired!")
}

//------------------------demo2-------------------------
//上面的例程看起来像一个Sleep，是的，不过Timer是可以Stop的。你需要注意Timer只通知一次。如果你要像C中的Timer能持续通知的话，你需要使用Ticker。下面是Ticker的例程：

package main
 
import "time"
import "fmt"
 
func main() {
    ticker := time.NewTicker(time.Second)
 
    for t := range ticker.C {
        fmt.Println("Tick at", t)
    }
}
//上面的这个ticker会让你程序进入死循环，我们应该放其放在一个goroutine中。下面这个程序结合了timer和ticker
package main
 
import "time"
import "fmt"
 
func main() {
 
    ticker := time.NewTicker(time.Second)
 
    go func () {
        for t := range ticker.C {
            fmt.Println(t)
        }
    }()
 
    //设置一个timer，10钞后停掉ticker
    timer := time.NewTimer(10*time.Second)
    <- timer.C
 
    ticker.Stop()
    fmt.Println("timer expired!")
}
