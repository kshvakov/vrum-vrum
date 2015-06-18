package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type request struct {
	HTTP   *http.Request
	params httprouter.Params
}

func (r *request) Param(key string) string {

	return r.params.ByName(key)
}

func (r *request) IsPost() bool {

	return r.HTTP.Method == "POST"
}
