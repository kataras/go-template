package main

import (
	"github.com/kataras/go-template/amber"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

type mypage struct {
	Title   string
	Message string
}

// Iris examples covers the most part, including all 6 template engines and their configurations:
// https://github.com/iris-contrib/examples/tree/master/template_engines

func main() {

	iris.UseTemplate(html.New()) // the Iris' default if no template engines are setted.

	// add our second template engine with the same directory but with .amber file extension
	iris.UseTemplate(amber.New(amber.Config{})).Directory("./templates", ".amber")

	iris.Get("/render_html", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "hiHTML.html", map[string]interface{}{"Name": "You!"})
	})

	iris.Get("/render_amber", func(ctx *iris.Context) {
		ctx.MustRender("hiAMBER.amber", map[string]interface{}{"Name": "You!"})
	})

	println("Open a browser tab & go to localhost:8080/render_html  & localhost:8080/render_amber")
	iris.Listen(":8080")
}

// Iris examples covers the most part, including all 6 template engines and their configurations:
// https://github.com/iris-contrib/examples/tree/master/template_engines
