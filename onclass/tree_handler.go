package main

import (
	"net/http"
	"strings"
)

type HandlerBaseOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  handlerFunc
}

func NewHandlerBasedOnTree() Handler {
	return &HandlerBaseOnTree{
		root: &node{},
	}
}

func (h *HandlerBaseOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBaseOnTree) findRouter(path string) (handlerFunc, bool) {
	// 去除头尾可能有的/，然后按照/切割成段
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		// 从子节点里边找一个匹配到了当前 path 的节点
		matchChild, found := cur.findChildTree(p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	// 到这里，应该是找完了
	if cur.handler == nil {
		// 到达这里是因为这种场景
		// 比如说你注册了 /user/profile
		// 然后你访问 /user
		return nil, false
	}
	return cur.handler, true
}

func (h *HandlerBaseOnTree) Route(method string, pattern string, handleFunc func(c *Context)) {

	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root

	for index, path := range paths {
		//先查找子节点
		//如果找到了子节点，当前根节点替换为子节点 继续往下查找  没有的话执行创建
		child, find := h.root.findChildTree(path)
		if find {
			cur = child
			continue
		} else {
			//创建子节点
			cur.createRootTree(paths[index:], handleFunc)
			return
		}
	}

	//循环完了就是最后一级了 赋值进去
	cur.handler = handleFunc
}

func (n *node) findChildTree(path string) (*node, bool) {
	for _, child := range n.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

func (n *node) createRootTree(path []string, handlefunc handlerFunc) {

	cur := n
	for _, childPath := range path {
		newPath := newPath(childPath)
		newPath.handler = handlefunc
		n.children = append(n.children, newPath)
		cur = newPath
	}
	cur.handler = handlefunc
}

func newPath(path string) *node {

	return &node{
		path:     path,
		children: make([]*node, 0, 100),
	}
}
