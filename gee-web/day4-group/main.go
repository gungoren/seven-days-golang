package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee </h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			//expect /hello?name=mehmet
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(context *gee.Context) {
			context.String(http.StatusOK, "Hello %s, you're at %s\n", context.Param("name"), context.Path)
		})

		v2.POST("/login", func(context *gee.Context) {
			context.JSON(http.StatusOK, gee.H{
				"username": context.PostForm("username"),
				"password": context.PostForm("password"),
			})
		})
	}

	r.Run(":9995")
}
