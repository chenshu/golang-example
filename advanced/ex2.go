package main

import (
    "fmt"
    "errors"
    "container/list"
)

type Value int

type Node struct {
    Value
    prev, next *Node
}

func (n *Node) Next() *Node {
    return n.next
}

type List struct {
    head, tail *Node
}

func (l *List) Front() *Node {
    return l.head
}

func (l *List) PushBack(v Value) *List {
    node := &Node{Value: v}
    if l.head == nil {
        l.head = node
    } else {
        node.prev = l.tail
        l.tail.next = node
    }
    l.tail = node
    return l
}

var errEmpty = errors.New("List is empty")

func (l *List) Pop() (v Value, err error) {
    if l.tail == nil {
        err = errEmpty
    } else {
        v = l.tail.Value
        l.tail = l.tail.prev
        if l.tail == nil {
            l.head = nil
        } else {
            l.tail.next = nil
        }
    }
    return v, err
}

func main() {
    l := list.New()
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(4)

    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Printf("%v\n", e.Value)
    }

    ll := new(List)
    ll.PushBack(1)
    ll.PushBack(2)
    ll.PushBack(4)

    for e := ll.Front(); e != nil; e = e.Next() {
        fmt.Printf("%v\n", e.Value)
    }
    fmt.Println()
    for v, err := ll.Pop(); err == nil; v, err = ll.Pop() {
        fmt.Printf("%v\n", v)
    }
}
