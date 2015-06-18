package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
)

type Handler func(c *Context)

func handler(handlers []Handler) func(http.ResponseWriter, *http.Request, httprouter.Params) {

	return func(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {

		stop := timerStart(req.URL.Path)

		c := Context{
			mutex:          &sync.Mutex{},
			ResponseWriter: rw,
			Request: request{
				HTTP:   req,
				params: params,
			},
			handlers: handlers,
			values:   make(map[string]interface{}),
		}

		c.run()

		stop()
	}
}
