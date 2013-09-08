package main

import (
    "fmt"
)

func m(arr []int, f func(a int) int) []int {
    rs := make([]int, len(arr))
    for i, v := range arr {
        rs[i] = f(v)
    }
    return rs
}

func main() {
    arr := []int{5, 4, 2, 8, 9}
    f := func(i int) int {
        return i * i
    }
    for _, v := range m(arr, f) {
        fmt.Println(v)
    }
}
