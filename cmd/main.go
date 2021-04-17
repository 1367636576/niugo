package main

import (
	//"fmt"
	//"net/http"
	"niu"
	"niu/context"
)

func main() {
	app := niu.New()



	app.Get("/go", func(ctx context.Context) { ctx.String("hello niu! a wondderful web framework!") })


	app.Run("127.0.0.1:8888")
}



