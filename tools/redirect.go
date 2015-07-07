package tools

import (
	"github.com/kshvakov/vrum-vrum/server"
	"net/http"
)

func Redirect(c *server.Context, urlStr string, code int) {
	// FF
	c.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header().Set("Expires", "0")

	http.Redirect(c.ResponseWriter, c.Request.HTTP, urlStr, code)

	c.Stop()
}
