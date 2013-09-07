package main

import (
    "fmt"
)

func main() {
    s := make([]float64, 10)
    for i := range s {
        s[i] = float64(3.0 * i + i)
    }
    fmt.Printf("%v\n", s)
    sum := 0.0
    for _, v := range s {
        sum += v
    }
    avg := sum/float64(len(s))
    fmt.Printf("%v\n", avg)
}
