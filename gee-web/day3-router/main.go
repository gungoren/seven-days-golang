package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(context *gee.Context) {
		context.String(http.StatusOK, "Hello %s, you're at %s\n", context.Query("name"), context.Path)
	})

	r.GET("/hello/:name", func(context *gee.Context) {
		context.String(http.StatusOK, "Hello %s, you're at %s\n", context.Param("name"), context.Path)
	})

	r.GET("/assets/*filepath", func(context *gee.Context) {
		context.JSON(http.StatusOK, gee.H{"filepath": context.Param("filepath")})
	})

	r.Run(":9995")
}
