// Code generated by hertz generator.

package main

import (
	"fzuhelper-empty-room/biz/dal"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	dal.Init()
}
func main() {
	h := server.Default()
	dal.Init()
	register(h)
	h.Spin()
}
