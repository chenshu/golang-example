package main

import (
    "io"
    "os"
)

func main() {
    f, _ := os.Open("ex7.go")
    io.Copy(os.Stdout, f)
}
