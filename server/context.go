package server

import (
	"net/http"
	"sync"
)

type Context struct {
	http.ResponseWriter

	mutex    *sync.Mutex
	Request  request
	values   map[string]interface{}
	handlers []Handler
	index    int
	stop     bool
}

func (c *Context) Set(key string, value interface{}) {

	c.mutex.Lock()

	c.values[key] = value

	c.mutex.Unlock()
}

func (c *Context) Get(key string) interface{} {

	defer c.mutex.Unlock()

	c.mutex.Lock()

	if v, found := c.values[key]; found {

		return v
	}

	return nil
}

func (c *Context) Next() {

	c.index++

	c.run()
}

func (c *Context) Stop() {

	c.stop = true
}

func (c *Context) run() {

	for c.index < len(c.handlers) {

		c.handlers[c.index](c)

		c.index++

		if c.stop {

			return
		}
	}
}
