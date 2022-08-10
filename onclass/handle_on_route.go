package main

import (
	"fmt"
	"net/http"
)

type Handler interface {
	http.Handler
}


type HandlerBaseOnMap struct {
	handlers map[string]func(c *Context)
}

func (h *HandlerBaseOnMap) ServeHTTP(write http.ResponseWriter, request *http.Request) {

	//先使用方法生成路由的key
	key := h.key(request.Method, request.URL.Path)

	//如果路由存在于路由组中
	if handler, ok := h.handlers[key]; ok {
		//实例化content
		c := NewContext(write,request)
		//方法传入content
		handler(c)
	}else{
		//返回未找到
		write.WriteHeader(http.StatusNotFound)
		_,_ = write.Write([]byte("not any router match"))
	}

}

//生成路由map的key
func (h *HandlerBaseOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s#%s", method, path)
}
