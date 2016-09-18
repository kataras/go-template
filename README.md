<p align="center">
  <img src="/logo.jpg" height="400">
</p>

<a href="https://travis-ci.org/kataras/go-template"><img src="https://img.shields.io/travis/kataras/go-template.svg?style=flat-square" alt="Build Status"></a>
<a href="https://github.com/avelino/awesome-go"><img src="https://img.shields.io/badge/awesome-%E2%9C%93-ff69b4.svg?style=flat-square" alt="https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg"></a>
<a href="http://goreportcard.com/report/kataras/go-template"><img src="https://img.shields.io/badge/report%20card-A%2B-F44336.svg?style=flat-square" alt="Report card"></a>
<a href="https://github.com/kataras/go-template/blob/master/LICENSE"><img src="https://img.shields.io/badge/%20license-MIT%20%20License%20-E91E63.svg?style=flat-square" alt="License"></a>
<a href="https://github.com/kataras/go-template/releases"><img src="https://img.shields.io/badge/%20release%20-%20v0.0.3-blue.svg?style=flat-square" alt="Releases"></a>
<a href="#docs"><img src="https://img.shields.io/badge/%20docs-reference-5272B4.svg?style=flat-square" alt="Read me docs"></a>
<a href="https://kataras.rocket.chat/channel/go-template"><img src="https://img.shields.io/badge/%20community-chat-00BCD4.svg?style=flat-square" alt="Chat"></a>
<a href="https://golang.org"><img src="https://img.shields.io/badge/powered_by-Go-3362c2.svg?style=flat-square" alt="Built with GoLang"></a>

The package go-template provides the easier way to use templates via template engine,
supports multiple template engines with different file extensions.

> It's a cross-framework means that it's 100% compatible with standard net/http, iris & fasthttp


6 template engines are supported:
- standard html/template
- amber
- django
- handlebars
- pug(jade)
- markdown

It's already tested on production & used on [Iris](https://github.com/kataras/iris) and [Q](https://github.com/kataras/q) web framework.

Installation
------------
The only requirement is the [Go Programming Language](https://golang.org/dl), at least v1.7.

```bash
$ go get -u github.com/kataras/go-template
```


Examples
------------

Run them from the [/examples](https://github.com/kataras/go-template/tree/master/examples) folder.


Otherwise, you can view examples via Iris example's repository [here](https://github.com/iris-contrib/examples/tree/master/template_engines).

Read the [kataras/iris/template.go](https://github.com/kataras/iris/blob/master/template.go) to see how Iris uses this package.

Docs
------------

Read the [godocs](https://godoc.org/github.com/kataras/go-template).


**Iris Quick look**

[Iris](https://github.com/kataras/iris) is the fastest web framework for Go, so far. It's based on [fasthttp](https://github.com/valyala/fasthttp), check that out if you didn't yet.

[Examples](https://github.com/iris-contrib/examples/tree/master/template_engines) covers the big picture, this is just a small code overview*


Make sure that you read & run the [iris-contrib/examples/template_engines](https://github.com/iris-contrib/examples/tree/master/template_engines) to cover the Iris + go-template part.

```go
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


```


**NET/HTTP Quick look**

```go
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

	// templates := template.NewMux()
	// templates.AddEngine(html.New()).Directory("./templates", ".html") // the defaults
	// or just use the default package-level mux:
	template.AddEngine(html.New())

	// add our second template engine with the same directory but with .amber file extension
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


```

> Note: All template engines have optional configuration which can be passed within $engine.New($engine.Config{})

FAQ
------------

- Q: Did this package works only with net/http ?
- A: No, this package can work with [Iris](https://github.com/kataras/iris) & [fasthttp](https://github.com/valyala/fasthttp) too, look [here for more](https://github.com/kataras/iris/blob/master/template.go).



- Q: How can I make my own template engine?
- A: Simply, you have to implement only **3  functions**, for load and execute the templates. One optionally (**Funcs() map[string]interface{}**) which is used to register any SharedFuncs.
	The simplest implementation, which you can look as example, is the Markdown Engine, which is located [here](https://github.com/kataras/go-template/tree/master/markdown).

```go

type (
	// Engine the interface that all template engines must implement
	Engine interface {
		// LoadDirectory builds the templates, usually by directory and extension but these are engine's decisions
		LoadDirectory(directory string, extension string) error
		// LoadAssets loads the templates by binary
		// assetFn is a func which returns bytes, use it to load the templates by binary
		// namesFn returns the template filenames
		LoadAssets(virtualDirectory string, virtualExtension string, assetFn func(name string) ([]byte, error), namesFn func() []string) error

		// ExecuteWriter finds, execute a template and write its result to the out writer
		// options are the optional runtime options can be passed by user
		// an example of this is the "layout" or "gzip" option
		ExecuteWriter(out io.Writer, name string, binding interface{}, options ...map[string]interface{}) error
	}

	// EngineFuncs is optional interface for the Engine
	// used to insert the Iris' standard funcs, see var 'usedFuncs'
	EngineFuncs interface {
		// Funcs should returns the context or the funcs,
		// this property is used in order to register the iris' helper funcs
		Funcs() map[string]interface{}
	}

	// EngineRawExecutor is optional interface for the Engine
	// used to receive and parse a raw template string instead of a filename
	EngineRawExecutor interface {
		// ExecuteRaw is super-simple function without options and funcs, it's not used widely
		ExecuteRaw(src string, wr io.Writer, binding interface{}) error
	}
)

```

Explore [these questions](https://github.com/kataras/go-template/issues?go-template=label%3Aquestion) or navigate to the [community chat][Chat].

Versioning
------------

Current: **v0.0.3**


People
------------
The author of go-template is [@kataras](https://github.com/kataras).

If you're **willing to donate**, feel free to send **any** amount through paypal

[![](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=makis%40ideopod%2ecom&lc=GR&item_name=Iris%20web%20framework&item_number=iriswebframeworkdonationid2016&currency_code=EUR&bn=PP%2dDonationsBF%3abtn_donateCC_LG%2egif%3aNonHosted)


Contributing
------------
If you are interested in contributing to the go-template project, please make a PR.

License
------------

This project is licensed under the MIT License.

License can be found [here](LICENSE).

[Travis Widget]: https://img.shields.io/travis/kataras/go-template.svg?style=flat-square
[Travis]: http://travis-ci.org/kataras/go-template
[License Widget]: https://img.shields.io/badge/license-MIT%20%20License%20-E91E63.svg?style=flat-square
[License]: https://github.com/kataras/go-template/blob/master/LICENSE
[Release Widget]: https://img.shields.io/badge/release-v0.0.3-blue.svg?style=flat-square
[Release]: https://github.com/kataras/go-template/releases
[Chat Widget]: https://img.shields.io/badge/community-chat-00BCD4.svg?style=flat-square
[Chat]: https://kataras.rocket.chat/channel/go-template
[ChatMain]: https://kataras.rocket.chat/channel/go-template
[ChatAlternative]: https://gitter.im/kataras/go-template
[Report Widget]: https://img.shields.io/badge/report%20card-A%2B-F44336.svg?style=flat-square
[Report]: http://goreportcard.com/report/kataras/go-template
[Documentation Widget]: https://img.shields.io/badge/documentation-reference-5272B4.svg?style=flat-square
[Documentation]: https://www.gitbook.com/book/kataras/go-template/details
[Language Widget]: https://img.shields.io/badge/powered_by-Go-3362c2.svg?style=flat-square
[Language]: http://golang.org
[Platform Widget]: https://img.shields.io/badge/platform-Any--OS-gray.svg?style=flat-square
