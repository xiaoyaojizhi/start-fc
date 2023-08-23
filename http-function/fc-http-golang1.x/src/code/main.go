package main

import (
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"main/handler"
)

func main() {
	fc.StartHttp(handler.HandleHttpRequest)
}
