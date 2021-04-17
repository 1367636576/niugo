package main

import (
	"fmt"
	"net/http"
	"niu/router"
	"niu"
)

func main() {
	app := &niu.Application{}

	router := &router.Route{}
	router.RList = make(map[string]func(http.ResponseWriter, *http.Request))
	router.Add("/go", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello niu")) })


	app.Run("127.0.0.1:8888", router)
}



