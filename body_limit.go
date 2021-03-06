// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

/*
Package bodylimit is a HTTP middleware that limit the request body size.

Example

	package main

	import (
		"log"

		"github.com/go-gem/gem"
		"github.com/go-gem/middleware-body-limit"
	)

	func main() {
		// maximum size: 2M.
		limiter := bodylimit.New(int64(2 * 1024))
		router := gem.NewRouter()
		router.POST("/upload", func(ctx *gem.Context) {
			// upload files.
		}, &gem.HandlerOption{Middlewares: []gem.Middleware{limiter}})
		log.Println(gem.ListenAndServe(":8080", router.Handler()))
	}
*/
package bodylimit

import (
	"net/http"

	"github.com/go-gem/gem"
)

// New returns Limiter instance by the
// given maximum allowed size.
func New(max int64) *Limiter {
	return &Limiter{
		max: max,
	}
}

// Limiter request body limit middleware.
type Limiter struct {
	// Maximum allowed size for a request body,
	// it's unit is byte.
	max int64
}

// Wrap implements Middleware's interface.
func (l *Limiter) Wrap(next gem.Handler) gem.Handler {
	return gem.HandlerFunc(func(ctx *gem.Context) {
		// response Bad Request if the content length is unknown.
		if ctx.Request.ContentLength == -1 || (ctx.Request.ContentLength == 0 && ctx.Request.Body != nil) {
			ctx.Response.WriteHeader(http.StatusBadRequest)
			return
		}

		// response Request Entity Too Large if content length is
		// large than maximum size.
		if ctx.Request.ContentLength > l.max {
			ctx.Response.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}

		next.Handle(ctx)
	})
}
