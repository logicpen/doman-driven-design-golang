package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler
type Chain []Middleware

func CreateMiddlewareChain(middlewares ...Middleware) Chain {
	var chain []Middleware
	return append(chain, middlewares...)
}

func (c Chain) Then(originalHandler http.Handler) http.Handler {
	for i := len(c) - 1; i >= 0; i-- {
		originalHandler = c[i](originalHandler)
	}
	return originalHandler
}
