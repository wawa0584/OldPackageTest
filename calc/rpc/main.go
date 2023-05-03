package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

type Company struct {
	Name string
}
type Employee struct {
	Name    string
	company Company
}

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintLn(employee Employee) {
	/*
		1。建立连接 http
		2.序列化成json
		3.发送json字符串
		4.调用成功后实际上是你接收到一个二进制的数据
		5.等待服务器发送结果
		6.将服务器返回的结果解析成PrintResult结果 反序列化
	*/
}

func main() {
	//现在我们想把add函数变成一个远程的函数调用 也就意味着需要把Add函数放到远程服务器上去运行

	//这里有一段逻辑 这个逻辑扣件库存 但是库存服务是一个独立的系统reduce 如何调用牵扯到网络 做成一个web服务(gin beego net/httpServer）
	//1.函数调用的参数如何传递-json 数据格式的协议  /xml/protobuf/msgpack
	//现在网络调用两个短 客户端   将数据传输到gin
	//gin 服务器 服务端解析数据格式
	//json并不少一个高性能的方式
	//fmt.Println(Add(1, 2))\

	//放两外一台服务器上  需要将本地的内存对象struct 这样不行 可行的方式将struct转换成json
	fmt.Println(Employee{
		Name: "wjb",
		company: Company{
			Name: "1111",
		},
	})
	//远程的服务需要将二进制反解成struct对象  但是大型的分布式系统是不行的
}
