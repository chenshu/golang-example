package main

import (
    "fmt"
)

func checkArgs(arg... int) {
    for i, n := range arg {
        fmt.Printf("%d\t%d\n", i, n)
    }
}

func main() {
    checkArgs(1, 3, 4, 5)
}
