package main

import (
	"net/http"
)

type Server interface {
	Routable
	Start(address string) error
}

//http服务结构体
type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(c *Context)) {
	/*http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer,request)
		handleFunc(c)
	})*/
	//生成请求的方法+路由key
	//key := s.handler.key(method,pattern)

	//需要初始化一下 不然会报空指针引用  推荐在上层初始化
	/*s.handler = &HandlerBaseOnMap{
		handlers : make(map[string]func(c *Context)),
	}*/
	//赋值处理方法 交由路由树匹配的方法处理
	//s.handler.handlers[key] = handleFunc

	s.handler.Route(method, pattern, handleFunc)

}

func (s *sdkHttpServer) Start(address string) error {

	http.HandleFunc("/", func(writer http.ResponseWriter,request *http.Request) {
		c:= NewContext(writer,request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}

//生成一个sdkHttpServer结构体  该结构体实现了Server接口定义的两个路由

func NewHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBaseOnMap()
	var root Filter = handler.ServeHTTP
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	Data int
}

func SignUp(c *Context) {
	req := &signUpReq{}

	err := c.ReadJson(req)
	if err != nil {
		c.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}

	err = c.WriteJson(http.StatusOK, resp)

	if err != nil {
		c.BadRequestJson(err)
	}
}
