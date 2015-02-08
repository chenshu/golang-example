package main

import (
    "fmt"
    "reflect"
    "./funcmap"
    "encoding/json"
)

type A struct {
    P1 B
    P2 string
}

type B struct {
    P3 int
}

type C struct {
    P4 string
    P5 D
}

type D struct {
    P6 int
}

func (a *A) Foo(name string, age int, c C) (r1 string, r2 int) {
    return "Hi, " + name + "-" + c.P4, age + c.P5.P6
}

func main() {
    a1 := new(A)
    d1 := D{11}
    c1 := C{"world", d1}
    fmt.Println(a1.Foo("abc", 100, c1))

    funcs1 := funcmap.New()
    a2 := new(A)
    d2 := D{99}
    c2 := C{"china", d2}
    v := reflect.ValueOf(a2)
    for i := 0; i < v.NumMethod(); i++ {
        err := funcs1.Bind("test." + v.Type().Method(i).Name, v.Method(i).Interface())
        if err != nil {
            fmt.Printf("bind failed: %v\n", err)
        }
        ret, err := funcs1.Call("test." + v.Type().Method(i).Name, "xyz", 900, c2)
        if err != nil {
            fmt.Printf("call failed: %v\n", err)
        }
        fmt.Printf("%v %v\n", ret[0].String(), ret[1].Int())
    }

    jsonBytes, err := json.Marshal([]interface{}{"xyz", 900, c2})
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    fmt.Println(string(jsonBytes))
    var pp []interface{}
    err = json.Unmarshal(jsonBytes, &pp)
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    fmt.Printf("%d, %v\n", len(pp), pp)
    funcs2 := funcmap.New()
    a3 := new(A)
    v2 := reflect.ValueOf(a3)
    for i := 0; i < v2.NumMethod(); i++ {
        err := funcs2.Bind("test." + v2.Type().Method(i).Name, v2.Method(i).Interface())
        if err != nil {
            fmt.Printf("bind failed: %v\n", err)
        }
        ret, err := funcs2.Call("test." + v2.Type().Method(i).Name, pp...)
        if err != nil {
            fmt.Printf("call failed: %v\n", err)
        }
        fmt.Printf("%v %v\n", ret[0].String(), ret[1].Int())
    }
}
