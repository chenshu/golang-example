package stack

import (
    "testing"
)

func TestPushPop(t *testing.T) {
    s := new(Stack)
    s.Push(5)
    if s.Pop() != 5 {
        t.Log("Pop doesn't give 5")
        t.Fail()
    }
}
