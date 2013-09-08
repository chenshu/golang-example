package main

import (
    "fmt"
)

func sort(arr []int) {
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i] > arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
}

func main() {
    var arr []int = []int{83, -3, 23, 9, 76}
    fmt.Printf("%v\n", arr)
    sort(arr)
    fmt.Printf("%v\n", arr)
}
