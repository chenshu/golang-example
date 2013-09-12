package main

import (
    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    "github.com/chenshu/golang-example/arith"
)

func main() {
    ar := new(arith.Arith)

    server := rpc.NewServer()
    server.Register(ar)

    server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)

    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go server.ServeCodec(jsonrpc.NewServerCodec(conn))
    }
}
