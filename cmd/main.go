package main

import (
	//"fmt"
	"net/http"
	"niu/router"
	"niu"
)

func main() {
	app := &niu.Application{}

	routerInstance := &router.Router{}
	routerInstance.RList = make(map[string] router.PatternHandlerMap)
	routerInstance.Get("/go", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello niu")) })


	app.Run("127.0.0.1:8888", routerInstance)
}



