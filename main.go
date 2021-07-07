package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
		return
	}
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板,路径中的点代表项目所在目录
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse tamplate failed, err:", err)
	}
	//渲染模板,execute第一个值是解析出来的模板要返回给io。第二个代表传什么数据。
	// 利用给定数据渲染模板，并将结果写入w
	user := UserInfo{
		Name:   "rickyyyy",
		Gender: "男",
		Age:    25,
	}
	err = t.Execute(w, user)
	if err != nil {
		fmt.Println("render template failed, err:", err)
	}
}
