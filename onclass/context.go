package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, res interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(res)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

func NewContext(write http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: write,
		R: request,
	}
}

func (c *Context) BadRequestJson(data interface{}) error {
	return c.WriteJson(http.StatusBadRequest, data)
}

func (c *Context) OkRequestJson(data interface{}) error {
	return c.WriteJson(http.StatusOK, data)
}
