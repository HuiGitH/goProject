package main
import (
    "fmt"
    "os"
    "github.com/go-vgo/robotgo"
    hook "github.com/robotn/gohook"
)  
func main() {  
    go func() {
        fmt.Println("请按回车键，并在5秒内把鼠标放在‘下一话’位置")

        // b := make([]byte, 1)
        // os.Stdin.Read(b)

        var n string
        fmt.Scanln("%s",&n)
        fmt.Println("倒计时5秒")
        robotgo.Sleep(5)
        mouseX, mouseY := robotgo.GetMousePos()
        for {
                robotgo.ScrollMouse(1, "down")
                robotgo.MoveClick(mouseX, mouseY, "left", false)
                robotgo.Sleep(1)
        }
        }()

        fmt.Println("--- 请在cmd窗口内按下esc键，程序将退出 ---")
        robotgo.EventHook(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
                fmt.Println("esc")
                robotgo.EventEnd()
        })
        s := robotgo.EventStart()
        <-robotgo.EventProcess(s)
        ok := robotgo.AddEvents("esc")
        if ok {
                os.Exit(1)
        }
        }