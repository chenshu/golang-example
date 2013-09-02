package main

import (
    "fmt"
)

func main() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d\n", i)
    }

    i := 0
    Here:
    fmt.Printf("%d\n", i)
    i++
    if i < 10 {
        goto Here
    }

    var j []int = []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
    for _, v := range j {
        fmt.Printf("%d\n", v)
    }

    var arr [10]int;
    for i := 0; i < 10; i++ {
        arr[i] = i
    }
    fmt.Printf("%v\n", arr)
}
