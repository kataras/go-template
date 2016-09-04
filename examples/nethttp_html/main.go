// Package main is a small example for html/template, same for all other
package main

import (
	"github.com/kataras/go-template"
	"github.com/kataras/go-template/html"
	"net/http"
)

type mypage struct {
	Title   string
	Message string
}

func main() {
	// create our html template engine, using a Layout
	tmplEngine := html.New(html.Config{Layout: "layouts/layout.html"})

	loader := template.NewLoader()
	loader.Directory("./templates", ".html") // defaults
	// for binary assets:
	// loader.Directory(dir string, ext string).Binary(assetFn func(name string) ([]byte, error), namesFn func() []string)

	// load the templates inside ./templates with extension .html
	//using the html template engine
	err := loader.LoadEngine(tmplEngine)
	if err != nil {
		panic("While parsing the template files: " + err.Error())
	}

	http.Handle("/", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// first parameter the writer
		// second parameter any page context (look ./templates/mypage.html) and you will understand
		// third parameter is optionally, is a map[string]interface{}
		// which you can pass a "layout" to change the layout for this specific render action
		// or the "charset" to change the defaults which is "utf-8"
		err := tmplEngine.ExecuteWriter(res, "mypage.html", mypage{"My Page title", "Hello world!"})
		if err != nil {
			res.Write([]byte(err.Error()))
		}
	}))

	println("Open a browser tab & go to localhost:8080")
	http.ListenAndServe(":8080", nil)
}
