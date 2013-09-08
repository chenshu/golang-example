package main

import (
    "fmt"
    "strconv"
)

type stack struct {
    i int
    data [10]int
}

func (s *stack) Push(v int) {
    s.data[s.i] = v
    s.i++
}

func (s *stack) Pop() int {
    s.i--
    return s.data[s.i]
}

func (s stack) String() string {
    var str string
    for i := 0; i < s.i; i++ {
        str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
    }
    return str
}

func main() {
    var s stack
    //s := new(stack)
    s.Push(25)
    s.Push(14)
    fmt.Printf("stack %v\n", s)
}
