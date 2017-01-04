# Request Body Limit Middleware

[![Build Status](https://travis-ci.org/go-gem/middleware-body-limit.svg?branch=master)](https://travis-ci.org/go-gem/middleware-body-limit)
[![GoDoc](https://godoc.org/github.com/go-gem/middleware-body-limit?status.svg)](https://godoc.org/github.com/go-gem/middleware-body-limit)
[![Coverage Status](https://coveralls.io/repos/github/go-gem/middleware-body-limit/badge.svg?branch=master)](https://coveralls.io/github/go-gem/middleware-body-limit?branch=master)

Request Body Limit middleware for [Gem](https://github.com/go-gem/gem) Web framework.

## Getting Started

**Install**

```
$ go get -u github.com/go-gem/middleware-body-limit
```

**Example**

```
package main

import (
	"log"

	"github.com/go-gem/gem"
	"github.com/go-gem/middleware-body-limit"
)

func main() {
	limiter := bodylimit.New(int64(1024))
	router := gem.NewRouter()
	router.POST("/upload", func(ctx *gem.Context) {
		// upload files.
	}, &gem.HandlerOption{Middlewares: []gem.Middleware{limiter}})
	log.Println(gem.ListenAndServe(":8080", router.Handler()))
}
```