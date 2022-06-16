package main

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestNewApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	// redirects to /admin without basic auth
	e.GET("/").Expect().Status(httptest.StatusUnauthorized)
	// without basic auth
	e.GET("/admin").Expect().Status(httptest.StatusUnauthorized)

	// with valid basic auth
	e.GET("/").WithBasicAuth("kataras", "kataras_pass").Expect().
		Status(httptest.StatusOK).Body().Equal("/ - Hello kataras:kataras_pass")
	e.GET("/admin/dashboard").WithBasicAuth("kataras", "kataras_pass").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/dashboard - Hello kataras:kataras_pass")
	e.GET("/admin/profile").WithBasicAuth("kataras", "kataras_pass").Expect().
		Status(httptest.StatusOK).Body().Equal("/admin/profile - Hello kataras:kataras_pass")

}
