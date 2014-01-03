package main

import (
    "fmt"
)

type Stack struct {
    array []string
    cur int
    size int
}

func NewStack(size int) *Stack {
    stack := &Stack{
        array: make([]string, size),
        cur: 0,
        size: size,
    }
    return stack
}

func (s *Stack) Push (item string) {
    if s.cur >= s.size {
        panic("stack is full")
    }
    s.array[s.cur] = item
    s.cur++
}

func (s *Stack) Pop () string {
    if s.cur == 0 {
        panic("stack is empty")
    }
    s.cur -= 1
    return s.array[s.cur]
}

func (s *Stack) Clear() {
    s.cur = 0
}

func (s *Stack) String() string {
    ret := fmt.Sprintf("%d, [", s.cur)
    for i := 0; i < s.cur; i++ {
        ret += s.array[i] + ", "
    }
    ret += "]"
    return ret
}

func main() {
    var size int = 2
    fmt.Println(size)
    stack := NewStack(size)
    fmt.Println(stack.String())
    stack.Push("foo")
    fmt.Printf("push %s\n", "foo")
    fmt.Println(stack.String())
    stack.Push("bar")
    fmt.Printf("push %s\n", "bar")
    fmt.Println(stack.String())
    fmt.Printf("pop %s\n", stack.Pop())
    fmt.Println(stack.String())
    fmt.Printf("pop %s\n", stack.Pop())
    fmt.Println(stack.String())
}
