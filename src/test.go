package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("利用できる CPU のコア数",runtime.NumCPU()) //利用できる CPU のコア数
    fmt.Println("使用するコア数",runtime.GOMAXPROCS(0)) //使用するコア数
    fmt.Println("現存している goroutine の数",runtime.NumGoroutine()) //現存している goroutine の数
}