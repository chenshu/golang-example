package main

import (
    "fmt"
)

func plusTwo() func(i int) int {
    return func(x int) int {
        return 2 + x
    }
}

func plusX(x int) func(i int) int {
    return func(y int) int {
        return x + y
    }
}

func main() {
    p := plusTwo()
    fmt.Printf("%v\n", p(2))

    p1 := plusX(2)
    fmt.Printf("%v\n", p1(2))
}
