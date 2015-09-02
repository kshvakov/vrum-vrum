package server

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func New() *App {

	return &App{
		router: httprouter.New(),
	}
}

var CollectRequestsStat = true

type App struct {
	router           *httprouter.Router
	handlers         []Handler
	PanicHandler     func(response http.ResponseWriter, _ *http.Request, err interface{})
	NotFound         http.HandlerFunc
	MethodNotAllowed http.HandlerFunc
}

func (a *App) Use(handler Handler) {

	a.handlers = append(a.handlers, handler)
}

func (a *App) Get(pattern string, handlers ...Handler) {

	handler := handler(append(a.handlers, handlers...))

	a.router.GET(pattern, handler)
	a.router.HEAD(pattern, handler)
}

func (a *App) Post(pattern string, handlers ...Handler) {

	a.router.POST(pattern, handler(append(a.handlers, handlers...)))
}

func (a *App) Form(pattern string, handlers ...Handler) {

	handler := handler(append(a.handlers, handlers...))

	a.router.GET(pattern, handler)
	a.router.POST(pattern, handler)
}

func (a *App) ServeFiles(path string, root http.FileSystem) {

	a.router.ServeFiles(path, root)
}

func (a *App) Run(addr string) {

	if a.PanicHandler != nil {

		a.router.PanicHandler = a.PanicHandler
	}

	if a.NotFound != nil {

		a.router.NotFound = a.NotFound
	}

	if a.MethodNotAllowed != nil {

		a.router.MethodNotAllowed = a.MethodNotAllowed
	}

	a.router.GET("/metrics/", handler([]Handler{metricsHandler}))

	log.Fatal(http.ListenAndServe(addr, a.router))
}
