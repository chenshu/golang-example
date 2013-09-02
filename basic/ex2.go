package main

import (
    "fmt"
)

func m1() {
    for i := 1; i <= 100; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

func m2() {
    var p bool
    for i := 1; i <= 100; i++ {
        p = false
        if i % 3 == 0 {
            fmt.Printf("Fizz")
            p = true
        }
        if i % 5 == 0 {
            fmt.Printf("Buzz")
            p = true
        }
        if !p {
            fmt.Printf("%d", i)
        }
        fmt.Println()
    }
}

func main() {
    m1()
    m2()
}
