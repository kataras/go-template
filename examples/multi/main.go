// Package main uses the template.Mux here to simplify the steps
// and support of loading more than one template engine for a single app
// you can do the same things without the Mux, as you saw on the /html example folder
package main

import (
	"github.com/kataras/go-template"
	"github.com/kataras/go-template/amber"
	"github.com/kataras/go-template/html"
	"net/http"
)

type mypage struct {
	Title   string
	Message string
}

func main() {
	//  templates := template.NewMux()
	// 	templates.AddEngine(tmplEngine).Directory("./templates", ".html") // the defaults
	//  or just use the default package-level mux
	// defaults to dir "./templates" and ".html" as extension, you can change it by .Directory("./mydir", ".html")
	template.AddEngine(html.New())
	// set our second template engine with the same directory but with .amber file extension
	template.AddEngine(amber.New()).Directory("./templates", ".amber")

	// load all the template files using the correct template engine for each one of the files
	err := template.Load()
	if err != nil {
		panic("While parsing the template files: " + err.Error())
	}

	http.Handle("/render_html", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// first parameter the writer
		// second parameter any page context (look ./templates/mypage.html) and you will understand
		// third parameter is optionally, is a map[string]interface{}
		// which you can pass a "layout" to change the layout for this specific render action
		// or the "charset" to change the defaults which is "utf-8"

		// Does the same thing but returns the parsed template file results as string
		// useful when you want to send rich e-mails with a template
		// template.ExecuteString(name, pageContext, options...)

		err := template.ExecuteWriter(res, "hiHTML.html", map[string]interface{}{"Name": "You!"}) // yes you can pass simple maps instead of structs
		if err != nil {
			res.Write([]byte(err.Error()))
		}
	}))

	http.Handle("/render_amber", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		template.ExecuteWriter(res, "hiAMBER.amber", map[string]interface{}{"Name": "You!"})
	}))

	println("Open a browser tab & go to localhost:8080/render_html  & localhost:8080/render_amber")
	http.ListenAndServe(":8080", nil)
}
