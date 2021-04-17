package main

import (
	//"fmt"
	"net/http"
	"niu"
)

func main() {
	app := niu.New()


	app.Get("/go", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello niu")) })


	app.Run("127.0.0.1:8888")
}



