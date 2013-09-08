package main

import (
    "fmt"
)

func max(arr []int) int {
    m := arr[0]
    for i := 1; i < len(arr); i++ {
        if (arr[i] > m) {
            m = arr[i]
        }
    }
    return m
}

func min(arr []int) int {
    m := arr[0]
    for i := 1; i < len(arr); i++ {
        if (arr[i] < m) {
            m = arr[i]
        }
    }
    return m
}

func main() {
    arr := []int{4, 6, 1, 9, 7, 2}
    fmt.Printf("%T\n", arr)
    fmt.Println(max(arr))
    fmt.Println(min(arr))
    var arr1 []int = []int{4, 6, 1, 9, 7, 2}
    fmt.Printf("%T\n", arr1)
    fmt.Println(max(arr1))
    fmt.Println(min(arr1))
    var arr2 [6]int = [6]int{4, 6, 1, 9, 7, 2}
    fmt.Printf("%T\n", arr2)
    //fmt.Println(max(arr2))
    //fmt.Println(min(arr2))
}
