package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()

	engine.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hello From Gee</h1>")
	})

	engine.GET("/hello", func(context *gee.Context) {
		context.String(http.StatusOK, "Hello %s, you're at %s\n", context.Query("name"), context.Path)
	})

	engine.POST("/login", func(context *gee.Context) {
		context.JSON(http.StatusOK, gee.H{
			"username": context.PostForm("username"),
			"password": context.PostForm("password"),
		})
	})

	engine.Run(":9995")
}
