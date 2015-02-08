package main

import (
    "fmt"
)

type Node struct {
    data string
    next *Node
}

type Stack struct {
    top *Node
    size, capacity int
}

func NewStack(size int) *Stack {
    stack := &Stack{
        capacity: size,
    }
    return stack
}

func (s *Stack) Push (item string) {
    if s.size >= s.capacity {
        panic("stack is full")
    }
    node := &Node{
        data: item,
        next: s.top,
    }
    s.top = node
    s.size++
}

func (s *Stack) Pop () string {
    if s.size == 0 {
        panic("stack is empty")
    }
    n := s.top
    s.top = n.next
    n.next = nil
    ret := n.data
    n = nil
    s.size--
    return ret
}

func (s *Stack) Clear() {
    for n := s.top; n != nil; {
        t := n.next
        n.next = nil
        n = nil
        n = t
        s.size--
    }
}

func (s *Stack) String() string {
    ret := fmt.Sprintf("%d, [", s.size)
    for n := s.top; n != nil; {
        ret += n.data + ", "
        n = n.next
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
