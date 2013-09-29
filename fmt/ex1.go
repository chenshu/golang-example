package main

import (
    "fmt"
)

func main() {
    var foo string = "This is a simple string"
    fmt.Printf("%v\n", foo)
    fmt.Printf("%+v\n", foo)
    fmt.Printf("%#v\n", foo)
    fmt.Printf("%T\n", foo)
    fmt.Printf("%%\n", foo)
}
