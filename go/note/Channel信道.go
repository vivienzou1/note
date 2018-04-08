
//--------------demo1----------------
package main
 
import "fmt"
 
func main() {
    //创建一个string类型的channel
    channel := make(chan string)
 
    //创建一个goroutine向channel里发一个字符串
    go func() { channel <- "hello" }()
 
    msg := <- channel
    fmt.Println(msg)
}

//---------------------demo2 指定channel的buffer -----------------------------
package main
import "fmt"
 
func main() {
    channel := make(chan string, 2)
 
    go func() {
        channel <- "hello"
        channel <- "World"
    }()
 
    msg1 := <-channel
    msg2 := <-channel
    fmt.Println(msg1, msg2)
}

//---------Channel的阻塞-----------------
package main
 
import "fmt"
import "time"
 
func main() {
 
    channel := make(chan string) //注意: buffer为1
 
    go func() {
        channel <- "hello"
        fmt.Println("write \"hello\" done!")
 
        channel <- "World" //Reader在Sleep，这里在阻塞
        fmt.Println("write \"World\" done!")
 
        fmt.Println("Write go sleep...")
        time.Sleep(3*time.Second)
        channel <- "channel"
        fmt.Println("write \"channel\" done!")
    }()
 
    time.Sleep(2*time.Second)
    fmt.Println("Reader Wake up...")
 
    msg := <-channel
    fmt.Println("Reader: ", msg)
 
    msg = <-channel
    fmt.Println("Reader: ", msg)
 
    msg = <-channel //Writer在Sleep，这里在阻塞
    fmt.Println("Reader: ", msg)
}

//-----------------多个Channel的select---------------------------
package main
import "time"
import "fmt"
 
func main() {
    //创建两个channel - c1 c2
    c1 := make(chan string)
    c2 := make(chan string)
 
    //创建两个goruntine来分别向这两个channel发送数据
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "Hello"
    }()
    go func() {
        time.Sleep(time.Second * 1)
        c2 <- "World"
    }()
 
    //使用select来侦听两个channel
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
//注意：上面的select是阻塞的，所以，才搞出ugly的for i <2这种东西。

//-----------Channel select阻塞的Timeout--------------------
for {
    timeout_cnt := 0
    select {
    case msg1 := <-c1:
        fmt.Println("msg1 received", msg1)
    case msg2 := <-c2:
        fmt.Println("msg2 received", msg2)
    case  <-time.After(time.Second * 30)：
        fmt.Println("Time Out")
        timout_cnt++
    }
    if time_cnt > 3 {
        break
    }
}

//-----------------------Channel的无阻塞------------------
for {
    select {
    case msg1 := <-c1:
        fmt.Println("received", msg1)
    case msg2 := <-c2:
        fmt.Println("received", msg2)
    default: //default会导致无阻塞
        fmt.Println("nothing received!")
        time.Sleep(time.Second)
    }
}

//-------------------Channel的关闭----------------------------
package main
 
import "fmt"
import "time"
import "math/rand"
 
func main() {
 
    channel := make(chan string)
    rand.Seed(time.Now().Unix())
 
    //向channel发送随机个数的message
    go func () {
        cnt := rand.Intn(10)
        fmt.Println("message cnt :", cnt)
        for i:=0; i<cnt; i++{
            channel <- fmt.Sprintf("message-%2d", i)
        }
        close(channel) //关闭Channel
    }()
 
    var more bool = true
    var msg string
    for more {
        select{
        //channel会返回两个值，一个是内容，一个是还有没有内容
        case msg, more = <- channel:
            if more {
                fmt.Println(msg)
            }else{
                fmt.Println("channel closed!")
            }
        }
    }
}








