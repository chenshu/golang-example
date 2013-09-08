package main

import (
    "fmt"
)

func avg(slice []float64) (ret float64) {
    var sum float64 = 0.0
    for _, v := range slice {
        sum += v
    }
    return sum / float64(len(slice))
}

func main() {
    slice := make([]float64, 10)
    for i, _ := range slice {
        slice[i] = float64(3 * i + i)
    }
    fmt.Printf("%f\n", avg(slice))
}
