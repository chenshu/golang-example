package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc/jsonrpc"
    "github.com/chenshu/golang-example/arith"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing: ", err)
    }
    defer conn.Close()

    args := &arith.Args{7, 8}
    var reply int

    args1 := &arith.Args{6, 4}
    var reply1 arith.Quotient

    c := jsonrpc.NewClient(conn)

    for i := 0; i < 1; i++ {
        err = c.Call("Arith.Multiply", args, &reply)
        if err != nil {
            log.Fatal("arith error: ", err)
        }
        fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

        err = c.Call("Arith.Divide", args1, &reply1)
        if err != nil {
            log.Fatal("arith error: ", err)
        }
        fmt.Printf("Arith: %d/%d=%d(%d)\n", args1.A, args1.B, reply1.Quo, reply1.Rem)
    }
}
