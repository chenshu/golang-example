package main

import (
    "fmt"
    "./funcmap"
    "reflect"
)

type A struct {
    B int
    C string
}

func (a *A) Foo(name string) (ret, err string) {
    return "Hi, " + name, "foo"
}
func (a *A) Bar(name string) (ret, err string) {
    return "Hello, " + name, "bar"
}

func main() {
    funcs := funcmap.New()

    a := new(A)

    err := funcs.Bind("a.foo", a.Foo)
    if err != nil {
        fmt.Printf("Bind: %s\n", err)
    }
    ret, err := funcs.Call("a.foo", "fred");
    if err != nil {
        fmt.Printf("Call: %s\n", err)
    }
    fmt.Printf("%v\n", ret)

    err = funcs.Bind("a.bar", a.Bar)
    if err != nil {
        fmt.Printf("Bind: %s\n", err)
    }
    ret, err = funcs.Call("a.bar", "fred");
    if err != nil {
        fmt.Printf("Call: %s\n", err)
    }
    fmt.Printf("%v\n", ret)

    t := new(A)
    v := reflect.ValueOf(t)
    fmt.Println(v)
    fmt.Println(v.NumMethod())
    s := v.Type()
    fmt.Println(s)
    for i:= 0; i < v.NumMethod(); i++ {
        m := v.Method(i)
        fmt.Printf("%d: %s %s = %v\n", i, s.Method(i).Name, m.Type(), m.Interface())
        fmt.Printf("%s\n", s.Method(i).Func)
        err := funcs.Bind("A."+s.Method(i).Name, m.Interface())
        if err != nil {
            fmt.Printf("Bind: %s\n", err)
        }
        ret, err = funcs.Call("A."+s.Method(i).Name, "fred");
        if err != nil {
            fmt.Printf("Call: %s\n", err)
        }
        fmt.Printf("%v\n", ret)
    }
}
