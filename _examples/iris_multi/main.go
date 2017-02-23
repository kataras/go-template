package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

type mypage struct {
	Title   string
	Message string
}

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	// the html engine on ./templates folder compile all *.html files
	app.Adapt(view.HTML("./templates", ".html"))

	// add our second template engine with the same directory but with .amber file extension
	app.Adapt(view.Amber("./templates", ".amber"))

	app.Get("/render_html", func(ctx *iris.Context) {
		ctx.RenderWithStatus(iris.StatusOK, "hiHTML.html", map[string]interface{}{"Name": "You!"})
	})

	app.Get("/render_amber", func(ctx *iris.Context) {
		ctx.MustRender("hiAMBER.amber", map[string]interface{}{"Name": "You!"})
	})

	println("Open a browser tab & go to localhost:8080/render_html  & localhost:8080/render_amber")
	app.Listen(":8080")
}

// More can be found there:
// https://github.com/iris-contrib/examples/tree/master/template_engines
