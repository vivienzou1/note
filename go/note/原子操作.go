/**
说到并发就需要说说原子操作，相信大家还记得我写的那篇《无锁队列的实现》一文，里面说到了一些CAS – CompareAndSwap的操作。Go语言也支持。你可以看一下相当的文档

我在这里就举一个很简单的示例：下面的程序有10个goroutine，每个会对cnt变量累加20次，所以，最后的cnt应该是200。如果没有atomic的原子操作，那么cnt将有可能得到一个小于200的数。

下面使用了atomic操作，所以是安全的。
*/
package main
 
import "fmt"
import "time"
import "sync/atomic"
 
func main() {
    var cnt uint32 = 0
    for i := 0; i < 10; i++ {
        go func() {
            for i:=0; i<20; i++ {
                time.Sleep(time.Millisecond)
                atomic.AddUint32(&cnt, 1)
            }
        }()
    }
    time.Sleep(time.Second)//等一秒钟等goroutine完成
    cntFinal := atomic.LoadUint32(&cnt)//取数据
    fmt.Println("cnt:", cntFinal)
}
