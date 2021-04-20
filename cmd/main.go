package main

import (
	//"fmt"
	//"net/http"
	"niu"
	"niu/context"
)

func hello(ctx context.Context)  {
	ctx.String("hello niu! a wondderful web framework!")
}

func main() {
	app := niu.New()



	app.Get("/go", func(ctx context.Context) {

		//ctx.String("hello niu! a wondderful web framework!")
	})
	app.Post("/go2", hello)
	app.Get("/hello", func(ctx context.Context) { ctx.String("hello world") })


	app.Run("127.0.0.1:8888")
}



