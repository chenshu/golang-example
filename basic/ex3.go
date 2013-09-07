package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    str := "A"
    for i := 0; i < 100; i++ {
        fmt.Println(str)
        str += "A"
    }

    s := "asSASA ddd dsjkdsjs dk我们"
    b := []byte(s)
    fmt.Printf("String %s\nLength: %d, Runes: %d\n", s, len(b), utf8.RuneCount(b))

    fmt.Printf("%v\n", s)
    sl := []rune(s)
    fmt.Printf("%v\n", sl)
    sl[3], sl[4], sl[5] = 'a', 'b', 'c'
    s2 := string(sl)
    fmt.Printf("%v\n", s2)

    t := "foobar"
    for i := len(t); i > 0; i-- {
        fmt.Printf("%c", t[i-1])
    }
    fmt.Println()
}
