// Code generated by hertz generator. DO NOT EDIT.

package main

import (
	router "fzuhelper-empty-room/biz/router"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
