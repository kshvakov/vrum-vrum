package main

import (
	"fmt"
	real_pongo2 "github.com/flosch/pongo2"
	"github.com/kshvakov/vrum-vrum/render"
	"github.com/kshvakov/vrum-vrum/render/pongo2"
	"github.com/kshvakov/vrum-vrum/server"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
)

type User struct {
	Name string
}

func init() {

	real_pongo2.RegisterFilter("HelloFunc", func(in *real_pongo2.Value, param *real_pongo2.Value) (*real_pongo2.Value, *real_pongo2.Error) {

		return real_pongo2.AsValue(fmt.Sprintf("Hello, %s", in.String())), nil
	})

	pongo2.ParseHTMLTemplates("template/")
}

func main() {

	app := server.New()

	errorHandlers(app)

	app.Use(func(c *server.Context) {

		c.Set("user", User{Name: "UserName"})
	})

	app.Use(func(c *server.Context) {

		c.Next()

		log.Print("end")
	})

	app.Get("/",
		func(c *server.Context) {

			c.Next()

			fmt.Println("log")
		},

		func(c *server.Context) {

			user := User{Name: "Undefined"}

			if u, ok := c.Get("user").(User); ok {

				user = u
			}

			pongo2.HTML(c, "index.html", map[string]interface{}{"user": user})

			c.Stop()
		},

		func(c *server.Context) {

			fmt.Println("stop")
		},
	)

	app.Get("/panic/", func(_ *server.Context) {

		panic("aaa")
	})

	app.Get("/test/:param/", func(c *server.Context) {

		fmt.Fprint(c, c.Request.Param("param"))
		fmt.Fprint(c, c.Request.IsPost())
	})

	app.Post("/post/", func(_ *server.Context) {

	})

	app.Get("/json/", func(c *server.Context) {

		render.JSON(c, map[string]string{"a": "A", "b": "B"})
	})

	app.Get("/xml/", func(c *server.Context) {

		type XML struct {
			A string `xml:"a"`
			B string `xml:"b"`
		}

		render.XML(c, XML{A: "A", B: "B"})
	})

	app.Get("/html/subdirectory/", func(c *server.Context) {

		user := User{Name: "Undefined"}

		if u, ok := c.Get("user").(User); ok {

			user = u
		}

		pongo2.HTML(c, "subdirectory/main.html", map[string]interface{}{"user": user})
	})

	go func() {
		//pprof
		log.Println(http.ListenAndServe(":8081", nil))
	}()

	app.Run(":8080")
}

func errorHandlers(app *server.App) {

	app.PanicHandler = func(rw http.ResponseWriter, _ *http.Request, err interface{}) {

		rw.Header().Set("Content-Type", "text/html")

		rw.WriteHeader(http.StatusInternalServerError)

		fmt.Fprint(rw, string(debug.Stack()))

		log.Printf("PANIC: %s\n", debug.Stack())
	}

	app.NotFound = func(rw http.ResponseWriter, _ *http.Request) {

		fmt.Fprint(rw, "Not found")
	}

	app.MethodNotAllowed = func(rw http.ResponseWriter, _ *http.Request) {

		fmt.Fprint(rw, "MethodNotAllowed")
	}
}
