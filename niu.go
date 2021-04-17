package niu

import "net/http"
//import "fmt"
import "niu/router"

type Application struct {
	//router
	route *router.Router

}

func (a *Application) Run(addr string,router *router.Router) {
	http.ListenAndServe(addr, router)
}


