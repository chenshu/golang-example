package main

import (
    "fmt"
)

func gen(base int) []int {
    var i int = 0
    r := make([]int, base)
    for {
        switch i {
            case 0:
            r[i] = 1
            i++
            case 1:
            r[i] = 1
            i++
        default:
            r[i] = r[i-1] + r[i-2]
            i++
        }
        if i == base {
            break
        }
    }
    return r
}

func main() {
    arr := gen(5)
    for _, v := range arr {
        fmt.Printf("%d\t", v)
    }
    fmt.Println()
}
