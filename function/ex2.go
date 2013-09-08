package main

import (
    "fmt"
)

func sort(x, y int) (int, int) {
    if x > y {
        return y, x
    } else {
        return x, y
    }
}

func main() {
    fmt.Println(sort(2, 7))
    fmt.Println(sort(7, 2))
}
