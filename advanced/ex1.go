package main

import (
    "fmt"
)

type e interface{}

func fn(v e) e {
    switch v.(type) {
    case int:
        return v.(int) * 2
    case string:
        return v.(string) + v.(string)
    }
    return v
}

func map1(f func(i e) e, arr []e) []e {
    ret := make([]e, len(arr))
    for k, v := range arr {
        ret[k] = f(v)
    }
    return ret
}

func main() {
    arr := []e{5, 4, 2, 8, 9}
    for _, v := range map1(fn, arr) {
        fmt.Println(v)
    }
    arr = []e{"abc", "def", "xyz"}
    for _, v := range map1(fn, arr) {
        fmt.Println(v)
    }
}
