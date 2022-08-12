package main

import (
	"fmt"
	"net/http"
)


type Routable interface {
	Route(method string, pattern string, handleFunc func(c *Context))
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

type HandlerBaseOnMap struct {
	handlers map[string]func(c *Context)
}


func (h *HandlerBaseOnMap) Route(method string, pattern string, handleFunc func(c *Context)) {

	//生成请求的方法+路由key
	key := h.key(method, pattern)

	//赋值处理方法 交由路由树匹配的方法处理
	h.handlers[key] = handleFunc

}


func (h *HandlerBaseOnMap) ServeHTTP(c *Context) {

	//先使用方法生成路由的key
	key := h.key(c.R.Method, c.R.URL.Path)

	//如果路由存在于路由组中
	if handler, ok := h.handlers[key]; ok {
		//方法传入content
		handler(c)
	} else {
		//返回未找到
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("not any router match"))
	}
}

//生成路由map的key
func (h *HandlerBaseOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}

func NewHandlerBaseOnMap() Handler {
	return &HandlerBaseOnMap{
		handlers: make(map[string]func(c *Context)),
	}
}
