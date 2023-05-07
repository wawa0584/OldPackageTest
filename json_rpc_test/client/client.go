package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//1。建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic("连接失败1")
	}
	var reply *string = new(string)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "bobby", reply)
	//client.Hello("bobby", reply)
	if err != nil {
		panic("连接失败")
	}
	fmt.Println(*reply)
}
