package main

import (
    "fmt"
)

func shower(c chan int, quit chan bool) {
    for {
        select {
            case num := <- c:
            fmt.Println(num)
            case <- quit:
            break
        }
    }
}

func main() {
    ch := make(chan int)
    quit := make(chan bool)
    go shower(ch, quit)
    for i := 0; i < 10; i++ {
        ch <- i
    }
    quit <- false
}
