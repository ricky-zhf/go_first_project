package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	r := gin.Default()
	//解析模板文件
	r.LoadHTMLFiles("templates/index.tmpl")
	//渲染模板文件，访问index时，调用func函数。
	r.GET("/index", func(c *gin.Context) {
		//写返回值，写返回好的渲染文件,没有通过define命名是需要默认的文件名
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			//gin.H本质上就是一个map[string]interface{}
			"title": "rickyyy.com",
		})
	})
	r.Run(":9090") //启动server
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板,路径中的点代表项目所在目录
	t, err := template.ParseFiles("./index.tmpl")
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
