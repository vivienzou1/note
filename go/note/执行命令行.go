package main
import "os/exec"
import "fmt"
func main() {
    cmd := exec.Command("ping", "127.0.0.1")
    out, err := cmd.Output()
    if err!=nil {
        println("Command Error!", err.Error())
        return
    }
    fmt.Println(string(out))
}

//demo2
//正规一点的用来处理标准输入和输出的示例如下：

package main
 
import (
    "strings"
    "bytes"
    "fmt"
    "log"
    "os/exec"
)
 
func main() {
    cmd := exec.Command("tr", "a-z", "A-Z")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("in all caps: %q\n", out.String())
}
