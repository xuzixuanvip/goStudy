package main

import (
	"net/http"
)

/*type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}*/

type Server interface {
	Route(method string,pattern string, handleFunc func(c *Context))
	Start(address string) error
}


//http服务结构体
type sdkHttpServer struct {
	Name string
	handler *HandlerBaseOnMap
}


func (s *sdkHttpServer) Route(method string,pattern string, handleFunc func(c *Context)) {
	/*http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer,request)
		handleFunc(c)
	})*/
	//生成请求的方法+路由key
	key := s.handler.key(method,pattern)

	//赋值处理方法 交由路由树匹配的方法处理
	s.handler.handlers[key] = handleFunc

}

/*func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handleFunc)
}*/

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, s.handler)
}

//生成一个sdkHttpServer结构体  该结构体实现了Server接口定义的两个路由

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
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
