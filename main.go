package main

import (
	helloworld "github.com/LiamZhuangDev/gin/hello_world"
	"github.com/LiamZhuangDev/gin/routing"
)

func main() {
	helloworld.Pong()
	routing.HttpMethods()
	routing.PathParams()
	routing.QueryParams()
	routing.FormParams()
	routing.JSONFormParams()
}
