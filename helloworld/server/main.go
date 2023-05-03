package main

import (
	"net"
	"net/rpc"
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
	_ = rpc.RegisterName("helloService", &HelloService{})

	conn, _ := listener.Accept()
	rpc.ServeConn(conn)

}
