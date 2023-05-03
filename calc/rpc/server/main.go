package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//http:127.0.0.1/add?a=1&b=2
	//返回给前端
	//1.callId r.URL.Path 2.数据的传输协议http的协议
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm() //解析参数
		fmt.Println("path", r.URL.Path)
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		JData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(JData)
	})
	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)

}
