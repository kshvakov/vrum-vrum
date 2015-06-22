package tools

import (
	"github.com/kshvakov/vrum-vrum/server"
	"net/http"
)

func Redirect(c *server.Context, urlStr string, code int) {

	http.Redirect(c.ResponseWriter, c.Request.HTTP, urlStr, code)

	c.Stop()
}
