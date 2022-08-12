package main

type HandlerBaseOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  handlerFunc
}
