package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply的值
	*reply = "hello, " + request
	return nil
}

func main() {
	listener, _ := net.Listen("tcp", ":1234")
	/*
		1.实力化一个sever
		2.注册处理逻辑
		3. 启动服务
	*/
	_ = rpc.RegisterName("HelloService", &HelloService{})

	for {
		//启动服务
		conn, _ := listener.Accept() //当一个新的连接进来的时候

		//1.callId  2.序列化和反序列化
		//跨语言调用 1.go语言的rpc序列和反序列协议 2能否替换成常见的序列化
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
