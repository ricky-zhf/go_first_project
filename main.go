package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//定义一个中间件，统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("go into m1...")
	start := time.Now()
	//c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	//c.Next()//执行当前请求的后续的操作
	c.Abort() //不调用后续处理请求的函数,但是会执行完当前函数。
	//如果不想执行当前函数后续流程需要使用return
	cost := time.Since(start) // 计算耗时
	fmt.Println("time costs", cost)
}
func m2(c *gin.Context) {
	fmt.Println("go into m2")
	c.Next()
	fmt.Println("exit m2")
}
func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddleware()) //全局注册中间件
	r.GET("/index", func(c *gin.Context) {
		fmt.Println("indexhandler...")
		c.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	})
	r.Run(":9090")
}

//通常中间件不会写成函数的形式，而是使用闭包来实现。
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*判断当前用户是否登录的中间件
		if 用户已登录 {
			c.Next()
		} else {
			c.Abort()
		}
		*/
	}
}
