package main
 
import (
    "fmt"
    "reflect"
)

func t1() {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))
    fmt.Println("value:", reflect.ValueOf(x))
    v := reflect.ValueOf(x)
    fmt.Println("type:", v.Type())
    fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
    fmt.Println("value:", v.Float())

    fmt.Println(v.Interface())
    fmt.Println(v.Interface().(float64))
}

func t2() {
    var x uint8 = 'x'
    fmt.Println("type:", reflect.TypeOf(x))
    fmt.Println("value:", reflect.ValueOf(x))
    v := reflect.ValueOf(x)
    fmt.Println("type:", v.Type())
    fmt.Println("kind is uint8:", v.Kind() == reflect.Uint8)
    fmt.Println("value:", v.Uint())
    fmt.Println("value:", uint8(v.Uint()))

    fmt.Println(v.Interface())
    fmt.Println(v.Interface().(uint8))
}

func t3() {
    type MyInt int
    var x MyInt = 7
    fmt.Println("type:", reflect.TypeOf(x))
    fmt.Println("value:", reflect.ValueOf(x))
    v := reflect.ValueOf(x)
    fmt.Println("type:", v.Type())
    fmt.Println("kind is int:", v.Kind() == reflect.Int)
    fmt.Println("value:", v.Int())
    fmt.Println("value:", int32(v.Int()))

    fmt.Println(v.Interface())
    fmt.Println(v.Interface().(MyInt))
}

func t4() {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(&x))
    fmt.Println("value:", reflect.ValueOf(&x))
    p := reflect.ValueOf(&x)
    fmt.Println("type of p:", p.Type())
    fmt.Println("settability of p:" , p.CanSet())
    v := p.Elem()
    fmt.Println("settability of v:" , v.CanSet())
    v.SetFloat(7.1)
    fmt.Println(v.Interface())
    fmt.Println(x)
}

func t5() {
    type T struct {
        A int
        B string
    }
    t := T{23, "skidoo"}
    s := reflect.ValueOf(&t).Elem()
    fmt.Println(s)
    fmt.Println("settability of s:" , s.CanSet())
    typeOfT := s.Type()
    fmt.Println(typeOfT)
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v\n", i,
            typeOfT.Field(i).Name, f.Type(), f.Interface())
    }
    s.Field(0).SetInt(77)
    s.Field(1).SetString("Sunset Strip")
    fmt.Println("t is now", t)
}
 
func main() {
    t1()
    t2()
    t3()
    t4()
    t5()
}
