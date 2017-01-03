// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package bodylimitmidware

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/go-gem/gem"
)

func TestBodyLimit(t *testing.T) {
	body := []byte("hello")
	midware := New(int64(len(body) - 1))

	var pass bool

	handler := midware.Wrap(gem.HandlerFunc(func(ctx *gem.Context) {
		pass = true
	}))

	req := httptest.NewRequest("/", "/", nil)
	resp := httptest.NewRecorder()
	ctx := &gem.Context{Request: req, Response: resp}

	handler.Handle(ctx)
	if pass {
		t.Error("expected that the middleware blocked this request, but failed")
	}

	req = httptest.NewRequest("/", "/", bytes.NewReader(body))
	resp = httptest.NewRecorder()
	ctx = &gem.Context{Request: req, Response: resp}
	handler.Handle(ctx)
	if pass {
		t.Error("expected that the middleware blocked this request, but failed")
	}

	req = httptest.NewRequest("/", "/", bytes.NewReader([]byte("foo")))
	resp = httptest.NewRecorder()
	ctx = &gem.Context{Request: req, Response: resp}
	handler.Handle(ctx)
	if !pass {
		t.Error("failed to pass the middleware")
	}
}
